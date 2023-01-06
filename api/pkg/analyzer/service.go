// Copyright 2022 Cisco Systems, Inc. and its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package analyzer

import (
	"context"
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/db"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/cisco-developer/api-insights/api/pkg/analyzer/apiclarity"
	"github.com/cisco-developer/api-insights/api/pkg/analyzer/completeness"
	"github.com/cisco-developer/api-insights/api/pkg/analyzer/contract"
	"github.com/cisco-developer/api-insights/api/pkg/analyzer/documentation"
	"github.com/cisco-developer/api-insights/api/pkg/analyzer/guidelines"
	"github.com/cisco-developer/api-insights/api/pkg/analyzer/security"
	"github.com/cisco-developer/api-insights/api/pkg/analyzer/woke"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"time"
)

type Service interface {
	Analyze(req *models.SpecAnalysisRequest) (*models.SpecAnalysisResponse, error)
	Reporter(optionalAnalyzers map[analyzer.SpecAnalyzer]*analyzer.Analyzer) (Reporter, error)
}

func NewService(analyzerLister func(ctx context.Context, filter *db.ListFilter, withRules bool) ([]*analyzer.Analyzer, error)) (Service, error) {
	svc := &service{
		listAnalyzers: analyzerLister,
	}
	return svc, nil
}

type service struct {
	listAnalyzers func(ctx context.Context, filter *db.ListFilter, withRules bool) ([]*analyzer.Analyzer, error)
}

func (s *service) listActiveAnalyzers() ([]*analyzer.Analyzer, error) {
	analyzers, err := s.listAnalyzers(context.Background(), &db.ListFilter{
		Model:   &analyzer.Analyzer{},
		Indexes: map[string]string{"status": analyzer.AnalyzerStatusActive},
	}, true)
	if err != nil {
		return nil, err
	}
	return analyzers, nil
}

func (s *service) Reporter(optionalAnalyzers map[analyzer.SpecAnalyzer]*analyzer.Analyzer) (Reporter, error) {
	var analyzers map[analyzer.SpecAnalyzer]*analyzer.Analyzer
	if len(optionalAnalyzers) > 0 {
		analyzers = optionalAnalyzers
	} else {
		analyzerList, err := s.listActiveAnalyzers()
		if err != nil {
			return nil, err
		}
		analyzers = analyzer.ListToMap(analyzerList)
	}
	analyzersScoreConfigs, err := analyzer.NewAnalyzersScoreConfigsFrom(analyzers)
	if err != nil {
		return nil, err
	}
	reporter, err := NewReporter(analyzersScoreConfigs, analyzers)
	if err != nil {
		return reporter, err
	}
	return reporter, nil
}

func (s *service) Analyze(req *models.SpecAnalysisRequest) (*models.SpecAnalysisResponse, error) {
	if !req.HasSpec() {
		return nil, fmt.Errorf("analyzer: SpecAnalysisRequest.Spec cannot be nil")
	}

	res := &models.SpecAnalysisResponse{
		Results: make(map[analyzer.SpecAnalyzer]*models.SpecAnalysis, len(req.Analyzers)),
	}

	for _, analyzerName := range req.Analyzers {
		var (
			cfg            = req.AnalyzersConfigs[analyzerName]
			analyzerClient models.SpecDocAnalyzer
			err            error
		)
		switch analyzerName {
		case analyzer.CiscoAPIGuidelines:
			analyzerClient, err = guidelines.NewClient()
		case analyzer.Completeness:
			analyzerClient, err = completeness.NewClient()
		case analyzer.Contract:
			analyzerClient, err = contract.NewClient()
		case analyzer.Documentation:
			analyzerClient, err = documentationp.NewClient()
		case analyzer.InclusiveLanguage:
			analyzerClient, err = woke.NewClient()
		case analyzer.Drift:
			analyzerClient, err = apiclarity.NewClient()
		case analyzer.Security:
			analyzerClient, err = security.NewClient()
		default:
			return nil, fmt.Errorf("analyzer: unsupported analyzer(%s)", analyzerName)
		}
		if err != nil {
			return nil, fmt.Errorf("failed to create analyzer(%s): %v", analyzerName, err)
		}
		var serviceNameID string
		if req.Service != nil {
			serviceNameID = req.Service.GetNameID(analyzerName, cfg)
		}
		result, err := analyzerClient.Analyze(req.Spec.Doc, cfg, &serviceNameID)
		if err != nil {
			shared.LogErrorf("failed to run analyzer(%s): %v", analyzerName, err)
			continue
			//return nil, err  // TODO Handle analyzer failures.
		}
		now := time.Now().UTC()
		specAnalysis := &models.SpecAnalysis{
			ID:        shared.TimeUUID(),
			Analyzer:  analyzerName,
			ServiceID: req.Spec.ServiceID,
			SpecID:    req.Spec.ID,
			CreatedAt: now,
			UpdatedAt: now,
		}
		specAnalysis.Config = cfg

		if err := specAnalysis.SetResult(result, "Analyzed"); err != nil {
			return nil, err
		}

		res.Results[analyzerName] = specAnalysis
	}

	reporter, err := s.Reporter(req.ActiveAnalyzers)
	if err != nil {
		return nil, err
	}

	serviceSpecReport, err := reporter.GenerateSpecReport(req.Spec, res.Results)
	if err != nil {
		return nil, err
	}
	for analyzerName, specAnalysis := range res.Results {
		specAnalysisReport, ok := serviceSpecReport.specAnalysisReports[analyzerName]
		if ok {
			if err := specAnalysis.SetScore(specAnalysisReport.score); err != nil {
				return nil, err
			}
		}
	}
	res.SpecScore = serviceSpecReport.score

	return res, nil
}

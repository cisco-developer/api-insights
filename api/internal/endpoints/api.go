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

package endpoints

import (
	"encoding/json"
	"github.com/cisco-developer/api-insights/api/internal/access"
	"github.com/cisco-developer/api-insights/api/internal/db"
	"github.com/cisco-developer/api-insights/api/internal/info"
	"github.com/cisco-developer/api-insights/api/internal/middleware"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/pkg/analyzer"
	"github.com/cisco-developer/api-insights/api/pkg/apiclarity"
	"github.com/cisco-developer/api-insights/api/pkg/differ"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"github.com/emicklei/go-restful/v3"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

var (
	se restful.ServiceError

	// sorting query params
	sort      = restful.QueryParameter("sort", "csv fields for sorting").DataType("string").DefaultValue("").CollectionFormat(restful.CollectionFormatCSV)
	sortOrder = restful.QueryParameter("order", "sorting order: asc(default), desc").DataType("string").DefaultValue("asc").AllowableValues(map[string]string{"asc": "asc", "desc": "desc"})

	max = restful.QueryParameter("max", "max").DataType("string").DefaultValue("100")
)

// API Exposes the restful container for API's
// API: /v1/
func API(cfg *shared.AppConfig) (*restful.Container, error) {
	db.Init(cfg)

	_, _, err := shared.InitTracer(cfg.AppName)
	if err != nil {
		return nil, err
	}

	container := restful.NewContainer()
	container.EnableContentEncoding(true)
	container.Filter(middleware.NewTracingMiddleware())

	validate := validator.New()

	runtimeServerInfo, err := info.GetInfo()
	if err != nil {
		return nil, err
	}

	differSvc, err := differ.NewService()
	if err != nil {
		return nil, err
	}

	specDao, err := db.NewSpecDAO(cfg)
	if err != nil {
		return nil, err
	}

	analyzerDao, err := db.NewAnalyzerDAO(cfg)
	if err != nil {
		return nil, err
	}

	analyzerRuleDao, err := db.NewAnalyzerRuleDAO(cfg)
	if err != nil {
		return nil, err
	}

	infoRes := &infoResource{
		config:   cfg,
		validate: validate,
		info:     runtimeServerInfo,
	}
	infoRes.Register(cfg, container, "/v1/apiregistry/info")

	analyzerRes := &analyzerResource{
		config:          cfg,
		dao:             analyzerDao,
		validate:        validate,
		analyzerRuleDao: analyzerRuleDao,
	}
	analyzerRes.Register(cfg, container, "/v1/apiregistry/analyzers")

	specAnalysisDao, err := db.NewSpecAnalysisDAO(cfg)
	if err != nil {
		return nil, err
	}

	analyzerSvc, err := analyzer.NewService(analyzerDao.List)
	if err != nil {
		return nil, err
	}

	specAnalysisRes := &specAnalysisResource{
		config:      cfg,
		dao:         specAnalysisDao,
		validate:    validate,
		analyzerSvc: analyzerSvc,
	}
	specAnalysisRes.Register(cfg, container, "/v1/apiregistry/specs/analyses")

	specDiffDao, err := db.NewSpecDiffDAO(cfg)
	if err != nil {
		return nil, err
	}

	specDiffRes := &specDiffResource{
		config:    cfg,
		dao:       specDiffDao,
		validate:  validate,
		differSvc: differSvc,
		specDAO:   specDao,
	}
	specDiffRes.Register(cfg, container, "/v1/apiregistry/specs/diffs")

	specValidationRes := &specValidationResource{
		config: cfg,
	}
	specValidationRes.Register(cfg, container, "/v1/apiregistry/specs/validations")

	serviceDao, err := db.NewServiceDAO(cfg)
	if err != nil {
		return nil, err
	}

	apiclarityClient, err := apiclarity.New(nil)
	if err != nil {
		return nil, err
	}

	accessChecker, err := access.NewNoopAccessChecker()
	if err != nil {
		return nil, err
	}

	organizationDao, err := db.NewOrganizationDAO(cfg)
	if err != nil {
		return nil, err
	}

	organizationRes := &organizationResource{
		config:        cfg,
		dao:           organizationDao,
		validate:      validate,
		accessChecker: accessChecker,
	}
	organizationRes.Register(cfg, container, "/v1/apiregistry/organizations")

	serviceRes := &serviceResource{
		config:           cfg,
		dao:              serviceDao,
		validate:         validate,
		specDAO:          specDao,
		specDiffDAO:      specDiffDao,
		specAnalysisDAO:  specAnalysisDao,
		analyzerDAO:      analyzerDao,
		organizationDAO:  organizationDao,
		analyzerSvc:      analyzerSvc,
		differSvc:        differSvc,
		apiclarityClient: apiclarityClient,
		accessChecker:    accessChecker,
		info:             runtimeServerInfo,
	}
	serviceRes.Register(cfg, container, "/v1/apiregistry/services")

	hr := HealthCheckResource{}
	hr.Register(container, "/v1/healthz")

	return container, nil
}

func handleError(res *restful.Response, err error) {
	if schemaErr, ok := err.(*ValidationError); ok {
		result := struct {
			Errors map[string][]string `json:"errors"`
		}{map[string][]string{}}
		err = json.Unmarshal([]byte(schemaErr.Error()), &result)
		if err != nil {
			_ = res.WriteHeaderAndEntity(http.StatusBadRequest, schemaErr.Error())
		} else {
			_ = res.WriteHeaderAndEntity(http.StatusBadRequest, result)
		}
	} else if _, ok := err.(*models.UnauthorizedResourceAccessError); ok {
		res.WriteHeader(http.StatusForbidden)
	} else {
		res.WriteHeader(http.StatusInternalServerError)
	}
}

// parseQueryBool returns true if query flag is given and true, otherwise false.
func parseQueryBool(req *restful.Request, flag string) bool {
	b, err := strconv.ParseBool(req.QueryParameter(flag))
	if err == nil && b {
		return true
	}

	return false
}

// ValidationError - Validation error class.
type ValidationError struct {
	jsonErrors []byte
}

// Error - Print error as string.
func (err *ValidationError) Error() string {
	return string(err.jsonErrors)
}

// MarshalJSON - Print error as json bytes.
func (err *ValidationError) MarshalJSON() ([]byte, error) {
	return err.jsonErrors, nil
}

func orgServiceAccessDataFilterFromReq(req *restful.Request) *models.OrgServiceAccessDataFilter {
	filtersRaw := middleware.AccessDataFiltersFromReq(req)
	for _, filterRaw := range filtersRaw {
		filter, ok := filterRaw.(*models.OrgServiceAccessDataFilter)
		if ok {
			return filter
		}
	}
	return nil
}

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

package model

import (
	"fmt"
	"github.com/cisco-developer/api-insights/cli/pkg/utils"
	"github.com/olekukonko/tablewriter"
	"io"
	"strconv"
	"time"
)

const (
	SpecAnalyzerCiscoAPIGuidelines = SpecAnalyzer("guidelines")
)

type SpecAnalyzer string

var severityConfigs = map[SeverityName]struct {
	Color tablewriter.Colors
}{
	SeverityNameError:   {tablewriter.Colors{tablewriter.Normal, tablewriter.FgRedColor}},
	SeverityNameWarning: {tablewriter.Colors{tablewriter.Normal, tablewriter.FgYellowColor}},
	SeverityNameInfo:    {tablewriter.Colors{tablewriter.Normal, tablewriter.FgBlueColor}},
	SeverityNameHint:    {tablewriter.Colors{tablewriter.Normal, tablewriter.Normal}},
}

// SpecAnalysis represents a specAnalysis
type SpecAnalysis struct {
	ID        string          `json:"id,omitempty" validate:"required"`
	Analyzer  SpecAnalyzer    `json:"analyzer"`
	Config    *AnalyzerConfig `json:"config,omitempty"`
	Result    *Result         `json:"result"`
	Score     int             `json:"score"`
	ServiceID string          `json:"service_id"`
	SpecID    string          `json:"spec_id"`
	Status    string          `json:"status"` // Submitted, Invalid, Analyzed
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type SpecAnalysisList []*SpecAnalysis

func (m SpecAnalysisList) Print(w io.Writer) {
	table := utils.NewTable(w)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"#", "ID", "Analyzer", "Status", "Score", "Spec ID", "Service ID"})

	for i, analysis := range m {
		row := []string{fmt.Sprintf("%d", i+1), analysis.ID, string(analysis.Analyzer), analysis.Status, fmt.Sprint(analysis.Score), analysis.SpecID, analysis.ServiceID}
		table.Append(row)
	}

	table.Render()
}

// AnalyzerConfig represents configs for an analyzer (SpecAnalysis.Analyzer)
type AnalyzerConfig map[string]interface{}

// SpecAnalysisRequest represents a request for a SpecAnalysis
type SpecAnalysisRequest struct {
	Analyzers        []SpecAnalyzer                   `json:"analyzers"`
	AnalyzersConfigs map[SpecAnalyzer]*AnalyzerConfig `json:"analyzers_configs,omitempty"`

	Spec    *Spec    `json:"spec,omitempty"`
	Service *Service `json:"service,omitempty"`
}

type SpecAnalysisResponse struct {
	Results   map[SpecAnalyzer]*SpecAnalysis `json:"results,omitempty"`
	SpecScore int                            `json:"spec_score"`
}

// HasErrorFindings returns true if there are any error findings, otherwise false.
func (s *SpecAnalysisResponse) HasErrorFindings() bool {
	for _, analysis := range s.Results {
		if analysis.Result.Summary.Stats.Error != nil && analysis.Result.Summary.Stats.Error.Count > 0 {
			return true
		}
	}

	return false
}

// Print prints results in console
func (s *SpecAnalysisResponse) Print(w io.Writer, analyzers map[string]*Analyzer) {
	type analysisSummary struct {
		Analyzer string
		Score    string
		Error    string
		Warning  string
		Info     string
		Hint     string
	}
	var analysisSummaries []analysisSummary
	normalColor := tablewriter.Colors{tablewriter.Normal, tablewriter.Normal}

	for analyzer, analysis := range s.Results {
		table := utils.NewTable(w)
		table.SetHeader([]string{"#", "Severity", "Code", "Findings", "Recommendation", "Affected Items"})
		severityColumnIndex := 1
		columColors := []tablewriter.Colors{normalColor, normalColor, normalColor, normalColor, normalColor, normalColor}

		a := analyzers[string(analyzer)]
		fmt.Printf("%s Compliance\n", a.Title)

		i := 0
		for severityName, findings := range analysis.Result.Findings {
			for code, finding := range findings.Rules {
				row := []string{fmt.Sprintf("%d", i+1), string(severityName), string(code), finding.Message, finding.Mitigation, fmt.Sprintf("%d", len(finding.Data))}

				if conf, found := severityConfigs[severityName]; found {
					columColors[severityColumnIndex] = conf.Color
				} else {
					columColors[severityColumnIndex] = normalColor
				}

				table.Rich(row, columColors)
				i++
			}
		}

		summary := analysis.Result.Summary.String()
		table.SetCaption(true, summary)

		table.Render()
		fmt.Println()

		stats := analysis.Result.Summary.Stats
		analysisSummaries = append(analysisSummaries, analysisSummary{
			Analyzer: a.Title,
			Score:    fmt.Sprintf("%d", analysis.Score),
			Error:    strconv.Itoa(stats.TotalError()),
			Warning:  strconv.Itoa(stats.TotalWarning()),
			Info:     strconv.Itoa(stats.TotalInfo()),
			Hint:     strconv.Itoa(stats.TotalHint()),
		})
	}

	// summary
	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"#", "Analyzer", "Score", "Error", "Warning", "Info", "Hint"})
	table.SetBorder(false)
	table.SetAutoWrapText(false)
	for i, as := range analysisSummaries {
		rows := []string{fmt.Sprintf("%d", i+1), as.Analyzer, as.Score, as.Error, as.Warning, as.Info, as.Hint}
		table.Append(rows)
	}
	table.Render()

	fmt.Printf("\nAPI Score: %d\n", s.SpecScore)
}

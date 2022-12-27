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

package cmd

import (
	"errors"
	"fmt"
	"github.com/cisco-developer/api-insights/cli/pkg/model"
	"github.com/cisco-developer/api-insights/cli/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	// used for stateless analysis
	dummySpecID        = "10000000-0000-0000-0000-000000000000"
	dummyServiceID     = "20000000-0000-0000-0000-000000000000"
	dummyServiceNameID = "placeholder"

	flagAnalyzer        = "analyzer"
	flagFailBelowScore  = "fail-below-score"
	flagFailOnErrorRule = "fail-on-error-rule"
)

var (
	analyzer        string
	failBelowScore  int
	failOnErrorRule bool
)

func init() {
	rootCmd.AddCommand(analyzeCmd())
}

var (
	apiSpecFilePath = ""
)

func analyzeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "analyze LOCAL_SPEC",
		Short: "Analyze local API spec",
		Example: `  # Analyze local spec with all analyzers
  api-insights-cli analyze testdata/carts.json

  # Analyze local spec with specific analyzer
  api-insights-cli analyze testdata/carts.json --analyzer guidelines
  api-insights-cli analyze testdata/carts.json --analyzer completeness
  api-insights-cli analyze testdata/carts.json --analyzer inclusive-language
  api-insights-cli analyze testdata/carts.json --analyzer drift
  api-insights-cli analyze testdata/carts.json --analyzer security`,
		Run:  analyzeSpec,
		Args: cobra.MinimumNArgs(1),
	}

	cmd.Flags().StringVarP(&analyzer, flagAnalyzer, "a", "", "API spec analyzer")
	cmd.Flags().IntVarP(&failBelowScore, flagFailBelowScore, "", 0, "Fail if API score is below specified score, defaults to 0")
	cmd.Flags().BoolVarP(&failOnErrorRule, flagFailOnErrorRule, "", false, "Fail if there are any error findings")
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		logDebugln("Failed to bind flags", err.Error())
	}

	return cmd
}

func analyzeSpec(cmd *cobra.Command, args []string) {
	logDebugln("started")
	if len(args) < 1 {
		utils.ExitWithCode(utils.ExitInvalidInput, errors.New("API spec file is required, for example: api-insights-cli analyze api.yaml"))
	}

	filename := args[0]
	logDebugf("loading local spec: %s\n", filename)

	spec, err := os.ReadFile(filename)
	if err != nil {
		utils.ExitWithCode(utils.ExitError, fmt.Errorf("failed to load spec: %s", err.Error()))
	}
	logDebugf("loaded local spec: %s\n", filename)

	analyzers, err := apiInsightsClient.ListAnalyzers(cmd.Context(), map[string]string{"status": "active"})
	if err != nil {
		utils.ExitWithCode(utils.ExitError, fmt.Errorf("failed to list analyzers: %s", err.Error()))
	}
	as := map[string]*model.Analyzer{}
	var choices []string
	for _, a := range analyzers {
		as[a.NameID] = a
		choices = append(choices, a.NameID)
	}

	req := &model.SpecAnalysisRequest{
		Service: &model.Service{NameID: dummyServiceNameID},
		Spec: &model.Spec{
			ID:        dummySpecID,
			ServiceID: dummyServiceID,
			Doc:       model.NewSpecDoc(spec),
		},
	}

	analyzer = viper.GetString(flagAnalyzer)
	if len(analyzer) > 0 {
		a, found := as[analyzer]
		if !found {
			utils.ExitWithCode(utils.ExitInvalidInput, fmt.Errorf("invalid analyzer, choices: %s", strings.Join(choices, ", ")))
		}
		req.Analyzers = append(req.Analyzers, model.SpecAnalyzer(a.NameID))
	} else {
		for _, a := range analyzers {
			req.Analyzers = append(req.Analyzers, model.SpecAnalyzer(a.NameID))
		}
	}

	logDebugf("analyzing local spec: %s, analyzers: %v\n", filename, req.Analyzers)
	res, err := apiInsightsClient.AnalyzeAPISpec(cmd.Context(), req)
	if err != nil {
		utils.ExitWithCode(utils.ExitError, fmt.Errorf("failed to analyze spec: %s", err.Error()))
	}
	res.Print(os.Stdout, as)

	if res.SpecScore < failBelowScore {
		utils.ExitWithCode(utils.ExitFailBelowScore)
	}

	if failOnErrorRule && res.HasErrorFindings() {
		utils.ExitWithCode(utils.ExitErrorBlockerFindings)
	}

	utils.ExitWithCode(utils.ExitSuccess)
}

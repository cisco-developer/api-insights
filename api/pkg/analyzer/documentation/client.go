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

package documentationp

import (
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/cisco-developer/api-insights/api/pkg/analyzer/spectral"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

// documentationRuleset represents custom spectral ruleset for documentation.
var documentationRuleset = "node_modules/@cisco-developer/api-insights-openapi-rulesets/documentation.js"

func Flags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "documentation-ruleset",
			Usage:       "Documentation ruleset",
			Value:       documentationRuleset,
			Destination: &documentationRuleset,
			EnvVars:     []string{"DOCUMENTATION_RULESET"},
		}),
	}
}

func NewClient() (models.SpecDocAnalyzer, error) {
	sc, err := spectral.NewClient(documentationRuleset)
	if err != nil {
		return nil, err
	}
	return &cliClient{
		Client: sc,
	}, nil
}

// cliClient implements Linter.
type cliClient struct {
	*spectral.Client
}

func (c *cliClient) Analyze(doc models.SpecDoc, cfgMap analyzer.Config, serviceNameID *string) (*analyzer.Result, error) {
	return c.Client.Analyze(doc, cfgMap, serviceNameID)
}

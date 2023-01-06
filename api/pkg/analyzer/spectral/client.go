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

package spectral

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
)

// NewClient return a spectral client for some ruleset.
// ruleset is like "node_modules/@cisco-developer/api-insights-openapi-rulesets/api-insights-openapi-ruleset.js"
// Client is an instance of models.SpecDocAnalyzer
func NewClient(ruleset string) (*Client, error) {
	if err := preRunCheck(); err != nil {
		return nil, err
	}
	return &Client{
		ruleset: ruleset,
	}, nil
}

// Client implements Linter.
type Client struct {
	ruleset string
}

func (c *Client) Analyze(doc models.SpecDoc, cfgMap analyzer.Config, serviceNameID *string) (*analyzer.Result, error) {
	cfg := &analyzer.SpectralConfig{}
	if cfgMap != nil {
		if err := cfgMap.UnmarshalInto(cfg); err != nil {
			return nil, fmt.Errorf("analyzer: invalid config: %v", err)
		}
	}
	cfg.SetRuleset(c.ruleset)

	var (
		timestamp            = fmt.Sprintf("%v", time.Now().UnixNano())
		inputFilenamePattern = "lint-" + timestamp + "-in-*.json"
		outputFilename       = "/tmp/lint-" + timestamp + "-out.json"
	)

	inputFile, err := os.CreateTemp("/tmp", inputFilenamePattern)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = inputFile.Close()
		_ = os.Remove(inputFile.Name())
	}()
	if doc == nil || *doc == "" {
		return nil, fmt.Errorf("doc is nil or empty")
	}
	_, err = inputFile.Write([]byte(*doc))
	if err != nil {
		return nil, err
	}

	cmd := c.commandFromOpts(context.TODO(), cfg, inputFile.Name(), outputFilename)

	cmdString := cmd.String()
	shared.LogInfof("Running CMD[%s]...", cmdString)
	if out, err := cmd.CombinedOutput(); len(out) != 0 {
		shared.LogErrorf("Unexpected output returned by CMD[%s] (err=%v): %v", cmdString, err, string(out))
		//return nil, err
		// TODO cisco-openapi-ruleset uses an outdated version of @stoplight/spectral@5.9.2, which at times validates the document incorrectly.
		// Although it incorrectly validates the document & outputs the errors to the console, it still successfully lints & writes to the output file,
		// so we shouldn't blindly treat console output as a fatal error, for now. Need to revisit.
	}
	shared.LogInfof("Successfully ran CMD[%s].", cmdString)

	lintFileData, err := os.ReadFile(outputFilename)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = os.Remove(outputFilename)
	}()

	var result analyzer.SpectralResult
	err = json.Unmarshal(lintFileData, &result)
	if err != nil {
		return nil, err
	}
	shared.LogInfof("CMD[%s] resulted in %v results.", cmdString, len(result))
	if result == nil {
		shared.LogErrorf("CMD[%s] failed to read results from %v.", cmdString, string(lintFileData))
		return nil, fmt.Errorf("analyzer: failed to analyze")
	}

	return result.Result()
}

// commandFromOpts builds the `npx @cisco/cisco-openapi-ruleset ...` command.
func (c *Client) commandFromOpts(ctx context.Context, cfg *analyzer.SpectralConfig, inputFilename, outputFilename string) *exec.Cmd {

	if ctx == nil {
		ctx = context.Background()
	}

	var args []string

	args = append(args, "lint")

	args = append(args, "-f", "json")

	if cfg.Ruleset != nil && *cfg.Ruleset != "" {
		args = append(args, "-r", *cfg.Ruleset)
	}

	args = append(args, "-q")

	args = append(args, "-o", outputFilename)

	args = append(args, inputFilename)

	return exec.CommandContext(ctx, "spectral", args...)
}

func preRunCheck() error {
	// TODO Check dependencies like node & npm
	return nil
}

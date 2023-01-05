package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

var validationRuleset = "node_modules/@cisco-developer/api-insights-openapi-rulesets/validation.js"

// SpecValidationRequest represents payload for validating spec doc
type SpecValidationRequest struct {
	Doc SpecDoc `json:"doc"`
}

// SpecValidationResult represents spec validation result
type SpecValidationResult struct {
	Valid    bool                    `json:"valid"`
	Findings analyzer.SpectralResult `json:"findings"`
}

// ValidateSpecDoc checks if sd is a valid spec via spectral.
func ValidateSpecDoc(ctx context.Context, sd SpecDoc) (*SpecValidationResult, error) {
	if sd == nil || len(*sd) == 0 {
		return nil, errors.New("spec: empty")
	}

	var (
		timestamp            = fmt.Sprintf("%v", time.Now().UnixNano())
		inputFilenamePattern = "validation-" + timestamp + "-in-*.json"
		outputFilename       = "/tmp/validation-" + timestamp + "-out.json"
	)

	inputFile, err := os.CreateTemp("/tmp", inputFilenamePattern)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = inputFile.Close()
		_ = os.Remove(inputFile.Name())
	}()
	if sd == nil || *sd == "" {
		return nil, fmt.Errorf("doc is nil or empty")
	}

	_, err = inputFile.Write([]byte(*sd))
	if err != nil {
		return nil, err
	}

	cmd := exec.CommandContext(ctx, "spectral", "lint", "-f", "json", "-q", "-r", validationRuleset, "-o", outputFilename, inputFile.Name())

	if out, err := cmd.CombinedOutput(); len(out) != 0 {
		return nil, err
	}

	lintFileData, err := ioutil.ReadFile(outputFilename)
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
	if result == nil {
		return nil, fmt.Errorf("analyzer: failed to analyze")
	}

	vr := &SpecValidationResult{
		Valid:    len(result) == 0,
		Findings: result,
	}

	return vr, nil
}

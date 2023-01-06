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

// SpecAnalyzer represents the name of a spec analyzer.
type SpecAnalyzer string

const (
	CiscoAPIGuidelines = SpecAnalyzer("guidelines")
	InclusiveLanguage  = SpecAnalyzer("inclusive-language")
	Drift              = SpecAnalyzer("drift")
	Completeness       = SpecAnalyzer("completeness")
	Contract           = SpecAnalyzer("contract")
	Documentation      = SpecAnalyzer("documentation")
	Security           = SpecAnalyzer("security")
)

type Resulter interface{ Result() (*Result, error) }

var (
	_ Resulter = (*SpectralResult)(nil)
	_ Resulter = (*WokeResult)(nil)
	_ Resulter = (*APIClarityDriftResult)(nil)
)

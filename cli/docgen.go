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

package main

import (
	"github.com/cisco-developer/api-insights/cli/cmd"
	"github.com/spf13/cobra/doc"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir := "./docs"
	err := doc.GenMarkdownTree(cmd.GetRootCmd(), dir)
	if err != nil {
		log.Fatal(err)
	}

	// generate docs/README.md
	readme := filepath.Join(dir, "README.md")
	f, err := os.Create(readme)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = doc.GenMarkdownCustom(cmd.GetRootCmd(), f, func(s string) string {
		return s
	})
	if err != nil {
		log.Fatal(err)
	}
}

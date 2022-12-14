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

package shared

import (
	"os"
	"strings"
)

var defaultWhiteList = `
http://localhost:8080
http://localhost:8008`

func GetCORSWhitelist() []string {
	var whiteList []byte
	var err error
	filename := GetStringFromEnv(EnvCorsWhiteListFilePath, "/etc/cors/cors.list")
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		whiteList = []byte(defaultWhiteList)
	} else {
		whiteList, err = os.ReadFile(filename)
		if err != nil {
			LogErrorf("failed to read the cors whitelist file at %s", filename)
			panic(err)
		}
	}
	domains := strings.Split(string(whiteList), "\n")
	for _, domain := range domains {
		LogDebugf("adding following domain to CORS: %s", domain)
	}
	return domains
}

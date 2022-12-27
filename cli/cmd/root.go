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
	"fmt"
	"github.com/cisco-developer/api-insights/cli/pkg/client"
	"github.com/cisco-developer/api-insights/cli/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	flagHost               = "host"
	flagBasePath           = "base-path"
	flagHeader             = "header"
	flagAuthType           = "auth-type"
	flagAuthUsername       = "username"
	flagAuthPassword       = "password"
	flagAuthBearerToken    = "bearer-token"
	flagOAuth2GrantType    = "oauth2-grant-type"
	flagOAuth2TokenURL     = "oauth2-token-url"
	flagOAuth2ClientID     = "oauth2-client-id"
	flagOAuth2ClientSecret = "oauth2-client-secret"
	flagDebug              = "debug"
)

var debug bool
var cfgFile string
var apiInsightsClient client.APIInsightsClient
var host string
var basePath string
var rawHeaders []string
var (
	authType           string
	authUsername       string
	authPassword       string
	authBearerToken    string
	oAuth2GrantType    string
	oAuth2TokenURL     string
	oAuth2ClientID     string
	oAuth2ClientSecret string
)
var cliLongDescription = `The api-insights-cli tool is a CLI for API Insights.
Services are registered for analysis with the API Insights tool and API specs are associated with a service.
API Insights manages versioned API specifications (Swagger2/OpenAPI Spec 3) for services registered with API Insights.
The CLI validates and scores API specs against guidelines and also generates useful API changelogs across multiple versions.

Environment variables:

| Name                       | Description                                                   |
|----------------------------|---------------------------------------------------------------|
| $API_INSIGHTS_HOST                 | API host, such as https://host.example.com                    |
| $API_INSIGHTS_BASE_PATH            | API base path, such as /v1/registry                           |
| $API_INSIGHTS_AUTH_TYPE            | auth type, such as basic, bearer, oauth2                      |
| $API_INSIGHTS_USERNAME             | username for 'basic' auth-type                                |
| $API_INSIGHTS_PASSWORD             | password for 'basic' auth-type                                |
| $API_INSIGHTS_BEARER_TOKEN         | bearer token for 'bearer-token' auth-type                     |
| $API_INSIGHTS_OAUTH2_GRANT_TYPE    | grant type for 'oauth2' auth-type, such as client_credentials |
| $API_INSIGHTS_OAUTH2_TOKEN_URL     | token URL for 'oauth2' auth-type                              |
| $API_INSIGHTS_OAUTH2_CLIENT_ID     | client id for 'oauth2' auth-type                              |
| $API_INSIGHTS_OAUTH2_CLIENT_SECRET | client secret for 'oauth2' auth-type                          |
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "api-insights-cli",
	Short: "api-insights-cli is a CLI for API Insights.",
	Long:  cliLongDescription,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.LoadConfig()
		if len(host) > 0 {
			cfg.APIInsightsHost = host
		}
		if len(basePath) > 0 {
			cfg.APIInsightsBasePath = basePath
		}
		if len(rawHeaders) > 0 {
			cfg.Headers = make(map[string]string, len(rawHeaders))
			for _, rh := range rawHeaders {
				rhSubstrings := strings.SplitN(rh, ":", 2)
				if len(rhSubstrings) == 2 {
					cfg.Headers[strings.TrimSpace(rhSubstrings[0])] = strings.TrimSpace(rhSubstrings[1])
				}
			}
		}
		if cfg.AuthConfig == nil {
			cfg.AuthConfig = &config.AuthConfig{}
		}
		if len(authType) > 0 {
			cfg.AuthConfig.Type = authType
		}
		if len(authUsername) > 0 {
			cfg.AuthConfig.Username = authUsername
		}
		if len(authPassword) > 0 {
			cfg.AuthConfig.Password = authPassword
		}
		if len(authBearerToken) > 0 {
			cfg.AuthConfig.BearerToken = authBearerToken
		}
		if len(oAuth2GrantType) > 0 {
			cfg.AuthConfig.OAuth2GrantType = oAuth2GrantType
		}
		if len(oAuth2TokenURL) > 0 {
			cfg.AuthConfig.OAuth2TokenURL = oAuth2TokenURL
		}
		if len(oAuth2ClientID) > 0 {
			cfg.AuthConfig.OAuth2ClientID = oAuth2ClientID
		}
		if len(oAuth2ClientSecret) > 0 {
			cfg.AuthConfig.OAuth2ClientSecret = oAuth2ClientSecret
		}

		apiInsightsClient = client.NewAPIInsightsClient(cfg)
		return nil
	},
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.api-insights.yaml)")
	rootCmd.PersistentFlags().StringVarP(&host, flagHost, "H", "", "API host, for example: https://host.example.com")
	rootCmd.PersistentFlags().StringVarP(&basePath, flagBasePath, "", "", "API base path, for example: /v1/apiregistry")
	rootCmd.PersistentFlags().StringArrayVar(&rawHeaders, flagHeader, []string{}, "API header(s), for example: --header 'Content-Type: application/json' --header 'Accept: application/json'")
	rootCmd.PersistentFlags().BoolVarP(&debug, flagDebug, "", false, "verbose output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//
	// auth-specific flags
	//
	rootCmd.PersistentFlags().StringVarP(&authType, flagAuthType, "", "", "auth type, for example: basic, bearer, oauth2")
	// auth-type basic
	rootCmd.PersistentFlags().StringVarP(&authUsername, flagAuthUsername, "", "", "username for 'basic' auth-type")
	rootCmd.PersistentFlags().StringVarP(&authPassword, flagAuthPassword, "", "", "password for 'basic' auth-type")
	rootCmd.MarkFlagsRequiredTogether(flagAuthUsername, flagAuthPassword)
	// auth-type bearer-token
	rootCmd.PersistentFlags().StringVarP(&authBearerToken, flagAuthBearerToken, "", "", "bearer token for 'bearer-token' auth-type")
	// auth-type oauth2
	rootCmd.PersistentFlags().StringVarP(&oAuth2GrantType, flagOAuth2GrantType, "", "", "grant type for 'oauth2' auth-type, for example: client_credentials")
	rootCmd.PersistentFlags().StringVarP(&oAuth2TokenURL, flagOAuth2TokenURL, "", "", "token URL for 'oauth2' auth-type")
	rootCmd.PersistentFlags().StringVarP(&oAuth2ClientID, flagOAuth2ClientID, "", "", "client ID for 'oauth2' auth-type")
	rootCmd.PersistentFlags().StringVarP(&oAuth2ClientSecret, flagOAuth2ClientSecret, "", "", "client secret for 'oauth2' auth-type")
	rootCmd.MarkFlagsRequiredTogether(flagOAuth2GrantType, flagOAuth2TokenURL, flagOAuth2ClientID, flagOAuth2ClientSecret)

	// prevents Cobra from generating docs with "Auto generated by spf13/cobra..."
	rootCmd.DisableAutoGenTag = true

	// prevents Cobra from creating a default 'completion' command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := viper.BindPFlags(rootCmd.Flags())
	if err != nil {
		fmt.Println("failed to bind flags", err.Error())
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".api-insights" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".api-insights")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func logDebugln(a ...interface{}) {
	if debug {
		fmt.Println(a...)
	}
}

func logDebugf(format string, a ...interface{}) {
	if debug {
		fmt.Printf(format, a...)
	}
}

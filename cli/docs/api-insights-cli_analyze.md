## api-insights-cli analyze

Analyze local API spec

```
api-insights-cli analyze LOCAL_SPEC [flags]
```

### Examples

```
  # Analyze local spec with all analyzers
  api-insights-cli analyze testdata/carts.json

  # Analyze local spec with specific analyzer
  api-insights-cli analyze testdata/carts.json --analyzer guidelines
  api-insights-cli analyze testdata/carts.json --analyzer completeness
  api-insights-cli analyze testdata/carts.json --analyzer inclusive-language
  api-insights-cli analyze testdata/carts.json --analyzer drift
  api-insights-cli analyze testdata/carts.json --analyzer security
```

### Options

```
  -a, --analyzer string        API spec analyzer
      --fail-below-score int   Fail if API score is below specified score, defaults to 0
      --fail-on-error-rule     Fail if there are any error findings
  -h, --help                   help for analyze
```

### Options inherited from parent commands

```
      --auth-type string              auth type, for example: basic, bearer, oauth2
      --base-path string              API base path, for example: /v1/apiregistry
      --bearer-token string           bearer token for 'bearer-token' auth-type
      --config string                 config file (default is $HOME/.api-insights.yaml)
      --debug                         verbose output
      --header stringArray            API header(s), for example: --header 'Content-Type: application/json' --header 'Accept: application/json'
  -H, --host string                   API host, for example: https://host.example.com
      --oauth2-client-id string       client ID for 'oauth2' auth-type
      --oauth2-client-secret string   client secret for 'oauth2' auth-type
      --oauth2-grant-type string      grant type for 'oauth2' auth-type, for example: client_credentials
      --oauth2-token-url string       token URL for 'oauth2' auth-type
      --password string               password for 'basic' auth-type
      --username string               username for 'basic' auth-type
```

### SEE ALSO

* [api-insights-cli](api-insights-cli.md)	 - api-insights-cli is a CLI for API Insights.


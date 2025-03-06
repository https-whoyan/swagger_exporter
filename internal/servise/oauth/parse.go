package oauth

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type WebConfig struct {
	ClientID            string   `json:"client_id"`
	ProjectID           string   `json:"project_id"`
	AuthURI             string   `json:"auth_uri"`
	TokenURI            string   `json:"token_uri"`
	AuthProviderCertURL string   `json:"auth_provider_x509_cert_url"`
	ClientSecret        string   `json:"client_secret"`
	RedirectURIs        []string `json:"redirect_uris"`
}

type toGoogleConfig struct {
	Web *WebConfig `json:"web"`
}

func ParseOAuthConfig(cfg *WebConfig, applyURL string) (*oauth2.Config, error) {
	bytes, err := json.Marshal(toGoogleConfig{Web: cfg})
	if err != nil {
		return nil, err
	}
	return google.ConfigFromJSON(
		bytes,
		applyURL,
	)
}

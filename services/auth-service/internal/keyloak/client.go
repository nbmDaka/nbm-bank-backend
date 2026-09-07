package keycloak

import (
	"net/http"

	"github.com/nbmDaka/nbm-bank-backend/services/auth-service/config"
)

type Client struct {
	baseURL string

	realm string

	clientID string

	clientSecret string

	httpClient *http.Client
}

func NewClient(
	cfg config.KeycloakConfig,
) *Client {

	return &Client{

		baseURL: cfg.URL,

		realm: cfg.Realm,

		clientID: cfg.ClientID,

		clientSecret: cfg.ClientSecret,

		httpClient: &http.Client{},
	}
}

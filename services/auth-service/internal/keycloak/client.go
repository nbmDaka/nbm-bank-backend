package keycloak

import (
	"net/http"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"bytes"
	"io"
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

func (c *Client) GetAdminToken() (string, error) {


	tokenURL := fmt.Sprintf(
		"%s/realms/%s/protocol/openid-connect/token",
		c.baseURL,
		c.realm,
	)


	data := url.Values{}

	data.Set(
		"client_id",
		c.clientID,
	)


	data.Set(
		"client_secret",
		c.clientSecret,
	)


	data.Set(
		"grant_type",
		"client_credentials",
	)


	req, err := http.NewRequest(
		http.MethodPost,
		tokenURL,
		strings.NewReader(
			data.Encode(),
		),
	)


	if err != nil {
		return "", err
	}


	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)


	resp, err := c.httpClient.Do(req)


	if err != nil {
		return "", err
	}


	defer resp.Body.Close()



	if resp.StatusCode != http.StatusOK {

		return "",
			fmt.Errorf(
				"keycloak token request failed: %s",
				resp.Status,
			)
	}



	var token TokenResponse


	err = json.NewDecoder(
		resp.Body,
	).Decode(&token)


	if err != nil {
		return "", err
	}


	return token.AccessToken, nil
}

func (c *Client) CreateUser(
	token string,
	user CreateUserRequest,
) (string, error) {


	url := fmt.Sprintf(
		"%s/admin/realms/%s/users",
		c.baseURL,
		c.realm,
	)


	body, err := json.Marshal(user)

	if err != nil {
		return "", err
	}


	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(body),
	)


	if err != nil {
		return "", err
	}


	req.Header.Set(
		"Content-Type",
		"application/json",
	)


	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)


	resp, err := c.httpClient.Do(req)


	if err != nil {
		return "", err
	}


	defer resp.Body.Close()



	if resp.StatusCode != http.StatusCreated {

	body, _ := io.ReadAll(resp.Body)

	return "",
		fmt.Errorf(
			"create user failed: %s body=%s",
			resp.Status,
			string(body),
		)
}

	location := resp.Header.Get(
		"Location",
	)


	if location == "" {

		return "",
			fmt.Errorf(
				"missing location header",
			)
	}



	parts := strings.Split(
		location,
		"/",
	)


	userID := parts[len(parts)-1]


	return userID, nil
}
package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JWKS struct {
	Keys []JWK `json:"keys"`
}

type JWK struct {
	Kid string `json:"kid"`

	Kty string `json:"kty"`

	Alg string `json:"alg"`

	N string `json:"n"`

	E string `json:"e"`
}

func GetJWKS() (*JWKS, error) {

	resp, err := http.Get(
		"http://localhost:8080/realms/nbm-bank/protocol/openid-connect/certs",
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"failed to get jwks",
		)
	}

	var jwks JWKS

	err = json.NewDecoder(
		resp.Body,
	).Decode(&jwks)

	if err != nil {
		return nil, err
	}

	return &jwks, nil
}

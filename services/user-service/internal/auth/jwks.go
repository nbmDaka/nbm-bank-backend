package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
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

func decodeBase64URL(s string) ([]byte, error) {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		b, err = base64.URLEncoding.DecodeString(s)
	}
	return b, err
}

func (j *JWKS) GetPublicKey(kid string) (*rsa.PublicKey, error) {
	for _, key := range j.Keys {
		if key.Kid == kid {
			nBytes, err := decodeBase64URL(key.N)
			if err != nil {
				return nil, fmt.Errorf("failed to decode modulus: %w", err)
			}
			eBytes, err := decodeBase64URL(key.E)
			if err != nil {
				return nil, fmt.Errorf("failed to decode exponent: %w", err)
			}

			var eInt int
			for _, b := range eBytes {
				eInt = (eInt << 8) | int(b)
			}

			return &rsa.PublicKey{
				N: new(big.Int).SetBytes(nBytes),
				E: eInt,
			}, nil
		}
	}
	return nil, fmt.Errorf("public key not found for kid: %s", kid)
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


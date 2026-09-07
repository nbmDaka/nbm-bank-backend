package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email string `json:"email"`

	PreferredUsername string `json:"preferred_username"`

	Azp string `json:"azp"`

	jwt.RegisteredClaims
}

func ValidateToken(
	tokenString string,
	jwks *JWKS,
) (*jwt.Token, error) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {

			if token.Method.Alg() != "RS256" {
				return nil, fmt.Errorf(
				"unexpected signing algorithm",
			)
}

			if jwks == nil {
				var err error
				jwks, err = GetJWKS()
				if err != nil {
					return nil, fmt.Errorf("failed to fetch jwks: %w", err)
				}
			}

			kid, ok := token.Header["kid"].(string)
			if !ok {
				return nil, fmt.Errorf("missing kid in token header")
			}

			return jwks.GetPublicKey(kid)
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	issuer := claims.Issuer

	if issuer != "http://localhost:8080/realms/nbm-bank" {
		return nil, fmt.Errorf(
			"invalid issuer",
		)
	}

	if claims.Azp != "nbm-backend" {
		return nil, fmt.Errorf(
			"invalid client",
		)
	}

	return token, nil
}

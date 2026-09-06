package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)


func ValidateToken(
	tokenString string,
) (*jwt.Token,error){


	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token)(interface{},error){

			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf(
					"unexpected signing method",
				)
			}


			return nil,nil
		},
	)


	if err != nil {
		return nil,err
	}


	return token,nil
}
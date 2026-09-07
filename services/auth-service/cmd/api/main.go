package main

import (
	"log"
	"net/http"

	"github.com/nbmDaka/nbm-bank-backend/services/auth-service/config"
	"github.com/nbmDaka/nbm-bank-backend/services/auth-service/internal/keycloak"
)

func main() {

	err := config.LoadEnv()

	if err != nil {
		log.Println(
			"No .env file found",
		)
	}

	cfg := config.Load()

	kc := keycloak.NewClient(
		cfg.Keycloak,
	)

	token, err := kc.GetAdminToken()

	if err != nil {
		log.Fatal(err)
	}
	
	err = kc.CreateUser(
		token,
		keycloak.CreateUserRequest{

			Username: "john",

			Email: "john@test.com",

			FirstName: "John",

			LastName: "Smith",

			Enabled: true,


			Credentials: []keycloak.Credential{

				{
					Type: "password",
					Value: "Password123!",
					Temporary: false,
				},
			},
		},
	)


if err != nil {
	log.Fatal(err)
}


log.Println("user created")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("auth-service is running"))
	})

	log.Println(
		"auth-service starting on port",
		cfg.App.Port,
	)

	err = http.ListenAndServe(
		":"+cfg.App.Port,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}
}

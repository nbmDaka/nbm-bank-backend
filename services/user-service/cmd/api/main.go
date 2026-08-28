package main

import (
	"log"

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/config"
)


func main(){

	err := config.LoadEnv()

	if err != nil {
		log.Println(
		"No .env file found",
		)
	}


	cfg := config.Load()


	log.Println(
		"Environment:",
		cfg.App.Environment,
	)
	
	log.Println(
		"Port:",
		cfg.App.Port,
	)

}
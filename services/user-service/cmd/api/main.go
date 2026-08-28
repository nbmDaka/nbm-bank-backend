package main

import (
	"log"

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/config"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/app"
)

func main() {

	err := config.LoadEnv()

	if err != nil {
		log.Println(
			"No .env file found",
		)
	}

	cfg := config.Load()

	application, err := app.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	err = application.Start()

	if err != nil {
		log.Fatal(err)
	}

	for {
		// Do nothing, just keep the main goroutine alive
	}
}

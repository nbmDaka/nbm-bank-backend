package main

import (
	"log"

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/config"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/platform/database"
)

func main() {

	err := config.LoadEnv()

	if err != nil {
		log.Println(
			"No .env file found",
		)
	}

	cfg := config.Load()


	db, err := database.NewPostgres(
		cfg.Database,
	)

	if err != nil {
		log.Fatal(
			"database connection failed:",
			err,
		)
	}
	defer db.Close()
}

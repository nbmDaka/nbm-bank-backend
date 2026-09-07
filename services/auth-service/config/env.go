package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {

	err := godotenv.Load()

	if err != nil {
		return err
	}

	return nil
}


func getEnv(
	key string,
	defaultValue string,
) string {

	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
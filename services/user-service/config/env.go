package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {

	err := godotenv.Load()

	if err != nil {
		return nil
	}

	return nil
}

func GetEnv(key string, fallback string) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
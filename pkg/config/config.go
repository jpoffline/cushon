package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser   string
	PostgresPass   string
	PostgresHost   string
	PostgresPort   string
	PostgresDBName string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}
	return &Config{
		PostgresUser:   os.Getenv("POSTGRES_USER"),
		PostgresPass:   os.Getenv("POSTGRES_PASS"),
		PostgresHost:   os.Getenv("POSTGRES_HOST"),
		PostgresPort:   os.Getenv("POSTGRES_PORT"),
		PostgresDBName: os.Getenv("POSTGRES_DB_NAME"),
	}, nil
}

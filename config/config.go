package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var config *Config

const (
	envPath = ".env"
	port    = "2024"
)

type Config struct {
	DbUrl string
	Port  string
}

func LoadConfig() (*Config, error) {
	errEnv := godotenv.Load(envPath)

	if errEnv != nil {
		return nil, errors.New("unable to load the environment file")
	}

	loadedPort := os.Getenv("PORT")
	loadedDbUrl := os.Getenv("DB_URL")

	if loadedDbUrl == "" {
		return nil, errors.New("unable to load db url")
	}

	if loadedPort == "" {
		log.Printf("Set default port %v", port)
		loadedPort = port
	}

	config = &Config{
		loadedDbUrl,
		loadedPort,
	}

	return config, nil
}

package config

import (
	"fmt"
	"log"
	"os"
)

var config *Config

const (
	port = "2024"
)

type Config struct {
	DbUrl string
	Port  string
}

func LoadConfig() (*Config, error) {
	loadedPort := os.Getenv("PORT")

	loadedDbUrl := fmt.Sprintf(
		"%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

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

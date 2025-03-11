package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type Config struct {
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load env file!")
	}

	pgHost := failOnEmpty("POSTGRES_HOST", os.Getenv("POSTGRES_HOST"))
	pgUser := failOnEmpty("POSTGRES_USER", os.Getenv("POSTGRES_USER"))
	pgPassword := failOnEmpty("POSTGRES_PASSWORD", os.Getenv("POSTGRES_PASSWORD"))
	pgDb := failOnEmpty("POSTGRES_DB", os.Getenv("POSTGRES_DB"))

	return &Config{
		PostgresHost:     pgHost,
		PostgresUser:     pgUser,
		PostgresPassword: pgPassword,
		PostgresDb:       pgDb,
	}

}

func failOnEmpty(key string, value string) string {
	if strings.TrimSpace(value) == "" {
		log.Fatal("Env variable is empty: " + key)
		return ""
	}
	return value
}

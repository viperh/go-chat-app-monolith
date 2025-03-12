package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
	PostgresPort     string

	JwtSecret string
	ApiPort   string
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
	pgPort := failOnEmpty("POSTGRES_PORT", os.Getenv("POSTGRES_PORT"))

	jwtSecret := failOnEmpty("JWT_SECRET", os.Getenv("JWT_SECRET"))
	apiPort := failOnEmpty("API_PORT", os.Getenv("API_PORT"))

	return &Config{
		PostgresHost:     pgHost,
		PostgresUser:     pgUser,
		PostgresPassword: pgPassword,
		PostgresDb:       pgDb,
		JwtSecret:        jwtSecret,
		ApiPort:          apiPort,
		PostgresPort:     pgPort,
	}

}

func failOnEmpty(key string, value string) string {
	if strings.TrimSpace(value) == "" {
		log.Fatal("Env variable is empty: " + key)
		return ""
	}
	return value
}

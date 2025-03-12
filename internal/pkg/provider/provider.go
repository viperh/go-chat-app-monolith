package provider

import (
	"fmt"
	"go-chat-app-monolith/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Provider struct {
	Db *gorm.DB
}

func NewProvider(cfg *config.Config) *Provider {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDb, cfg.PostgresPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &Provider{Db: db}
}

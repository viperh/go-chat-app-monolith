package migrations

import (
	"fmt"
	"go-chat-app-monolith/internal/config"
	"go-chat-app-monolith/internal/models"
	"gorm.io/driver/postgres"
	"log"

	"gorm.io/gorm"
)

type Migration struct {
	Db *gorm.DB
}

func NewMigration(cfg *config.Config) *Migration {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDb, cfg.PostgresPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database!")
	}

	return &Migration{
		Db: db,
	}
}

func (m *Migration) Up() {
	err := m.Db.AutoMigrate(&models.User{}, &models.Message{}, &models.Room{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migration completed!")

}

func (m *Migration) Down() {
	err := m.Db.Migrator().DropTable(&models.User{}, &models.Message{}, &models.Room{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migration completed!")
}

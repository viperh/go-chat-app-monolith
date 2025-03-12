package main

import (
	"flag"
	"go-chat-app-monolith/internal/config"
	"go-chat-app-monolith/internal/migrations"
	"log"
)

func main() {
	cfg := config.MustLoad()
	mgr := migrations.NewMigration(cfg)

	action := flag.String("action", "up", "Migration action")

	flag.Parse()

	switch *action {
	case "up":
		mgr.Up()
	case "down":
		mgr.Down()
	default:
		log.Fatal("Invalid action!")
	}

}

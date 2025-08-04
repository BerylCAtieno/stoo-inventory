package migration

import (
	"log"

	"github.com/berylCAtieno/stoo-inventory/internal/database"
	"github.com/berylCAtieno/stoo-inventory/internal/models"
)

func Run() {
	err := database.DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migration complete.")
}

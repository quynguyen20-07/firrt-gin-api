package database

import (
	"log"

	"github.com/gin/api/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("✅ Migration completed successfully.")
}

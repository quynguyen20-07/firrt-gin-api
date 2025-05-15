package database

import (
	"log"

	"github.com/gin/api/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gin-api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
}

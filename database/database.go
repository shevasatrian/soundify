package database

import (
	"log"
	"os"
	"playlist-music/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Music{}, &models.Playlist{}, &models.PlaylistMusic{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database connection established and migrated")
}

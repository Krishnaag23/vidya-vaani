package storage

import (
	"log"

	"github.com/krishnaag23/vidya-vaani/internal/storage/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitPostgres() {
	var err error
	dsn := "host=db user=youruser password=yourpassword dbname=vidya sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	log.Println("Connected to PostgreSQL")

	err = db.AutoMigrate(&models.Student{}, &models.LogEntry{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

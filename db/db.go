package db

import (
	"log"
	"os"

	"inventory-supply-chain-system/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Retrieve database credentials from environment variables
	dsn := os.Getenv("DB_DSN")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Auto-migrate models
	err = DB.AutoMigrate(&models.User{}, &models.Inventory{}, &models.Order{}, &models.Shipment{}, &models.Vendor{})
	if err != nil {
		log.Fatalf("Error with auto-migration: %v", err)
	}

	log.Println("Connected to the database and applied migrations successfully!")
}

package db

import (
	"fmt"
	"log"
	"os"

	"inventory-supply-chain-system/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Build the DSN from environment variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Auto-migrate models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Inventory{},
		&models.Order{},
		&models.Shipment{},
		&models.Vendor{},
	)
	if err != nil {
		log.Fatalf("Error with auto-migration: %v", err)
	}

	log.Println("Connected to the database and applied migrations successfully!")
}

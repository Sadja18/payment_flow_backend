package config

import (
	"log"

	"github.com/Sadja18/payment-flow-backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("payments.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate tables
	err = db.AutoMigrate(&models.User{}, &models.Invoice{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Connected to DB and ran migrations")
	return db
}

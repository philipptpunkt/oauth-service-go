package utils

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatalf("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB = db

	log.Println("Database connection established using GORM")
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database is not initialized")
	}
	return DB
}

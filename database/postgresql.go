package database

import (
	"context"
	"database/sql"
	"go_postgtresql_pgx/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
		os.Exit(1)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	// using pgx as the undeerlying driver for the standard database/sql package
	sqlDB, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}
	// make sure the database is reachable
	err = sqlDB.PingContext(context.Background())
	if err != nil {
		log.Fatalf("Failed to ping DB: %v", err)

	}
	// Initialize GORM with the sql.DB object
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize GORM: %v", err)
	}
	log.Println("Connected to database using GORM and pgx")
	// Migrate the schema
	gormDB.AutoMigrate(&models.Person{})
	DB = gormDB
}

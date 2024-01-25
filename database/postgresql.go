package database

import (
	"go_postgtresql_pgx/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
		return nil, err
	}

	databaseURL := os.Getenv("DATABASE_URL")
	// using pgx as the undeerlying driver for the standard database/sql package
	// sqlDB, err := sql.Open("pgx", databaseURL)
	// if err != nil {
	// 	log.Fatalf("Failed to open a DB connection: %v", err)
	// 	return nil,err
	// }
	// // make sure the database is reachable
	// err = sqlDB.PingContext(context.Background())
	// if err != nil {
	// 	log.Fatalf("Failed to ping DB: %v", err)
	// 	return nil,err

	// }
	// Initialize GORM with the sql.DB object
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: databaseURL,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize GORM: %v", err)
		return nil, err
	}
	log.Println("Connected to database using GORM and pgx")
	// Migrate the schema
	gormDB.AutoMigrate(&models.Person{})
	return gormDB, nil
}

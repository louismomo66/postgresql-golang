package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() *pgxpool.Pool {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
		os.Exit(1)
	}
	databasUrl := os.Getenv("DATABASE_URL")
	dbpool, err := pgxpool.New(context.Background(), databasUrl)
	if err != nil {
		log.Fatalf("Failed to connect to db")
	}
	log.Printf("Connected to database")
	return dbpool
}

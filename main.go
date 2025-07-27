package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/joho/godotenv"
	"github.com/jubilant-gremlin/growlog/internal/config"
	"github.com/jubilant-gremlin/growlog/internal/database"
	"github.com/rivo/tview"
)

func readEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable not set.\n", key)
	}
	return value
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s\n", err)
	}

	cfg := config.Config{
		DbURL: readEnv("DB_URL"),
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)
}

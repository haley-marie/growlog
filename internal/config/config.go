package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/jubilant-gremlin/growlog/internal/database"
	_ "github.com/lib/pq"
)

type Config struct {
	DbQueries *database.Queries
}

func readEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable not set.\n", key)
	}
	return value
}

func ReadCfg() Config {
	DbURL := readEnv("DB_URL")

	db, err := sql.Open("postgres", DbURL)
	if err != nil {
		log.Fatalf("Error opening database connection: %s\n", err)
	}

	dbQueries := database.New(db)

	cfg := Config{
		DbQueries: dbQueries,
	}

	return cfg
}

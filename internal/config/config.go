package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/jubilant-gremlin/growlog/internal/database"
	migrate "github.com/jubilant-gremlin/growlog/internal/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
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

	goose.SetBaseFS(migrate.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Error setting goose dialect: %v", err)
	}

	if err := goose.Up(db, "schema"); err != nil {
		log.Fatalf("Error running goose migrations: %v", err)
	}

	return cfg
}

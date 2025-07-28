package commands

import (
	"context"
	"log"

	"github.com/jubilant-gremlin/growlog/internal"
	"github.com/jubilant-gremlin/growlog/internal/config"
)

func HandlerResetPlants(cfg *config.Config, cmd Command) error {
	err := cfg.DbQueries.ResetPlants(context.Background())
	if err != nil {
		return internal.ReturnFormattedError("Error resetting plant database", err)
	}

	log.Println("Successfully reset plant database.")
	return nil
}

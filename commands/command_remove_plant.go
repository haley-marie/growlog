package commands

import (
	"context"
	"log"
	"strconv"

	"github.com/jubilant-gremlin/growlog/internal"
	"github.com/jubilant-gremlin/growlog/internal/config"
)

func HandlerRemovePlant(cfg *config.Config, cmd Command) error {
	if len(cmd.Arguments) < 1 {
		return internal.ReturnMessageAsError("Error removing plant: must specify plant id to remove")
	}

	if len(cmd.Arguments) > 1 {
		return internal.ReturnMessageAsError("Error removing plant: may only specify one plant id to remove")
	}

	id, err := strconv.Atoi(cmd.Arguments[0])
	if err != nil {
		return internal.ReturnFormattedError("Error converting plantId to int", err)
	}
	err = cfg.DbQueries.RemovePlantById(context.Background(), int32(id))
	if err != nil {
		return internal.ReturnFormattedError("Error removing plant", err)
	}

	log.Printf("Successfully removed plant with id %d", id)
	return nil
}

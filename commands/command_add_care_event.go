package commands

import (
	"context"
	"log"

	"github.com/jubilant-gremlin/growlog/internal"
	"github.com/jubilant-gremlin/growlog/internal/config"
)

func HandlerAddCustomCareEvent(cfg *config.Config, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return internal.ReturnMessageAsError("Error creating custom care event: must specify event name")
	}

	event, err := cfg.DbQueries.CreateCareEvent(context.Background(), cmd.Arguments[0])
	if err != nil {
		return internal.ReturnFormattedError("Error creating custom care event", err)
	}

	log.Printf("Successfully added custom care event %s to care events", event.TaskName)
	return nil
}

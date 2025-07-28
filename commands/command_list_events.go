package commands

import (
	"context"
	"fmt"

	"github.com/jubilant-gremlin/growlog/internal"
	"github.com/jubilant-gremlin/growlog/internal/config"
)

func HandlerListAllEvents(cfg *config.Config, cmd Command) error {
	events, err := cfg.DbQueries.ListCareEvents(context.Background())
	if err != nil {
		return internal.ReturnFormattedError("Error getting plants from database", err)
	}

	fmt.Println("Event ID, Event Name")
	for _, event := range events {
		fmt.Printf("%d, %s\n", event.ID, event.TaskName)
	}

	return nil
}

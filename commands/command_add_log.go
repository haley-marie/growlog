package commands

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/jubilant-gremlin/growlog/internal"
	"github.com/jubilant-gremlin/growlog/internal/config"
	"github.com/jubilant-gremlin/growlog/internal/database"
)

func HandlerAddLog(cfg *config.Config, cmd Command) error {
	if len(cmd.Arguments) < 2 {
		return internal.ReturnMessageAsError("Error: must specify plant and event ids to add to growlog")
	}

	plantID, err := strconv.Atoi(cmd.Arguments[0])
	if err != nil {
		return internal.ReturnFormattedError("Error converting plant ID to int", err)
	}

	eventID, err := strconv.Atoi(cmd.Arguments[1])
	if err != nil {
		return internal.ReturnFormattedError("Error converting event ID to int", err)
	}

	timeOfCare := sql.NullTime{Valid: false}

	if len(cmd.Arguments) == 3 {
		careTime, err := time.Parse(time.RFC822, cmd.Arguments[2])
		if err != nil {
			return internal.ReturnFormattedError("Error adding log entry", err)
		}

		timeOfCare = sql.NullTime{Time: careTime, Valid: true}
	}

	logEntry, err := cfg.DbQueries.CreateLogEntry(
		context.Background(),
		database.CreateLogEntryParams{ID: int32(plantID), ID_2: int32(eventID), TimeOfCare: timeOfCare},
	)

	if err != nil {
		return internal.ReturnFormattedError("Error adding log entry", err)
	}

	log.Printf("Successfully added log: %v", logEntry)
	return nil
}

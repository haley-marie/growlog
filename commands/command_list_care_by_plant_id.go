package commands

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jubilant-gremlin/growlog/internal"
	"github.com/jubilant-gremlin/growlog/internal/config"
)

func HandlerGetLogsForPlant(cfg *config.Config, cmd Command) error {
	if len(cmd.Arguments) < 1 {
		return internal.ReturnMessageAsError("Must specify plant id.")
	}

	plantID := cmd.Arguments[0]
	plantIDInt, err := strconv.Atoi(plantID)
	if err != nil {
		return internal.ReturnFormattedError("Error converting plant ID to integer", err)
	}

	logs, err := cfg.DbQueries.GetLogsForPlant(
		context.Background(),
		sql.NullInt64{
			Int64: int64(plantIDInt),
			Valid: true,
		})
	if err != nil {
		return internal.ReturnFormattedError("Error getting logs for plant", err)
	}

	fmt.Println("Plant ID, Plant Type, Event Name, Time of Care")
	for _, log := range logs {
		timeOfCare := "---"
		if log.TimeOfCare.Valid == true {
			timeOfCare = log.TimeOfCare.Time.String()
		}
		fmt.Printf("%v, %v, %s, %v\n", log.PlantID.Int64, log.PlantType.String, log.CareEventName.String, timeOfCare)
	}

	return nil
}

package commands

import (
	"context"
	"fmt"

	"github.com/jubilant-gremlin/growlog/internal"
	"github.com/jubilant-gremlin/growlog/internal/config"
)

func HandlerListAllLogs(cfg *config.Config, cmd Command) error {
	logs, err := cfg.DbQueries.ListAllLogs(context.Background())
	if err != nil {
		return internal.ReturnFormattedError("Error getting logs from database", err)
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

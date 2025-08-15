package commands

import (
	"context"
	"fmt"

	"github.com/jubilant-gremlin/growlog/internal"
	"github.com/jubilant-gremlin/growlog/internal/config"
)

func HandlerListPlantsByType(cfg *config.Config, cmd Command) error {
	if len(cmd.Arguments) < 1 {
		return internal.ReturnMessageAsError("Plant type must be specified.")
	}

	plantType := cmd.Arguments[0]
	plants, err := cfg.DbQueries.GetPlantsByType(context.Background(), plantType)
	if err != nil {
		return internal.ReturnFormattedError("Error getting plants from database", err)
	}
	for _, plant := range plants {
		plantedAt := "---"
		if plant.PlantedAt.Valid == true {
			plantedAt = plant.PlantedAt.Time.String()
		}
		fmt.Printf("%d, %s, %v\n", plant.ID, plant.PlantType, plantedAt)
	}

	return nil
}

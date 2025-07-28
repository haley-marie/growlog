package commands

import (
	"context"
	"fmt"

	"github.com/jubilant-gremlin/growlog/internal"
	"github.com/jubilant-gremlin/growlog/internal/config"
)

func HandlerListAllPlants(cfg *config.Config, cmd Command) error {
	plants, err := cfg.DbQueries.GetAllPlants(context.Background())
	if err != nil {
		return internal.ReturnFormattedError("Error getting plants from database", err)
	}

	fmt.Println("Plant ID, Plant Type, Time Planted")
	for _, plant := range plants {
		plantedAt := "---"
		if plant.PlantedAt.Valid == true {
			plantedAt = plant.PlantedAt.Time.String()
		}
		fmt.Printf("%d, %s, %v\n", plant.ID, plant.PlantType, plantedAt)
	}

	return nil
}

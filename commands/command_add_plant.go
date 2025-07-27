package commands

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jubilant-gremlin/growlog/internal"
	"github.com/jubilant-gremlin/growlog/internal/config"
	"github.com/jubilant-gremlin/growlog/internal/database"
)

func HandlerAddPlant(cfg *config.Config, cmd Command) error {
	plantType := "unspecified"
	plantedAt := sql.NullTime{Valid: false}

	if len(cmd.Arguments) < 1 {
		return internal.ReturnMessageAsError("Error adding plant: Plant type must be specified")
	} else if len(cmd.Arguments) == 1 {
		plantType = cmd.Arguments[0]
	} else {
		plantType = cmd.Arguments[0]
		plantedTime, err := time.Parse(time.RFC822, cmd.Arguments[1])
		if err != nil {
			return internal.ReturnFormattedError("Error adding plant", err)
		}
		plantedAt = sql.NullTime{Time: plantedTime, Valid: true}
	}

	plant, err := cfg.DbQueries.AddPlant(
		context.Background(),
		database.AddPlantParams{
			PlantType: plantType,
			PlantedAt: plantedAt,
		},
	)

	if err != nil {
		return internal.ReturnFormattedError("Error adding plant", err)
	}

	log.Printf("Successfully added %s to Grow Log. Happy Planting!", plant.PlantType)
	return nil
}

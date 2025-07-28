package main

import (
	"log"
	"os"

	//	"github.com/gdamore/tcell/v2"
	"github.com/joho/godotenv"
	"github.com/jubilant-gremlin/growlog/commands"
	"github.com/jubilant-gremlin/growlog/internal/config"
	// "github.com/rivo/tview"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s\n", err)
	}

	cfg := config.ReadCfg()

	cmd_map := commands.Commands{
		Cmds: make(map[string]func(*config.Config, commands.Command) error),
	}

	registerHandlers(&cmd_map)

	args := os.Args
	new_cmd := commands.Command{
		Name:      args[1],
		Arguments: args[2:],
	}

	err = cmd_map.Run(&cfg, new_cmd)
	if err != nil {
		log.Fatalf("Error running command: %s\n", err)
	}

}

func registerHandlers(cmd_map *commands.Commands) {
	cmd_map.Register("addPlant", commands.HandlerAddPlant)
	cmd_map.Register("resetPlants", commands.HandlerResetPlants)
	cmd_map.Register("removePlant", commands.HandlerRemovePlant)
	cmd_map.Register("addEvent", commands.HandlerAddCustomCareEvent)
	cmd_map.Register("addLog", commands.HandlerAddLog)
}

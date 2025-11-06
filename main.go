package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"github.com/jubilant-gremlin/growlog/commands"
	"github.com/jubilant-gremlin/growlog/internal/config"
	"github.com/jubilant-gremlin/growlog/internal/tui"
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

	p := tea.NewProgram(tui.NewModel(&cfg))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Sorry, an error occured: %v", err)
		os.Exit(1)
	}
}

func registerHandlers(cmd_map *commands.Commands) {
	cmd_map.Register("addPlant", commands.HandlerAddPlant)
	cmd_map.Register("resetPlants", commands.HandlerResetPlants)
	cmd_map.Register("removePlant", commands.HandlerRemovePlant)
	cmd_map.Register("addEvent", commands.HandlerAddCustomCareEvent)
	cmd_map.Register("addLog", commands.HandlerAddLog)
	cmd_map.Register("listPlants", commands.HandlerListAllPlants)
	cmd_map.Register("listEvents", commands.HandlerListAllEvents)
	cmd_map.Register("listLogs", commands.HandlerListAllLogs)
	cmd_map.Register("listPlantsByType", commands.HandlerListPlantsByType)
	cmd_map.Register("listLogsForPlant", commands.HandlerGetLogsForPlant)
}

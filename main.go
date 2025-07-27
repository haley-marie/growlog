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

	cmd_map.Register("addPlant", commands.HandlerAddPlant)

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

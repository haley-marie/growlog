package commands

import "github.com/jubilant-gremlin/growlog/internal/config"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmds map[string]func(*config.Config, command) error
}

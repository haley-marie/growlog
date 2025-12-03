package commands

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/jubilant-gremlin/growlog/internal"
	"github.com/jubilant-gremlin/growlog/internal/config"
)

type Command struct {
	Name      string
	Arguments []textinput.Model
}

type Commands struct {
	Cmds map[string]func(*config.Config, Command) error
}

func (c *Commands) Register(name string, f func(*config.Config, Command) error) error {
	c.Cmds[name] = f

	_, ok := c.Cmds[name]

	if !ok {
		return internal.ReturnMessageAsError("Error registering command")
	}

	return nil
}

func (c *Commands) Run(cfg *config.Config, cmd Command) error {
	handlerName := cmd.Name

	handler, ok := c.Cmds[handlerName]
	if !ok {
		return internal.ReturnMessageAsError("Error running command: Command not found. Try growlog help for a list of valid commands")
	}

	handler(cfg, cmd)
	return nil
}

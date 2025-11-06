package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jubilant-gremlin/growlog/internal/config"
)

type menuItem struct {
	title        string
	hasChildMenu bool
	hasCommand   bool
	childMenu    []menuItem
	command      string
}

type model struct {
	choices []menuItem
	cursor  int
	cfg     *config.Config
}

var mainMenu = []menuItem{
	{
		title:        "Quick Add",
		hasChildMenu: true,
		hasCommand:   false,
		childMenu:    quickAddMenu,
	},
	{
		title:        "Quick View",
		hasChildMenu: true,
		hasCommand:   false,
		childMenu:    quickViewMenu,
	},
	{
		title:        "Logs",
		hasChildMenu: true,
		hasCommand:   false,
		childMenu:    logsMenu,
	},
	{
		title:        "Settings",
		hasChildMenu: true,
		hasCommand:   false,
		childMenu:    settingsMenu,
	},
}

var quickAddMenu = []menuItem{
	{
		title:        "Add Plant",
		hasChildMenu: false,
		hasCommand:   true,
		command:      "addPlant",
	},
	{
		title:        "Add Event",
		hasChildMenu: false,
		hasCommand:   true,
		command:      "addEvent",
	},
	{
		title:        "Add Log",
		hasChildMenu: false,
		hasCommand:   true,
		command:      "addLog",
	},
}

var quickViewMenu []menuItem
var logsMenu []menuItem
var settingsMenu []menuItem

func NewModel(cfg *config.Config) model {
	return model{choices: mainMenu, cfg: cfg}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	// header
	s := "Welcome to Grow Log! What would you like to do?\n\n"

	// body
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice.title)
	}

	// footer
	s += "\nPress q to quit"

	return s
}

package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string // menu options
	cursor   int      // menu item cursor is pointing at
	selected int      // menu item selected
}

var InitialModel = model{
	choices: []string{"Plants", "Logs", "Settings"},
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl.c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.selected = m.cursor
		}
	}
	return m, nil
}

func (m model) View() string {
	// header
	s := "Welcome to grow log! What would you like to do?\n\n"

	// body
	for i, choice := range m.choices {

		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// footer
	s += "\nPress q or ctrl+c to quit.\n"

	return s
}

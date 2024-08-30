package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the state of the TUI.
type Model struct {
	count int
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

// InitialModel returns the initial state of the model.
func InitialModel() Model {
	return Model{
		count: 0,
	}
}

// Update is called to handle user input and update the model's state.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			m.count++
		case "down":
			m.count--
		}
	}
	return m, nil
}

// View returns the view that should be displayed.
func (m Model) View() string {
	return fmt.Sprintf("Press q to quit.\nCount: %d\nPress up/down to change the count.", m.count)
}

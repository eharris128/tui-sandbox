package models

import (
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the state of the TUI.
type Model struct {
	table   table.Model
	loaded  bool
	message string
}

// Image represents an individual image in the report
type Image struct {
	Cmd     string    `json:"cmd"`
	Created time.Time `json:"created"`
	ID      string    `json:"id"`
	Info    string    `json:"info"`
	Name    string    `json:"name"`
	Size    string    `json:"size"`
}

// ImageReport represents the report containing multiple images
type ImageReport struct {
	Images []Image `json:"images"`
}

// InitialModel returns the initial state of the model.
func InitialModel(data ImageReport) *Model {
	var rows []table.Row
	for _, imageInfo := range data.Images {
		// Should shorten the imageInfo.ID.
		imageRow := []string{imageInfo.Name, imageInfo.ID, imageInfo.Created.String(), imageInfo.Size}
		rows = append(rows, imageRow)
	}
	m := &Model{}
	columns := []table.Column{
		{Title: "Name", Width: 30},
		{Title: "Image ID", Width: 30},
		{Title: "Created", Width: 30},
		{Title: "Size", Width: 8},
	}
	table := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	m.table = table
	return m
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

// Update is called to handle user input and update the model's state.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit // Quit the program.
		case "z":
			m.message = "Handle error message displays here"
			return m, nil
		}
	}
	return m, nil
}

// View returns the view that should be displayed.
func (m Model) View() string {
	var statusBar string

	// Display messages to the user
	if m.message != "" {
		// Second argument should be informed based on terminal width
		statusBar = RightPadTrim(m.message, 1000)
	}

	if m.loaded {
		content := m.table.View()

		footerStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#282828")).
			Background(lipgloss.Color("#7c6f64"))

		footerStr := "Press q to quit"
		footer := footerStyle.Render(footerStr)
		return lipgloss.JoinVertical(lipgloss.Left,
			content,
			statusBar,
			footer,
		)
	} else {
		return "loading"
	}
}

package models

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

// Styling
var (
	columnStyle = lipgloss.NewStyle().Padding(1, 2)
)

// Model represents the state of the TUI.
type Model struct {
	list list.Model

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

type item struct {
	name string
	size string
}

func (i item) Title() string {
	return i.name
}

func (i item) Description() string {
	return fmt.Sprintf("Size: %v", i.size)
}

func (i item) FilterValue() string { return i.name }

// InitialModel returns the initial state of the model.
func InitialModel(data ImageReport) *Model {
	var items []list.Item
	for _, imageInfo := range data.Images {
		image := item{
			name: imageInfo.Name,
			size: imageInfo.Size,
		}
		items = append(items, image)
	}
	m := &Model{}
	m.initList(90, 32)
	m.list.SetItems(items)
	return m
}

func (m *Model) initList(width, height int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.list.Title = "Image Viewer"
	m.list.SetStatusBarItemName("image", "images")
	m.list.SetShowHelp(false)
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
			columnStyle.Width(msg.Width)
			columnStyle.Height(int(float64(msg.Height) * .5))
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
		imagesView := m.list.View()
		content := lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Render(imagesView),
		)
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

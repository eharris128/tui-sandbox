package models

import (
	"log"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cast"

	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	defaultState state = iota
	searchState
	MouseScrollSpeed = 3
	ListProportion   = 0.3
)

// Model represents the state of the TUI.
type Model struct {
	width, height int

	viewport viewport.Model

	// Add fields that represent the state of your TUI.
	count int
	list  list.Model

	statusMessage string
	ready         bool
	now           string

	keyMap
	state
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
}

func (i item) Name() string { return i.name }

func (i item) FilterValue() string { return i.name }

// InitialModel returns the initial state of the model.
func InitialModel(data ImageReport) *Model {
	var items []list.Item
	for _, imageInfo := range data.Images {
		image := item{
			name: imageInfo.Name,
		}
		items = append(items, image)
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Image Viewer"
	return &Model{
		list:   l,
		count:  0, // Initialize any state variables here.
		keyMap: defaultKeyMap(),
		state:  defaultState,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.tickCmd(), m.scanCmd(), m.countCmd())
}

// Update is called to handle user input and update the model's state.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case errMsg:
		m.statusMessage = msg.err.Error()
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		log.Printf("Width: %d Height: %d", m.width, m.height)
		statusBarHeight := lipgloss.Height(m.statusView())
		height := m.height - statusBarHeight
		listViewWidth := cast.ToInt(ListProportion * float64(m.width))
		listWidth := listViewWidth - listViewStyle.GetHorizontalFrameSize()
		log.Printf("list width height: %d %d", listWidth, height)
		m.list.SetSize(listWidth, height)

		detailViewWidth := m.width - listWidth
		log.Printf("viewport: %d %d", detailViewWidth, height)
		m.viewport = viewport.New(detailViewWidth, height)
		m.viewport.MouseWheelEnabled = true
		m.viewport.SetContent(m.viewportContent(m.viewport.Width))
	}
	switch m.state {
	case defaultState:
		cmds = append(cmds, m.handleDefaultState(msg))
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

// View returns the view that should be displayed.
// func (m Model) View() string {
// 	log.Println(m)
// 	return "Press q to quit.\n" + "Count: " + fmt.Sprint(m.count)
// }

func (m *Model) handleDefaultState(msg tea.Msg) tea.Cmd {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEscape:
			cmd = tea.Quit
			cmds = append(cmds, cmd)
		case tea.KeyCtrlC:
			cmd = tea.Quit
			cmds = append(cmds, cmd)
		case tea.KeyUp, tea.KeyDown, tea.KeyLeft, tea.KeyRight:
			m.list, cmd = m.list.Update(msg)
			cmds = append(cmds, cmd)
			cmds = append(cmds, cmd)
			// m.viewport.GotoTop()
			// m.viewport.SetContent(m.viewportContent(m.viewport.Width))
		}
	default:
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)

		cmds = append(cmds, cmd)

		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return tea.Batch(cmds...)
}

package models

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var (
	listViewStyle = lipgloss.NewStyle().
			PaddingRight(1).
			MarginRight(1).
			Border(lipgloss.RoundedBorder(), false, true, false, false)
	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"})

	statusNugget   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5")).Padding(0, 1)
	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})
	statusStyle = statusBarStyle.
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)
	encodingStyle = statusNugget.Background(lipgloss.Color("#A550DF")).Align(lipgloss.Right)
	statusText    = statusBarStyle
	datetimeStyle = statusNugget.Background(lipgloss.Color("#6124DF"))
)

func (m Model) listView() string {
	return listViewStyle.Render(m.list.View())
}

func (m Model) viewportContent(width int) string {
	var builder strings.Builder
	divider := dividerStyle.Render(strings.Repeat("-", width)) + "\n"
	if it := m.list.SelectedItem(); it != nil {
		key := fmt.Sprintf("Key: \n%s\n", it.(item).name)
		value := fmt.Sprintf("Value: \n%s\n", it.(item).name)
		builder.WriteString(divider)
		builder.WriteString(key)
		builder.WriteString(divider)
		builder.WriteString(value)
	}
	return wordwrap.String(builder.String(), width)
}

func (m Model) statusView() string {
	var status string
	var statusDesc string
	switch m.state {
	case searchState:
		status = "Search"
	default:
		status = "Ready"
		statusDesc = m.statusMessage
		if !m.ready {
			statusDesc = "Loading..."
		}
	}

	statusKey := statusStyle.Render(status)
	encoding := encodingStyle.Render("UTF-8")
	datetime := datetimeStyle.Render(m.now)

	statusVal := statusText.
		Width(m.width - lipgloss.Width(statusKey) - lipgloss.Width(encoding) - lipgloss.Width(datetime)).
		Render(statusDesc)

	bar := lipgloss.JoinHorizontal(lipgloss.Top, statusKey, statusVal, encoding, datetime)

	return statusBarStyle.Width(m.width).Render(bar)
}

func (m Model) View() string {
	// TODO: refresh status view only
	return lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.JoinHorizontal(lipgloss.Top, m.listView()),
		m.statusView(),
	)
}

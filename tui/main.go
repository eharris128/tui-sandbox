package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/eharris128/tui-sandbox/tui/models"
)

// RunTUI starts the TUI program.
// RunTUI starts the TUI program.
func RunTUI() {
	f, err := tea.LogToFile("debug.log", "debug")
	fmt.Println("Running TUI!!")
	if err != nil {
		fmt.Printf("RunTUI Logging - %v", err)
		os.Exit(1)
	}
	defer f.Close()
	p := tea.NewProgram(models.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("RunTUI error - %v", err)
		os.Exit(1)
	}
}

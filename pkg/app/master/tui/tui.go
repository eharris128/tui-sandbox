package tui

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/eharris128/tui-sandbox/pkg/app/master/tui/models"
)

// NewImageReport creates a new ImageReport from the given raw JSON data
func NewImageReport(rawData []byte) models.ImageReport {
	var report models.ImageReport
	err := json.Unmarshal(rawData, &report)
	if err != nil {
		fmt.Errorf("problem with unmarshal ImageReport: %v", err)
	}
	return report
}

func readData(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

// RunTUI starts the TUI program.
func RunTUI() {
	f, err := tea.LogToFile("debug.log", "debug")
	log.Println("Running TUI.")
	if err != nil {
		log.Printf("RunTUI Logging - %v", err)
		os.Exit(1)
	}

	defer f.Close()

	rawData := readData("pkg/app/master/tui/data/images.report.json")
	data := NewImageReport(rawData)

	p := tea.NewProgram(models.InitialModel(data))
	if _, err := p.Run(); err != nil {
		log.Printf("RunTUI error - %v", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"git.mattmohan.com/matt/advent2025/days"
)

func main() {
	dayArray := getDays()
	uiState := uiModel{days: dayArray,
		progressCh:   make(chan days.Progress, 500),
		progressBars: make([][2]progress.Model, len(dayArray)),
		table: table.New(
			table.WithColumns([]table.Column{
				{Title: "Day", Width: 6},
				{Title: "Name", Width: 20},
				{Title: "Part A", Width: 10},
				{Title: "Part B", Width: 10},
				{Title: "Time A", Width: 25},
				{Title: "Time B", Width: 25},
			}),
			table.WithFocused(true),
		),
	}

	for i := range uiState.progressBars {
		for j := 0; j < 2; j++ {
			uiState.progressBars[i][j] = progress.New(progress.WithDefaultGradient(), progress.WithWidth(25))
			uiState.progressBars[i][j].SetPercent(0)
		}
	}

	p := tea.NewProgram(uiState)

	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

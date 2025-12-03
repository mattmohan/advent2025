package main

import "fmt"

func (m uiModel) View() string {
	output := "Advent of Code 2025\n\n"
	output += m.table.View()
	output += fmt.Sprintf("\n%s\n", m.help.View(m.keymap))
	for i := range m.days {
		for j := range 2 {
			output += fmt.Sprintf("Day %d Part %d: %s\n", i+1, j+1, m.progressBars[i][j].View())
		}
	}
	return output
}

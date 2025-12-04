package main

import (
	"fmt"
	"os"
	"time"

	"git.mattmohan.com/matt/advent2025/days"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/table"

	tea "github.com/charmbracelet/bubbletea"
)

type uiModel struct {
	days         []DayPlugin
	table        table.Model
	progressCh   chan days.Progress
	progressBars [][2]progress.Model
	help         help.Model
	keymap       keyMap
}

type RunDay struct {
	day *days.Day
}
type LoadDay struct {
	day *days.Day
}
type UpdateTable struct{}

var _ tea.Model = uiModel{}

func (m uiModel) Init() tea.Cmd {
	return tea.Batch(
		updateTable(),
		listenProgress(m.progressCh),
	)
}

func (m uiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		case key.Matches(msg, m.keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.Execute):
			currentDay := m.table.Cursor()
			day := &(m.days)[currentDay]
			return m, func() tea.Msg { return LoadDay{day: &day.Day} }
		}
	case LoadDay:
		filename := fmt.Sprintf("inputs/day%d.txt", msg.day.Number)
		file, err := os.ReadFile(filename)
		if err != nil {
			msg.day.Parts[0].Result = "Input file missing"
			msg.day.Parts[1].Result = "Input file missing"
			return m, nil
		}

		msg.day.Input = file
		msg.day.Parts[0].Result = ""
		msg.day.Parts[1].Result = ""

		return m, func() tea.Msg { return RunDay{day: msg.day} }
	case RunDay:
		go msg.day.Run(days.PartA, m.progressCh)
		go msg.day.Run(days.PartB, m.progressCh)
		return m, nil

	case progress.FrameMsg:
		cmds := make([]tea.Cmd, 0, len(m.progressBars)*2+1)
		for i := range m.progressBars {
			for j := 0; j < 2; j++ {
				tmp, cmd1 := m.progressBars[i][j].Update(msg)
				m.progressBars[i][j] = tmp.(progress.Model)
				cmds = append(cmds, cmd1)
			}
		}
		cmds = append(cmds, updateTable())
		return m, tea.Batch(cmds...)

	case days.Progress:
		// Handle progress updates
		dayIdx := msg.Day
		partIdx := int(msg.Part)
		cmd := m.progressBars[dayIdx][partIdx].SetPercent(float64(msg.Percent))
		return m, tea.Sequence(cmd, listenProgress(m.progressCh))
	case UpdateTable:
		rows := make([]table.Row, 0, len(m.days))
		for _, day := range m.days {
			rows = append(rows, table.Row{
				fmt.Sprintf("%d", day.Day.Number),
				day.Day.Name,
				day.Day.Parts[0].Result,
				day.Day.Parts[1].Result,
				day.Day.Parts[0].Duration.Truncate(100 * time.Nanosecond).String(),
				day.Day.Parts[1].Duration.Truncate(100 * time.Nanosecond).String(),
			})
		}
		m.table.SetRows(rows)
		return m, nil
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func listenProgress(ch chan days.Progress) tea.Cmd {
	return func() tea.Msg {
		prog := <-ch
		return prog
	}
}

func updateTable() tea.Cmd {
	return func() tea.Msg {
		return UpdateTable{}
	}
}

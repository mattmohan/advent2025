package days

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type DayPartFunction func([]byte, chan string, chan float64)
type DayPart struct {
	PartFunc DayPartFunction
	Result   string
	Duration time.Duration
}
type Day struct {
	Number int
	Name   string
	Parts  [2]DayPart
	Input  []byte
}

func (day *Day) Run(partIdx PartNumber, taskProg chan Progress) tea.Cmd {
	part := &day.Parts[partIdx]
	progress := make(chan float64)
	result := make(chan string)

	taskProg <- Progress{Percent: 0.0, Done: false, Day: day.Number - 1, Part: partIdx}

	start := time.Now()
	go part.PartFunc(day.Input, result, progress)
	done := false
	nextProgress := float32(0.0)
	var timer *time.Timer
	for !done {
		select {
		case res := <-result:
			part.Duration = time.Since(start)
			part.Result = res
			if timer != nil {
				timer.Stop()
			}
			taskProg <- Progress{Percent: float32(1.0), Done: true, Day: day.Number - 1, Part: partIdx}
			done = true
		case p := <-progress:
			nextProgress = float32(p)
			if timer == nil {
				timer = time.AfterFunc(10*time.Millisecond, func() {
					taskProg <- Progress{Percent: nextProgress, Done: false, Day: day.Number - 1, Part: partIdx}
					timer = nil
				})
			}
		}
	}
	return nil
}

type PartNumber int

const (
	PartA PartNumber = iota
	PartB
)

type Progress struct {
	Percent float32
	Done    bool
	Day     int
	Part    PartNumber
}

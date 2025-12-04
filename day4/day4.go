package main

import (
	"strconv"
	"strings"

	"git.mattmohan.com/matt/advent2025/days"
	"git.mattmohan.com/matt/advent2025/utils"
)

func GetDay() days.Day {
	return days.Day{
		Number: 4,
		Name:   "Printing Department",
		Parts: [2]days.DayPart{
			{PartFunc: day4Part1},
			{PartFunc: day4Part2},
		},
	}
}

func parseInput(inputChars []byte) utils.Grid[bool] {
	lines := strings.Split(string(inputChars), "\n")
	rows := len(lines)
	cols := len(lines[0])

	grid := utils.NewGrid(rows, cols, false)
	for y, line := range lines {
		for x, char := range line {
			if char == '@' {
				grid.Set(x, y, true)
			}
		}
	}
	return grid
}

func day4Part1(inputChars []byte, result chan string, progress chan float64) {
	defer close(result)
	defer close(progress)
	grid := parseInput(inputChars)
	count := 0
	grid.Walk(func(x, y int, value bool) {
		if !value {
			return
		}
		neighbors := 0
		grid.WalkNeighbors(x, y, func(nx, ny int, value bool) {
			if value {
				neighbors++
			}
		})

		if neighbors < 4 {
			count++
		}
	})
	result <- strconv.Itoa(count)
}

func day4Part2(inputChars []byte, result chan string, progress chan float64) {
	defer close(result)
	defer close(progress)

	// Create two grids to let us alternate between them
	grids := [2]utils.Grid[bool]{parseInput(inputChars), parseInput(inputChars)}

	removedThisRound := 1
	removed := 0
	loop := 0
	for removedThisRound > 0 && loop < 1000 {
		removedThisRound = 0

		grid := &grids[loop%2]
		nextGrid := &grids[(loop+1)%2]

		grid.Walk(func(x, y int, value bool) {
			// Bail if the cell is already empty
			if !value {
				return
			}
			neighbors := 0
			grid.WalkNeighbors(x, y, func(nx, ny int, value bool) {
				if value {
					neighbors++
				}
			})

			if neighbors < 4 {
				removedThisRound++
				nextGrid.Set(x, y, false)
			}
		})

		loop++
		removed += removedThisRound
		// Copy nextGrid back to grid for the next iteration
		grid.Clone(*nextGrid)
	}

	result <- strconv.Itoa(removed)
}

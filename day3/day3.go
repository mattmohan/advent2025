package day3

import (
	"git.mattmohan.com/matt/advent2025/days"
)

func GetDay() days.Day {
	return days.Day{
		Number: 3,
		Name:   "Sample Day",
		Parts: [2]days.DayPart{
			{PartFunc: func(input []byte, result chan string, progress chan float64) {
				defer close(result)
				// Placeholder logic for Part A
				result <- "Not implemented yet"
			}},
			{PartFunc: func(input []byte, result chan string, progress chan float64) {
				defer close(result)
				// Placeholder logic for Part B
				result <- "Not implemented yet"
			}},
		},
	}
}

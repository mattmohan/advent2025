package day1

import (
	"strconv"
	"strings"

	"git.mattmohan.com/matt/advent2025/days"
)

func GetDay() days.Day {
	return days.Day{
		Number: 1,
		Name:   "Secret Entrance",
		Parts: [2]days.DayPart{
			{PartFunc: day1Part1},
			{PartFunc: day1Part2},
		},
	}
}

type Day1Input struct {
	steps []int
}

func day1Part1(inputChars []byte, result chan string, progress chan float64) {
	// Solve Part A
	defer close(result)
	defer close(progress)
	input := parseInputDay1(inputChars)

	cur := 50
	count := 0

	for i, step := range input.steps {
		cur = (cur + step) % 100
		if cur == 0 {
			count++
		}

		progress <- float64(i) / float64(len(input.steps))

	}
	result <- strconv.Itoa(count)
}

func day1Part2(inputChars []byte, result chan string, progress chan float64) {
	defer close(progress)
	defer close(result)
	input := parseInputDay1(inputChars)
	cur := 50
	count := int(0)
	for i, step := range input.steps {
		// Find the next position
		next := (cur + step) % 100
		if next < 0 {
			next += 100
		}

		count += countZeroCrossings(cur, step)
		cur = next
		progress <- float64(i) / float64(len(input.steps))
	}
	result <- strconv.Itoa(count)

}

func parseInputDay1(input []byte) Day1Input {
	if len(input) == 0 {
		panic("empty input")
	}
	lines := strings.Split(string(input), "\n")
	steps := make([]int, 0, len(lines))

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if trimmed[0] != 'L' && trimmed[0] != 'R' {
			panic("invalid direction")
		}
		tmp, err := strconv.Atoi(trimmed[1:])
		if err != nil {
			panic("invalid step count")
		}
		if trimmed[0] == 'L' {
			tmp *= -1
		}
		steps = append(steps, tmp)
	}
	// Placeholder parsing logic
	return Day1Input{steps: steps}
}

func countZeroCrossings(start, step int) (count int) {
	next := (start + step) % 100
	if next < 0 {
		next += 100
	}

	// Handle the simple case of ending on zero
	if next == 0 {
		count++
	} else if start != 0 {
		// Handle crossing zero
		crossingLeft := start < next
		steppingLeft := step < 0

		if steppingLeft == crossingLeft {
			count++
		}
	}

	// Handle extra passes
	if step > 0 {
		count += step / 100
	} else {
		count += step / -100
	}
	return count
}

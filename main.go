package main

import (
	"flag"
	"os"
	"strconv"
	"strings"
)

type Day1Input struct {
	steps []int
}

func main() {
	test := flag.Bool("test", false, "Run test cases instead of actual input")
	flag.Parse()

	fileName := "input.txt"

	if *test {
		fileName = "test_input.txt"
	}

	contents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	input := parseInputDay1(contents)
	outputA := solveDay1PartA(input)
	outputB := solveDay1PartB(input)
	println("Day 1 Part A:", outputA)
	println("Day 1 Part B:", outputB)
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

func solveDay1PartA(input Day1Input) int {
	cur := 50
	count := 0
	for _, step := range input.steps {
		cur = (cur + step) % 100
		if cur == 0 {
			count++
		}
	}
	return count
}

func solveDay1PartB(input Day1Input) int {
	cur := 50
	count := int(0)
	for _, step := range input.steps {
		// Find the next position
		next := (cur + step) % 100
		if next < 0 {
			next += 100
		}

		count += countZeroCrossings(cur, step)
		cur = next
	}
	return count
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

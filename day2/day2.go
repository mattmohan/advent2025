package main

import (
	"strconv"
	"strings"

	"git.mattmohan.com/matt/advent2025/days"
)

func GetDay() days.Day {
	return days.Day{
		Number: 2,
		Name:   "Gift Shop",
		Parts: [2]days.DayPart{
			{PartFunc: day2Part1},
			{PartFunc: day2Part2},
		},
	}
}

type Range struct {
	start int
	end   int
}

func parseInput(inputChars []byte) []Range {
	rangeStrings := strings.Split(string(inputChars), ",")
	ranges := make([]Range, 0, len(rangeStrings))

	for _, rangeString := range rangeStrings {
		trimmedRange := strings.TrimSpace(rangeString)
		parts := strings.Split(trimmedRange, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, Range{start: start, end: end})
	}
	return ranges
}

func day2Part1(inputChars []byte, result chan string, progress chan float64) {
	defer close(result)
	defer close(progress)

	count := 0
	ranges := parseInput(inputChars)

	for _, currentRange := range ranges {
	rangeloop:
		for i := currentRange.start; i <= currentRange.end; i++ {
			str := strconv.Itoa(i)
			if len(str)%2 != 0 {
				continue
			}
			midpoint := len(str) / 2
			firstHalf := str[:midpoint]
			secondHalf := str[midpoint:]
			for j := 0; j < midpoint; j++ {
				if firstHalf[j] != secondHalf[j] {
					continue rangeloop
				}
			}
			count += i
		}
	}
	// Placeholder logic for Part A
	result <- strconv.Itoa(count)
}

func day2Part2(inputChars []byte, result chan string, progress chan float64) {
	defer close(result)
	defer close(progress)

	count := 0
	ranges := parseInput(inputChars)

	// Iterate over all of the ranges
	for rangeIdx, currentRange := range ranges {
		// Iterate over all numbers in the range
		for i := currentRange.start; i <= currentRange.end; i++ {
			// For each number, try skipping digits to find a repeat pattern starting with a skip distance of 1 up to half the length of the number
			str := strconv.Itoa(i)
			longestSkip := len(str) / 2

		skiploop:
			for skip := 1; skip <= longestSkip; skip++ {
				// Bail if not divisible
				if len(str)%skip != 0 {
					continue skiploop
				}

				// Start at each position and check that everything within the window matches
				for startCheck := 0; startCheck < len(str)-skip; startCheck++ {
					if str[startCheck] != str[(startCheck+skip)] {
						continue skiploop
					}
				}
				// If we got here, we found a valid skip so accumulate and move on
				count += i
				break skiploop
			}
		}
		// Report back on progress
		progress <- float64(rangeIdx) / float64(len(ranges))
	}
	// Placeholder logic for Part A
	result <- strconv.Itoa(count)
}

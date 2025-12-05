package main

import (
	"sort"
	"strconv"
	"strings"

	"git.mattmohan.com/matt/advent2025/days"
	"git.mattmohan.com/matt/advent2025/utils"
)

func GetDay() days.Day {
	return days.Day{
		Number: 5,
		Name:   "Cafeteria",
		Parts: [2]days.DayPart{
			{PartFunc: part1},
			{PartFunc: part2},
		},
	}
}

func parseInput(inputChars []byte) ([]utils.Range, []int64) {
	lines := strings.Split(string(inputChars), "\n")

	ranges := make([]utils.Range, 0, len(lines)) //This overallocates but I don't want to scan twice, so 🤷
	var available []int64

	doneRanges := false
	for i, line := range lines {
		if line == "" {
			doneRanges = true
			available = make([]int64, 0, len(lines)-i-1)
			continue
		}
		if doneRanges {
			// Parse available items
			num := utils.ParseInt64(line)
			available = append(available, num)
		} else {
			// Parse range by splitting line on "-" and append to ranges
			parts := strings.Split(line, "-")
			start := utils.ParseInt64(parts[0])
			end := utils.ParseInt64(parts[1])
			ranges = append(ranges, utils.Range{Start: start, End: end})
		}
	}
	return ranges, available
}

func part1(inputChars []byte, result chan string, progress chan float64) {
	defer close(result)
	defer close(progress)
	ranges, available := parseInput(inputChars)

	count := 0
	for i, num := range available {
		for _, r := range ranges {
			if r.InRange(num) {
				count++
				break
			}
		}
		progress <- float64(i) / float64(len(available))
	}
	result <- strconv.Itoa(count)
}

func part2(inputChars []byte, result chan string, progress chan float64) {
	defer close(result)
	defer close(progress)

	ranges, _ := parseInput(inputChars)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})
	deleted := make([]bool, len(ranges))
	var count int64 = 0

	// Loop through ranges and merge overlapping ones
	for i := range ranges {
		// Skip deleted ranges
		if deleted[i] {
			continue
		}
		// Check for overlaps with other ranges
		for j := range ranges {
			// Skip self and deleted ranges
			if i == j || deleted[j] {
				continue
			}

			// If they overlap, merge and mark the other range as deleted
			if ranges[i].Overlaps(ranges[j]) {
				ranges[i] = ranges[i].Merge(ranges[j])
				deleted[j] = true
			}
		}
	}

	// Count total numbers covered by non-deleted ranges
	for i, r := range ranges {
		if deleted[i] {
			continue
		}
		count += r.End - r.Start + 1
		progress <- float64(i) / float64(len(ranges))
	}

	result <- strconv.FormatInt(count, 10)
}

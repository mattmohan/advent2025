package day3

import (
	"strconv"
	"strings"

	"git.mattmohan.com/matt/advent2025/days"
)

func GetDay() days.Day {
	return days.Day{
		Number: 3,
		Name:   "Lobby",
		Parts: [2]days.DayPart{
			{PartFunc: day3Part1},
			{PartFunc: day3Part2},
		},
	}
}

func day3Part1(inputChars []byte, result chan string, progress chan float64) {
	defer close(result)
	defer close(progress)

	count := 0
	banks := strings.Split(string(inputChars), "\n")
	for i, bank := range banks {
		count += findHighest(convertStringBankToIntSlice(bank), 2)
		progress <- float64(i) / float64(len(banks))
	}

	result <- strconv.Itoa(count)
}

func day3Part2(inputChars []byte, result chan string, progress chan float64) {
	defer close(result)
	defer close(progress)

	count := 0
	banks := strings.Split(string(inputChars), "\n")
	for i, bank := range banks {
		count += findHighest(convertStringBankToIntSlice(bank), 12)
		progress <- float64(i) / float64(len(banks))
	}

	result <- strconv.Itoa(count)
}

func convertStringBankToIntSlice(bankString string) []int {
	bank := make([]int, len(bankString))
	for i := range bankString {
		val, err := strconv.Atoi(string(bankString[i]))
		if err != nil {
			panic(err)
		}
		bank[i] = val
	}
	return bank
}

func findHighest(bank []int, length int) int {
	highIdxs := make([]int, length)

	// Start by assuming the first digits are the highest
	for i := range highIdxs {
		highIdxs[i] = i
	}

	start := 0
	// For each place in the output find the appropriate highest digit
	for place := 0; place < length; place++ {
		// After the first place start searching after the last highest digit found
		if place > 0 {
			start = highIdxs[place-1] + 1
		}

		// Stop searching so that there are enough digits left to fill the remaining places
		end := len(bank) - (length - place - 1)

		for i := start; i < end; i++ {
			if bank[i] > bank[highIdxs[place]] {
				// If found update the highest index for this place
				highIdxs[place] = i
				// And update all subsequent places to be one after this
				for j := place + 1; j < length; j++ {
					highIdxs[j] = highIdxs[j-1] + 1
				}
			}
		}
	}

	// Build the final number from the highest digits found
	sum := 0
	for _, idx := range highIdxs {
		sum = sum*10 + bank[idx]
	}

	return sum
}

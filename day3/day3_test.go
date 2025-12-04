package main

import (
	"testing"
)

var textInput = []byte(`987654321111111
811111111111119
234234234234278
818181911112111`)

func TestDay3Part1(t *testing.T) {
	// Placeholder test for Part 1
	expected := "357"
	resultChan := make(chan string)
	progressChan := make(chan float64)
	go day3Part1(textInput, resultChan, progressChan)
	result := <-resultChan
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestDay3Part2(t *testing.T) {
	// Placeholder test for Part 2
	table := []struct {
		input    []byte
		expected string
	}{
		{[]byte(
			"987654321111111"), "987654321111"},
		{[]byte(
			"811111111111119"), "811111111119"},
		{[]byte("234234234234278"), "434234234278"},
		{[]byte("818181911112111"), "888911112111"},
		{[]byte(`987654321111111
811111111111119
234234234234278
818181911112111`), "3121910778619"},
	}
	for _, test := range table {
		resultChan := make(chan string)
		progressChan := make(chan float64)
		go day3Part2(test.input, resultChan, progressChan)
		result := <-resultChan
		if result != test.expected {
			t.Errorf("For input: \n===\n%s\n===\n  expected: %s,\n  got:      %s", test.input, test.expected, result)
		}
	}
}

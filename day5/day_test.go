package main

import (
	"testing"
)

var textInput = []byte(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)

func TestPart1(t *testing.T) {
	// Placeholder test for Part 1
	expected := "3"
	resultChan := make(chan string)
	progressChan := make(chan float64, 500)
	go part1(textInput, resultChan, progressChan)
	result := <-resultChan
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	// Placeholder test for Part 2
	expected := "14"
	resultChan := make(chan string)
	progressChan := make(chan float64, 500)
	go part2(textInput, resultChan, progressChan)
	result := <-resultChan
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart2Extended(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name: "Extended Test Case 1",
			input: []byte(`1-10
2-5`),
			expected: "10",
		},
		{
			name: "Extended Test Case 2",
			input: []byte(`3-5
10-14
16-20
12-18
13-14
13-13`),
			expected: "14",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultChan := make(chan string)
			progressChan := make(chan float64, 500)
			go part2(tc.input, resultChan, progressChan)
			result := <-resultChan
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

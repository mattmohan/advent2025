package main

import (
	"testing"
)

var textInput = []byte(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`)

func TestDay4Part1(t *testing.T) {
	// Placeholder test for Part 1
	expected := "13"
	resultChan := make(chan string)
	progressChan := make(chan float64)
	go day4Part1(textInput, resultChan, progressChan)
	result := <-resultChan
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestDay4Part2(t *testing.T) {
	// Placeholder test for Part 2
	expected := "43"
	resultChan := make(chan string)
	progressChan := make(chan float64)
	go day4Part2(textInput, resultChan, progressChan)
	result := <-resultChan
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

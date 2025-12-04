package main

import (
	"strings"
	"testing"
)

func TestSolveDay1PartA(t *testing.T) {
	testInput := []byte(`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`)
	expectedOutput := "3"

	progress := make(chan float64, 100)
	result := make(chan string)
	go day1Part1(testInput, result, progress)
	output := <-result

	if strings.Compare(output, expectedOutput) != 0 {
		t.Errorf("Expected %s, got %s for Solve", expectedOutput, output)
	}
}

func TestSolveDay1PartB(t *testing.T) {
	testInput := []byte(`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`)
	expectedOutput := "6"

	progress := make(chan float64, 100)
	result := make(chan string)
	go day1Part2(testInput, result, progress)
	output := <-result

	if strings.Compare(output, expectedOutput) != 0 {
		t.Errorf("Expected %s, got %s for Solve", expectedOutput, output)
	}
}

package main

import (
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
	expectedOutput := 3

	input := parseInputDay1(testInput)
	output := solveDay1PartA(input)

	if output != expectedOutput {
		t.Errorf("Expected %d, got %d for Solve", expectedOutput, output)
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
	expectedOutput := 6

	input := parseInputDay1(testInput)
	output := solveDay1PartB(input)

	if output != expectedOutput {
		t.Errorf("Expected %d, got %d for Solve", expectedOutput, output)
	}
}

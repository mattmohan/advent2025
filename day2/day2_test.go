package main

import (
	"testing"
)

func TestSolveDay2PartA(t *testing.T) {
	testInput := []byte(`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`)
	expectedOutput := "1227775554"
	progress := make(chan float64, 100)
	result := make(chan string)
	go func() {
		day2Part1(testInput, result, progress)
	}()
	output := <-result

	if output != expectedOutput {
		t.Errorf("Expected %s, got %s for Solve", expectedOutput, output)
	}
	for range progress {

	}
}

func TestSolveDay2PartB(t *testing.T) {
	testInput := []byte(`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`)
	expectedOutput := "4174379265"
	progress := make(chan float64, 100)
	result := make(chan string)
	go func() {
		day2Part2(testInput, result, progress)
	}()
	output := <-result

	if output != expectedOutput {
		t.Errorf("Expected %s, got %s for Solve", expectedOutput, output)
	}
	for range progress {

	}
}

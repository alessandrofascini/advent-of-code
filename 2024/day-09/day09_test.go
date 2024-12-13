package day09

import (
	"os"
	"testing"
)

func createInput(filename string) []byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}

func testPart(t *testing.T, fn func(input []byte) int, expected int) {
	inputTest := createInput("test.txt")
	answer := fn(inputTest)
	if answer != expected {
		t.Errorf("Expected %d, got %d", expected, answer)
		return
	}
	inputPuzzle := createInput("puzzle.txt")
	answer = fn(inputPuzzle)
	t.Logf("Answer: %d", answer)
}

func TestPart1(t *testing.T) {
	testPart(t, Part1, 1928)
}

func TestPart2(t *testing.T) {
	// 6304576012713
	testPart(t, Part2, 2858)
}

func TestPart3(t *testing.T) {
	inputPuzzle := createInput("test.txt")
	Part2BruteForce(inputPuzzle)
}

package day12

import (
	"bytes"
	"os"
	"testing"
)

func createInput(filename string) [][]byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bytes.Split(b, []byte("\n"))
}

func testPart(t *testing.T, fn func(input [][]byte) int, expected int) {
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
	testPart(t, Part1, 1930)
}

func TestPart2(t *testing.T) {
	testPart(t, Part2, 1206)
}

package day23

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

func testPart(t *testing.T, fn func([][]byte) int, expected int, proceed bool) {
	inputTest := createInput("test.txt")
	answer := fn(inputTest)
	if answer != expected {
		t.Errorf("Expected %d, got %d", expected, answer)
		return
	}
	t.Logf("Test Answer: %d", answer)
	if !proceed {
		return
	}
	inputPuzzle := createInput("puzzle.txt")
	answer = fn(inputPuzzle)
	t.Logf("Puzzle Answer: %d", answer)
}

func TestPart1(t *testing.T) {
	testPart(t, Part1, 94, true)
}

func TestPart2(t *testing.T) {
	testPart(t, Part2, 154, false)
}

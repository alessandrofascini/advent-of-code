package day10

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func createInput(filename string) [][]byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	var input [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, []byte(line))
	}
	return input
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
	testPart(t, Part1, 36)
}

func TestPart2(t *testing.T) {
	// 6304576012713
	testPart(t, Part2, 81)
}

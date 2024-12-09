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

	input := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, []byte(line))
	}
	return input
}

func testPart(t *testing.T, fn func(eqs [][]byte) int, expected int) {
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
	testPart(t, Part1, 14)
}

func TestPart2(t *testing.T) {
	testPart(t, Part2, 34)
}

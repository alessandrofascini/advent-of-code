package day06

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func createInput(t *testing.T, filename string) [][]byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))

	input := make([][]byte, 0)
	for scanner.Scan() {
		bts := []byte(scanner.Text())
		input = append(input, bts)
	}
	return input
}

func testPart(t *testing.T, fn func(g [][]byte) int, expected int) {
	inputTest := createInput(t, "test.txt")
	answer := fn(inputTest)
	if answer != expected {
		t.Errorf("Expected %d, got %d", expected, answer)
		return
	}
	inputPuzzle := createInput(t, "puzzle.txt")
	answer = fn(inputPuzzle)
	t.Logf("Answer: %d", answer)
}

func TestPart1(t *testing.T) {
	testPart(t, Part1, 41)
}

func TestPart2(t *testing.T) {
	testPart(t, Part2, 6)
}

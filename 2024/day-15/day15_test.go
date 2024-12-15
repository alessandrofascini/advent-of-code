package day15

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func createInput(filename string) ([][]byte, []byte) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))

	grid := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		grid = append(grid, []byte(line))
	}

	movements := make([]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		movements = append(movements, []byte(line)...)
	}
	return grid, movements
}

func testPart(t *testing.T, fn func([][]byte, []byte) int, expected int) {
	g, m := createInput("test.txt")
	answer := fn(g, m)
	if answer != expected {
		t.Errorf("Expected %d, got %d", expected, answer)
		return
	}
	g, m = createInput("puzzle.txt")
	answer = fn(g, m)
	t.Logf("Puzzle Answer: %d", answer)
}

func TestSmallerInput(t *testing.T) {
	const expected = 2028
	g, m := createInput("smaller_test.txt")
	answer := Part1(g, m)
	if answer != expected {
		t.Errorf("Expected %d, got %d", expected, answer)
		return
	}
	t.Logf("Smaller Test Answer: %d", answer)
}

func TestPart1(t *testing.T) {
	testPart(t, Part1, 10092)
}

func TestSmallerInput2(t *testing.T) {
	const expected = 618
	g, m := createInput("smaller_test2.txt")
	answer := Part2(g, m)
	if answer != expected {
		t.Errorf("Expected %d, got %d", expected, answer)
		return
	}
	t.Logf("Smaller Test Answer: %d", answer)
}

func TestPart2(t *testing.T) {
	testPart(t, Part2, 9021)
}

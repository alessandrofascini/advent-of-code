package day16

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
	scanner := bufio.NewScanner(bytes.NewReader(b))

	grid := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		grid = append(grid, []byte(line))
	}
	return grid
}

func testPart(t *testing.T, fn func([][]byte) int, expected ...int) {
	if len(expected) < 1 {
		t.Fatalf("no expected values; required at least one value")
	}
	g := createInput("test.txt")
	answer := fn(g)
	if answer != expected[0] {
		t.Errorf("Test Expected %d, got %d", expected[0], answer)
		return
	}
	g = createInput("puzzle.txt")
	answer = fn(g)
	if len(expected) > 1 && answer != expected[1] {
		t.Errorf("Puzzle Expected %d, got %d", expected[1], answer)
		return
	}
}

func TestPart1(t *testing.T) {
	testPart(t, Part1, 7036, 107468)
}

func TestPart2(t *testing.T) {
	testPart(t, Part2, 45)
}

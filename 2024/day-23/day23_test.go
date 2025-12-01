package day23

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func createInput(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func TestPart1(t *testing.T) {
	numbers := createInput("test.txt")
	answer := Part1(numbers)
	t.Logf("Test Answer: %d", answer)
	numbers = createInput("puzzle.txt")
	answer = Part1(numbers)
	t.Logf("Puzzle Answer: %d", answer)
}

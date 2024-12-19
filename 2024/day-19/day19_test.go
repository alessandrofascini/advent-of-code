package day19

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"
)

func createInput(filename string) ([]string, []string) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))
	output := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, line)
	}
	return strings.Split(output[0], ", "), output[2:]
}

func TestPart1(t *testing.T) {
	const expected = 6
	p, d := createInput("test.txt")
	answer := Part1(p, d)
	if answer != expected {
		t.Fatalf("Expected %d, got %d", expected, answer)
	}
	p, d = createInput("puzzle.txt")
	answer = Part1(p, d)
	t.Logf("Puzzle Answer: %d", answer)
}

func TestPart2(t *testing.T) {
	const expected = 16
	p, d := createInput("test.txt")
	answer := Part2(p, d)
	if answer != expected {
		t.Fatalf("Expected %d, got %d", expected, answer)
	}
	p, d = createInput("puzzle.txt")
	answer = Part2(p, d)
	t.Logf("Puzzle Answer: %d", answer)
}

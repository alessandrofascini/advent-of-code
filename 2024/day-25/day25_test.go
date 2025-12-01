package day25

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func createInput(filename string) []Schematic {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	schematics := make([]Schematic, 0)
	curr := make([][]byte, 0)

	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			schematics = append(schematics, curr)
			curr = make([][]byte, 0, len(curr))
		} else {
			curr = append(curr, []byte(line))
		}
	}
	return schematics
}

func TestPart1(t *testing.T) {
	numbers := createInput("test.txt")
	answer := Part1(numbers)
	t.Logf("Test Answer: %d", answer)
	numbers = createInput("puzzle.txt")
	answer = Part1(numbers)
	t.Logf("Puzzle Answer: %d", answer)
}

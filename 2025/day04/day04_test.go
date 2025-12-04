package day04

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func ReadInput(t *testing.T, filename string) []Row {
	b, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	input := make([]Row, 0)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	return input
}

func TestPart1(t *testing.T) {
	input := ReadInput(t, "test.txt")
	t.Log("test:", Part1(input))

	input = ReadInput(t, "puzzle.txt")
	t.Log("puzzle:", Part1(input))
}
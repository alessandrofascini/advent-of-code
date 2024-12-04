package day_04

import (
	"bytes"
	"os"
	"testing"
)

func createInput(t *testing.T, filename string) [][]byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return bytes.Split(b, []byte{'\n'})
}

func TestPart1(t *testing.T) {
	input := createInput(t, "test.txt")
	answer := Part1(input)
	t.Log(answer)
}

func TestPart1Puzzle(t *testing.T) {
	input := createInput(t, "puzzle.txt")
	answer := Part1(input)
	t.Log(answer)
}

func TestPart2(t *testing.T) {
	input := createInput(t, "test.txt")
	answer := Part2(input)
	t.Log(answer)
}

func TestPart2Puzzle(t *testing.T) {
	input := createInput(t, "puzzle.txt")
	answer := Part2(input)
	t.Log(answer)
}

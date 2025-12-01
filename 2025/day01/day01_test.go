package day01

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"testing"
)

func createInput(t *testing.T, filename string) []Row {
	b, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	input := make([]Row, 0)
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		value, _ := strconv.Atoi(line[1:])
		input = append(input, Row{dir, value})
	}
	return input
}

func TestExamplePart1(t *testing.T) {
	input := createInput(t, "test.txt")
	answer := Part1(input)
	t.Log(answer)
}

func TestPart1(t *testing.T) {
	input := createInput(t, "puzzle.txt")
	t.Log(input)
	answer := Part1(input)
	t.Log(answer)
}

func TestExamplePart2(t *testing.T) {
	input := createInput(t, "test.txt")
	answer := Part2(input)
	t.Log(answer)
}

func TestPart2(t *testing.T) {
	input := createInput(t, "puzzle.txt")
	answer := Part2(input)
	t.Log(answer)
}

package day02

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func readInput(t *testing.T, filename string) []Range {
	b, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	input := make([]Range, 0)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		for _, r := range ranges {
			splitted := strings.Split(r, "-")
			lower, _ := strconv.ParseInt(splitted[0], 10, 64)
			upper, _ := strconv.ParseInt(splitted[1], 10, 64)
			input = append(input, Range{lower, upper})
		}
	}
	return input
}

func TestPart1(t *testing.T) {
	input := readInput(t, "test.txt")
	t.Log("Test: ", Part1(input))
	input = readInput(t, "puzzle.txt")
	t.Log("Puzzle: ", Part1(input))
}

func TestPart2(t *testing.T) {
	input := readInput(t, "test.txt")
	t.Log("Test: ", Part2(input))
	input = readInput(t, "puzzle.txt")
	t.Log("Puzzle: ", Part2(input))
}

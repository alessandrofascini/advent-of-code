package day02

import "testing"

func readInput(fileName string) []Range {
	return make([]Range, 0)
}

func TestPart1(t *testing.T) {
	input := readInput("test.txt")
	t.Log("Test: ", Part1(input))
	input = readInput("puzzle.txt")
	t.Log("Puzzle: ", Part1(input))
}
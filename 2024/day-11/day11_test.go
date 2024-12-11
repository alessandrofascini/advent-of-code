package day11

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func createInput(filename string) []int {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(b), " ")
	nums := make([]int, len(s))
	for i, v := range s {
		nums[i], _ = strconv.Atoi(v)
	}
	return nums
}

func testPart(t *testing.T, fn func(input []int) int, expected int) {
	inputTest := createInput("test.txt")
	answer := fn(inputTest)
	if answer != expected {
		t.Errorf("Expected %d, got %d", expected, answer)
		return
	}
	inputPuzzle := createInput("puzzle.txt")
	answer = fn(inputPuzzle)
	t.Logf("Answer: %d", answer)
}

func TestPart1(t *testing.T) {
	testPart(t, Part1, 55312)
}

func TestPart2(t *testing.T) {
	inputPuzzle := createInput("puzzle.txt")
	answer := Part2(inputPuzzle)
	t.Logf("Answer: %d", answer)
}

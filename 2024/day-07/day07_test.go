package day07

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func createInput(filename string) [][]int {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))

	input := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		split[0] = split[0][:len(split[0])-1]
		numbers := make([]int, len(split))
		for i, s := range split {
			numbers[i], _ = strconv.Atoi(s)
		}
		input = append(input, numbers)
	}
	return input
}

func testPart(t *testing.T, fn func(eqs [][]int) int, expected int) {
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
	testPart(t, Part1, 3749)
}

func TestPart2(t *testing.T) {
	testPart(t, Part2, 11387)
}

func BenchmarkMain(b *testing.B) {
	inputPuzzle := createInput("puzzle.txt")
	answer := Part2(inputPuzzle)
	b.Logf("Answer: %d", answer)
}

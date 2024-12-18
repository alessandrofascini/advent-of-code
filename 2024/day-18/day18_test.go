package day18

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
	input := make([][]int, 0)

	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		i, _ := strconv.Atoi(split[0])
		j, _ := strconv.Atoi(split[1])
		input = append(input, []int{j, i})
	}
	return input
}

func TestPart1(t *testing.T) {
	const expected = 22
	points := createInput("test.txt")
	answer := Part1(6, points[:12])
	if answer != expected {
		t.Errorf("Test Expected %d, got %d", expected, answer)
		return
	}
	points = createInput("puzzle.txt")
	answer = Part1(70, points[:1024])
	t.Logf("Puzzle Answer %d", answer)
}

func TestPart2(t *testing.T) {
	const expected = "6,1"
	points := createInput("test.txt")
	answer := Part2(6, points)
	if answer != expected {
		t.Errorf("Test Expected %s, got %s", expected, answer)
		return
	}
	points = createInput("puzzle.txt")
	answer = Part2(70, points)
	t.Logf("Puzzle Answer %s", answer)
}

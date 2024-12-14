package day14

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func createInput(filename string) []*Robot {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))
	robots := make([]*Robot, 0)
	for scanner.Scan() {
		robot := NewRobotFromString(scanner.Text())
		robots = append(robots, robot)
	}
	return robots
}

func testPart(t *testing.T, fn func(input []*Robot) int, expected int) {
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
	const expected = 12
	inputTest := createInput("test.txt")
	answer := Part1(inputTest, NewCoordinate(11, 7))
	if answer != expected {
		t.Errorf("Expected %d, got %d", expected, answer)
		return
	}
	inputPuzzle := createInput("puzzle.txt")
	answer = Part1(inputPuzzle, NewCoordinate(101, 103))
	t.Logf("Answer: %d", answer)
}

func TestPart2(t *testing.T) {
	inputPuzzle := createInput("puzzle.txt")
	answer := Part2(inputPuzzle, NewCoordinate(101, 103))
	t.Logf("Answer: %d", answer)
}

func TestVisualizePart2(t *testing.T) {
	inputPuzzle := createInput("puzzle.txt")
	VisualizePart2(inputPuzzle, NewCoordinate(101, 103), 500)
}

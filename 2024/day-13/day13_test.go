package day13

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func makeButton(s string) [2]int64 {
	s = s[12:]
	split := strings.Split(s, ",")
	f, s := split[0], split[1][3:]
	ans := [2]int64{}
	t1, _ := strconv.Atoi(f)
	t2, _ := strconv.Atoi(s)
	ans[0] = int64(t1)
	ans[1] = int64(t2)
	return ans
}

func makePrize(s string) [2]int64 {
	s = s[9:]
	split := strings.Split(s, ",")
	f, s := split[0], split[1][3:]
	ans := [2]int64{}
	t1, _ := strconv.Atoi(f)
	t2, _ := strconv.Atoi(s)
	ans[0] = int64(t1)
	ans[1] = int64(t2)
	return ans
}

func createInput(filename string) [][3][2]int64 {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))
	input := make([][3][2]int64, 0)
	var prev []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			prev = append(prev, line)
		} else {
			query := [3][2]int64{}
			query[0] = makeButton(prev[0])
			query[1] = makeButton(prev[1])
			query[2] = makePrize(prev[2])
			input = append(input, query)
			prev = prev[3:]
		}
	}
	query := [3][2]int64{}
	query[0] = makeButton(prev[0])
	query[1] = makeButton(prev[1])
	query[2] = makePrize(prev[2])
	input = append(input, query)
	return input
}

func testPart(t *testing.T, fn func(input [][3][2]int64) int64, expected int64) {
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
	testPart(t, Part1, 480)
}

func TestPart2(t *testing.T) {
	inputPuzzle := createInput("puzzle.txt")
	answer := Part2(inputPuzzle)
	t.Logf("Answer: %d", answer)
}

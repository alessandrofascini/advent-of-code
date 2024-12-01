package day_01

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func createInput(t *testing.T, filename string) [][]int {
	b, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	input := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		a, _ := strconv.Atoi(split[0])
		b, _ := strconv.Atoi(split[1])
		input = append(input, []int{a, b})
	}
	return input
}

func TestPart1(t *testing.T) {
	input := createInput(t, "puzzle.txt")
	answer := Part1(input)
	t.Log(answer)
}

func TestPart2(t *testing.T) {
	input := createInput(t, "puzzle.txt")
	answer := Part2(input)
	t.Log(answer)
}

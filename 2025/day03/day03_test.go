package day03

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func ReadInput(t *testing.T, filename string) []Bank {
	b, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	input := make([]Bank, 0)
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, "")
		bank := make(Bank, len(splitted))

		for i, s := range splitted {
			bank[i], _ = strconv.ParseInt(s, 10, 64)
		}

		input = append(input, bank)
	}
	return input
}

func TestPart1(t *testing.T) {
	input := ReadInput(t, "test.txt")
	t.Log("test:", Part1(input))
	input = ReadInput(t, "puzzle.txt")
	t.Log("puzzle:", Part1(input))
}

func TestPart2(t *testing.T) {
	input := ReadInput(t, "test.txt")
	t.Log("test:", Part2(input), Part2(input) == 3121910778619)
	input = ReadInput(t, "puzzle.txt")
	t.Log("puzzle:", Part2(input), Part2(input) == 168575096286051)
}

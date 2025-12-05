package day05

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestInterval_Compare(t *testing.T) {
	base := NewInterval(5, 10)
	compares := [][]int{
		{base.Compare(NewInterval(2, 3)), -1},
		{base.Compare(NewInterval(2, 6)), 0},
		{base.Compare(NewInterval(6, 8)), 0},
		{base.Compare(NewInterval(8, 11)), 0},
		{base.Compare(NewInterval(11, 12)), 1},
	}

	for _, compare := range compares {
		value, expected := compare[0], compare[1]
		if value != expected {
			t.Error("Expected", expected, "Got", value)
		}
	}

}

func ReadInput(t *testing.T, filename string) ([]Interval, []IngredientId) {
	b, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	ranges := make([]Interval, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, "-")
		lower, _ := strconv.Atoi(split[0])
		upper, _ := strconv.Atoi(split[1])
		ranges = append(ranges, NewInterval(lower, upper))
	}

	ingredients := make([]IngredientId, 0)
	for scanner.Scan() {
		line := scanner.Text()
		id, _ := strconv.Atoi(line)
		ingredients = append(ingredients, id)
	}

	return ranges, ingredients
}

func TestPart1(t *testing.T) {
	ranges, ingredients := ReadInput(t, "test.txt")
	t.Log("Test Part 1:", Part1(ranges, ingredients))

	ranges, ingredients = ReadInput(t, "puzzle.txt")
	t.Log("Puzzle Part 1:", Part1(ranges, ingredients))
}

func TestPart2(t *testing.T) {
	ranges, _ := ReadInput(t, "test.txt")
	t.Log("Test Part 2:", Part2(ranges))

	ranges, _ = ReadInput(t, "puzzle.txt")
	t.Log("Puzzle Part 2:", Part2(ranges))
}

package day05

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func ReadInput(t *testing.T, filename string) ([]Range, []IngredientId) {
	b, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	ranges := make([]Range, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, "-")
		lower, _ := strconv.Atoi(split[0])
		upper, _ := strconv.Atoi(split[1])
		ranges = append(ranges, Range{lower, upper})
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
}

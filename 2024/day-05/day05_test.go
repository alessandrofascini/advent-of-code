package day05

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func createInput(t *testing.T, filename string) (map[int]map[int]struct{}, [][]int) {
	b, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))

	adjGraph := make(map[int]map[int]struct{})

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, "|")
		u, _ := strconv.Atoi(split[0])
		v, _ := strconv.Atoi(split[1])
		if _, ok := adjGraph[u]; !ok {
			adjGraph[u] = map[int]struct{}{}
		}
		adjGraph[u][v] = struct{}{}
	}

	updates := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		numbers := make([]int, len(split))
		for i, s := range split {
			numbers[i], _ = strconv.Atoi(s)
		}
		updates = append(updates, numbers)
	}

	return adjGraph, updates
}

func TestPart1(t *testing.T) {
	adj, updates := createInput(t, "test.txt")
	answer := Part1(adj, updates)
	t.Log(answer)
}

func TestPart1Puzzle(t *testing.T) {
	adj, updates := createInput(t, "puzzle.txt")
	answer := Part1(adj, updates)
	t.Log(answer)
}

func TestPart2(t *testing.T) {
	adj, updates := createInput(t, "test.txt")
	answer := Part2(adj, updates)
	t.Log(answer)
}

func TestPart2Puzzle(t *testing.T) {
	adj, updates := createInput(t, "puzzle.txt")
	answer := Part2(adj, updates)
	t.Log(answer)
}

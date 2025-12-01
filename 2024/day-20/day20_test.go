package day_20

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func createInput(filename string) [][]byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bytes.Split(b, []byte("\n"))
}

func TestPart1(t *testing.T) {
	tcs := []struct {
		name string
		n    int
		want int
	}{
		{"test", 2, 44},
		{"test", 4, 30},
		{"test", 6, 16},
		{"test", 8, 14},
		{"test", 10, 10},
		{"test", 12, 8},
		{"test", 20, 5},
		{"test", 36, 4},
		{"test", 38, 3},
		{"test", 40, 2},
		{"test", 64, 1},
		{"puzzle", 100, 1422},
	}
	for _, tc := range tcs {
		grid := createInput(fmt.Sprintf("%s.txt", tc.name))
		answer := Part1(grid, tc.n)
		if answer != tc.want {
			t.Errorf("Expected %d, got %d", tc.want, answer)
		} else {
			t.Logf("Part1(grid, %3d): %d", tc.n, answer)
		}
	}
}

func TestPart2(t *testing.T) {
	tcs := []struct {
		name string
		n    int
		want int
	}{
		{"test", 76, 3},
		{"test", 74, 7},
		{"test", 50, 285},
		{"puzzle", 100, 0},
	}
	for _, tc := range tcs {
		grid := createInput(fmt.Sprintf("%s.txt", tc.name))
		answer := Part1(grid, tc.n)
		if answer != tc.want {
			t.Errorf("Expected %d, got %d", tc.want, answer)
		} else {
			t.Logf("Part1(grid, %3d): %d", tc.n, answer)
		}
	}
}

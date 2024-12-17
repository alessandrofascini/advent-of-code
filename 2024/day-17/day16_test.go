package day17

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func getRegister(s string) int {
	split := strings.Split(s, ": ")
	v, _ := strconv.Atoi(split[1])
	return v
}

func getProgram(s string) []int {
	split := strings.Split(s, ": ")
	split = strings.Split(split[1], ",")
	program := make([]int, len(split))
	for i, v := range split {
		program[i], _ = strconv.Atoi(v)
	}
	return program
}

func createInput(filename string) (int, int, int, []int) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))
	scanner.Scan()
	A := getRegister(scanner.Text())
	scanner.Scan()
	B := getRegister(scanner.Text())
	scanner.Scan()
	C := getRegister(scanner.Text())
	scanner.Scan()
	scanner.Text()
	scanner.Scan()
	program := getProgram(scanner.Text())
	return A, B, C, program
}

func testPart(t *testing.T, fn func(A, B, C int, program []int) string, expected ...string) {
	if len(expected) < 1 {
		t.Fatalf("no expected values; required at least one value")
	}
	A, B, C, program := createInput("test.txt")
	answer := fn(A, B, C, program)
	if answer != expected[0] {
		t.Errorf("Test Expected %s, got %s", expected[0], answer)
		return
	}
	A, B, C, program = createInput("puzzle.txt")
	answer = fn(A, B, C, program)
	if len(expected) > 1 && answer != expected[1] {
		t.Errorf("Puzzle Expected %s, got %s", expected[1], answer)
		return
	} else {
		t.Logf("Puzzle Answer %s", answer)
	}
}

func TestPart1(t *testing.T) {
	testPart(t, Part1, "4,6,3,5,6,3,5,2,1,0")
}

func TestPart2(t *testing.T) {
	const expected = 117440
	A, B, C, program := createInput("test2.txt")
	answer := Part2(A, B, C, program)
	if answer != expected {
		t.Errorf("Test Expected %d, got %d", expected, answer)
		return
	}
	A, B, C, program = createInput("puzzle.txt")
	answer = Part2(A, B, C, program)
	t.Logf("Puzzle Answer %d", answer)
}

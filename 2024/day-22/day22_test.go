package day22

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func createInput(filename string) []int {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	numbers := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
	}
	return numbers
}

func TestNextSecretNumber(t *testing.T) {
	secret := int64(123)
	expected := []int64{
		15887950,
		16495136,
		527345,
		704524,
		1553684,
		12683156,
		11100544,
		12249484,
		7753432,
		5908254,
	}

	for _, v := range expected {
		old := secret
		secret = NextSecretNumber(secret)
		if secret != v {
			t.Fatalf("NextNSecretNumber(%d) = %d, want %d", old, secret, v)
		}
	}
}

func TestPart1(t *testing.T) {
	const expected = 37327623
	numbers := createInput("test.txt")
	answer := Part1(numbers)
	if answer != expected {
		t.Fatalf("Expected %d, got %d", expected, answer)
	}
	numbers = createInput("puzzle.txt")
	answer = Part1(numbers)
	t.Logf("Puzzle Answer: %d", answer)
}

func TestPart2_1(t *testing.T) {
	secret := int64(1)
	digit := secret % 10
	fmt.Printf("%12d: %d\n", secret, digit)
	for i := 0; i < 2000; i++ {
		next := NextSecretNumber(secret)
		nd := next % 10
		fmt.Printf("%12d: %d (%d)\n", next, nd, nd-digit)
		secret, digit = next, nd
	}
}

func TestPart2(t *testing.T) {
	const expected = 23
	numbers := createInput("test2.txt")
	answer := Part2(numbers)
	if answer != expected {
		t.Fatalf("Expected %d, got %d", expected, answer)
	}
	numbers = createInput("puzzle.txt")
	answer = Part2(numbers)
	t.Logf("Puzzle Answer: %d", answer)
}

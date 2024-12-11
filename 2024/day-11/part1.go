package day11

import "fmt"

func digitsOf(n int) int {
	digits := 0
	for n > 0 {
		n /= 10
		digits++
	}
	return digits
}

func splitNumber(n, d int) (int, int) {
	s := 0
	p10 := 1
	d = d >> 1
	for i := 0; i < d; i++ {
		s += n % 10 * p10
		n /= 10
		p10 *= 10
	}
	return n, s
}

func nextStone(v int) []int {
	if v == 0 {
		return []int{1}
	}
	// digits are even
	if d := digitsOf(v); d&1 == 0 {
		p, s := splitNumber(v, d)
		return []int{p, s}
	}
	return []int{v * 2024}
}

func predict(stone, n int, mem map[string]int) int {
	makeKey := func(i, j int) string { return fmt.Sprintf("%d_%d", i, j) }
	var rec func(s, p int) int
	rec = func(s, p int) int {
		if p == n {
			return 1
		}
		key := makeKey(s, p)
		if v, ok := mem[key]; ok {
			return v
		}
		ns := nextStone(s)
		nOfStone := 0
		for _, v := range ns {
			nOfStone += rec(v, p+1)
		}
		mem[key] = nOfStone
		return nOfStone
	}
	return rec(stone, 0)
}

func blinkNTimes(stones []int, n int) int {
	ans := 0
	mem := map[string]int{}
	for _, stone := range stones {
		ans += predict(stone, n, mem)
	}
	return ans
}

func Part1(stones []int) int {
	return blinkNTimes(stones, 25)
}

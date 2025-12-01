package day22

func mix(a, b int64) int64 {
	return a ^ b
}

func prune(a int64) int64 {
	return a % 16777216
}

func NextSecretNumber(n int64) int64 {
	n = prune(mix(n, n*64))
	n = prune(mix(n, n/32))
	return prune(mix(n, n*2048))
}

func NextNSecretNumbers(from int64, times int) int64 {
	for times > 0 {
		from = NextSecretNumber(from)
		times--
	}
	return from
}

func Part1(input []int) int64 {
	sum := int64(0)
	for _, n := range input {
		sum += NextNSecretNumbers(int64(n), 2000)
	}
	return sum
}

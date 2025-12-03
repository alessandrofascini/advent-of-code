package day03

type Bank = []int64

func joltage(b Bank, m int) int64 {
	n := len(b)
	if n < m {
		return -1
	}
	dp := make([]int64, n)
	for i := 1; i <= m; i++ {
		temp := make([]int64, n)
		temp[i-1] = dp[i-1]*10 + b[i-1]
		for j := i; j < n; j++ {
			temp[j] = max(temp[j-1], dp[j-1]*10+b[j])
		}

		// copy
		for i, v := range temp {
			dp[i] = v
		}
	}
	return dp[n-1]
}

func Part1(bs []Bank) int64 {
	ans := int64(0)
	for _, b := range bs {
		ans += joltage(b, 2)
	}
	return ans
}

func Part2(bs []Bank) int64 {
	ans := int64(0)
	for _, b := range bs {
		ans += joltage(b, 12)
	}
	return ans
}

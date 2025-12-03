package day03

type Bank = []int64

func Joltage2(b Bank, m int) int64 {
	n := len(b)
	if n < m {
		return -1
	}
	dp := make([][]int64, m+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = max(dp[i][j-1], dp[i-1][j-1]*10+b[j-1])
		}
	}
	return dp[m][n]
}

func Joltage(b Bank) int64 {
	l, r := int64(0), int64(0)
	ans := int64(0)
	for _, v := range b {
		cases := [][]int64{{l, v}, {r, v}}
		for _, c := range cases {
			nl, nr := c[0], c[1]
			if t := nl*10 + nr; t > ans {
				l, r, ans = nl, nr, t
			}
		}
	}
	return ans
}

func Part1(bs []Bank) int64 {
	ans := int64(0)
	for _, b := range bs {
		ans += Joltage2(b, 2)
	}
	return ans
}

func Part2(bs []Bank) int64 {
	ans := int64(0)
	for _, b := range bs {
		ans += Joltage2(b, 12)
	}
	return ans
}

package day07

func concat(s, v int) int {
	cv := v
	powerOfTen := 1
	for cv > 0 {
		cv /= 10
		powerOfTen *= 10
	}
	return s*powerOfTen + v
}

func achieveGoal2(numbers []int, goal int) bool {
	var rec func(i, s int) bool
	rec = func(i, s int) bool {
		if i == len(numbers) {
			return s == goal
		}
		if s > goal {
			return false
		}
		v := numbers[i]
		return rec(i+1, s+v) || rec(i+1, s*v) || rec(i+1, concat(s, v))
	}
	return rec(0, 0)
}

func Part2(equations [][]int) int {
	ans := 0
	for _, equation := range equations {
		if goal := equation[0]; achieveGoal2(equation[1:], goal) {
			ans += goal
		}
	}
	return ans
}

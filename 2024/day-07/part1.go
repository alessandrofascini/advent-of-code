package day07

func achieveGoal(numbers []int, goal int) bool {
	var rec func(i, s int) bool
	rec = func(i, s int) bool {
		if i == len(numbers) {
			return s == goal
		}
		if s > goal {
			return false
		}
		v := numbers[i]
		return rec(i+1, s+v) || rec(i+1, s*v)
	}
	return rec(0, 0)
}

func Part1(equations [][]int) int {
	ans := 0
	for _, equation := range equations {
		if goal := equation[0]; achieveGoal(equation[1:], goal) {
			ans += goal
		}
	}
	return ans
}

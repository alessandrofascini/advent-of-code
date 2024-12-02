package day_02

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isSafe(l []int) bool {
	if len(l) < 2 {
		return true
	}
	var compare func(int, int) bool
	if l[0] < l[1] {
		compare = func(x, y int) bool {
			return x < y
		}
	} else {
		compare = func(x int, y int) bool {
			return x > y
		}
	}
	for i := 1; i < len(l); i++ {
		x, y := l[i-1], l[i]
		diff := abs(x - y)
		if !(compare(x, y) && 0 < diff && diff < 4) {
			return false
		}
	}
	return true
}

func Part1(levels [][]int) int {
	safe := 0
	for _, l := range levels {
		if isSafe(l) {
			safe++
		}
	}
	return safe
}

func isSafeWithDel(report []int) bool {
	for i := range report {
		temp := make([]int, len(report))
		copy(temp, report)
		if isSafe(append(temp[:i], temp[i+1:]...)) {
			return true
		}
	}
	return false
}

func Part2(reports [][]int) int {
	safe := 0
	for _, r := range reports {
		if isSafe(r) || isSafeWithDel(r) {
			safe++
		}
	}
	return safe
}

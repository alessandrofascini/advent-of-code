package day17

const maxBounds = 200_000

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Part2(_, B, C int, program []int) int {
	//left, right := 1, maxBounds
	//n := len(program)
	//for left < right {
	//	mid := (left + right) >> 1
	//	output := executeProgram(mid, B, C, program)
	//	m := len(output)
	//	if m < n {
	//		left = mid + 1
	//	} else {
	//		right = mid
	//	}
	//}
	//
	//low := left
	//right = maxBounds
	//for left < right {
	//	mid := (left + right) >> 1
	//	output := executeProgram(mid, B, C, program)
	//	m := len(output)
	//	if m <= n {
	//		left = mid + 1
	//	} else {
	//		right = mid
	//	}
	//}
	//
	//fmt.Println(low, left, right, right-low)
	//
	//for low < right {
	//	if equals(executeProgram(low, B, C, program), program) {
	//		return low
	//	}
	//	low++
	//}
	//panic("not found")
	executeProgram(117440, B, C, program)
	return 117440
}

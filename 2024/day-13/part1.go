package day13

const (
	btnACost = 3
	btnBCost = 1
)

func mcm(a, b int64) int64 {
	n := a * b
	for b != 0 {
		x := b
		b = a % b
		a = x
	}
	return n / a
}

func tryToWin(a, b, goal [2]int64) int64 {
	v1, v2, x := a[0], b[0], goal[0]
	v3, v4, y := a[1], b[1], goal[1]

	n := mcm(v1, v3)
	n1 := n / v1
	v2, x = v2*n1, x*n1

	n2 := n / v3
	v4, y = v4*n2, y*n2

	betaD, betaN := v2-v4, x-y
	if betaN < 0 && betaD < 0 {
		betaN, betaD = -betaN, -betaD
	} else if !(betaN >= 0 && betaD >= 0) {
		return 0
	}
	if betaN%betaD != 0 {
		return 0
	}
	beta := betaN / betaD

	alfaN, alfaD := goal[0]-beta*b[0], a[0]
	if alfaN < 0 && alfaD < 0 {
		alfaN, alfaD = -alfaN, -alfaD
	} else if !(alfaN >= 0 && alfaD >= 0) {
		return 0
	}
	if alfaN%alfaD != 0 {
		return 0
	}
	alfa := alfaN / alfaD
	return alfa*btnACost + beta*btnBCost
}

func Part1(machines [][3][2]int64) int64 {
	tokens := int64(0)
	for _, machine := range machines {
		btnA, btnB, goal := machine[0], machine[1], machine[2]
		tokens += tryToWin(btnA, btnB, goal)
	}
	return tokens
}

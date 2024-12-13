package day13

const offset = 10000000000000

func Part2(machines [][3][2]int64) int64 {
	tokens := int64(0)
	for _, machine := range machines {
		btnA, btnB, goal := machine[0], machine[1], machine[2]
		goal[0], goal[1] = goal[0]+offset, goal[1]+offset
		tokens += tryToWin(btnA, btnB, goal)
	}
	return tokens
}

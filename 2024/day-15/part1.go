package day15

const (
	wall  = '#'
	empty = '.'
	robot = '@'
	box   = 'O'

	up    = '^'
	right = '>'
	down  = 'v'
	left  = '<'
)

var dirs = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func moveToDir(move byte) []int {
	switch move {
	case up:
		return dirs[0]
	case right:
		return dirs[1]
	case down:
		return dirs[2]
	case left:
		return dirs[3]
	}
	panic("invalid move")
}

//func makeKey(i, j int) string {
//	return fmt.Sprintf("%d;%d", i, j)
//}

func findRobot(g [][]byte) (int, int) {
	for i, row := range g {
		for j, c := range row {
			if c == robot {
				return i, j
			}
		}
	}
	panic("missing robot")
}

func tryToMove(i, j int, move byte, g [][]byte) (int, int) {
	dir := moveToDir(move)
	ni, nj := i+dir[0], j+dir[1]
	symbol := g[ni][nj]
	if symbol == wall {
		return i, j
	}
	if symbol == empty {
		return ni, nj
	}

	for {
		ni, nj = ni+dir[0], nj+dir[1]
		symbol = g[ni][nj]
		if symbol == wall || symbol == empty {
			break
		}
	}
	if symbol == wall {
		return i, j
	}
	// symbol empty, can move
	for !(ni == i && nj == j) {
		pi, pj := ni-dir[0], nj-dir[1]
		g[ni][nj], g[pi][pj] = g[pi][pj], g[ni][nj]
		ni, nj = pi, pj
	}

	return i + dir[0], j + dir[1]
}

func moveBoxes(g [][]byte, movements []byte) {
	i, j := findRobot(g)
	g[i][j] = empty

	for _, move := range movements {
		i, j = tryToMove(i, j, move, g)
	}
}

func gps(g [][]byte) int {
	answer := 0
	for i, row := range g {
		for j, c := range row {
			if c == box || c == boxOpen {
				answer += 100*i + j
			}
		}
	}
	return answer
}

func Part1(g [][]byte, movements []byte) int {
	moveBoxes(g, movements)
	return gps(g)
}

package day06

const (
	empty = '.'
	wall  = '#'
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

func isInside(i, j int, g [][]byte) bool {
	n, m := len(g), len(g[0])
	return -1 < i && i < n && -1 < j && j < m
}

func getDir(i, j int, g [][]byte) int {
	if !isInside(i, j, g) {
		return -1
	}
	switch g[i][j] {
	case up:
		return 0
	case right:
		return 1
	case down:
		return 2
	case left:
		return 3
	}
	return -1
}

func findGuardPosition(g [][]byte) (int, int) {
	for i, row := range g {
		for j, cell := range row {
			if cell != wall && cell != empty {
				return i, j
			}
		}
	}
	panic("cannot find guard position")
}

func Part1(g [][]byte) int {
	i, j := findGuardPosition(g)
	move := getDir(i, j, g)

	visited := make([][]bool, len(g))
	for r := range visited {
		visited[r] = make([]bool, len(g[0]))
	}

	ans := 0
	for isInside(i, j, g) {
		if !visited[i][j] {
			visited[i][j] = true
			ans++
		}
		ni, nj := i+dirs[move][0], j+dirs[move][1]
		if !isInside(ni, nj, g) {
			break
		}
		if g[ni][nj] == wall {
			move = (move + 1) % 4
		} else {
			i, j = ni, nj
		}
	}

	return ans
}

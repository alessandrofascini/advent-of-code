package day23

import "fmt"

const (
	tree        = '#'
	path        = '.'
	upSlopes    = '^'
	leftSlopes  = '<'
	rightSlopes = '>'
	downSlopes  = 'v'

	up    = 0
	right = 1
	down  = 2
	left  = 3
)

func findStartPoint(grid [][]byte) (int, int) {
	for i := range grid {
		if grid[0][i] == path {
			return 0, i
		}
	}
	panic("missing start")
}

func findEndPoint(grid [][]byte) (int, int) {
	n := len(grid) - 1
	for i := range grid {
		if grid[n][i] == path {
			return n, i
		}
	}
	panic("missing end")
}

func makeKey(a, b int) string {
	return fmt.Sprintf("%d-%d", a, b)
}

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func isInside(i, j int, grid [][]byte) bool {
	return -1 < i && i < len(grid) && -1 < j && j < len(grid[0])
}

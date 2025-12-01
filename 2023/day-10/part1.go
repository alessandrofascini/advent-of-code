package day202310

import "fmt"

const (
	up    = 0
	right = 1
	down  = 2
	left  = 3
)

var dirs = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func findStart(g [][]byte) (int, int) {
	for i, row := range g {
		for j, c := range row {
			if c == start {
				return i, j
			}
		}
	}
	panic("cannot find start")
}

func isInside(i, j int, g [][]byte) bool {
	return -1 < i && i < len(g) && -1 < j && j < len(g[0])
}

const (
	ns     = '|' // north-south
	ew     = '-' // east-west
	ne     = 'L' // north-east
	nw     = 'J' // north-west
	sw     = '7' // south-west
	se     = 'F' // south-east
	ground = '.' // ground
	start  = 'S' // start point
)

func canGo(fi, fj, ti, tj, direction int, g [][]byte) bool {
	if !isInside(fi, fj, g) || !isInside(ti, tj, g) {
		return false
	}
	fc, tc := g[fi][fj], g[ti][tj]
	if fc == ground || tc == ground {
		return false
	}
	if fc == ns {
		if direction == up {
			return tc == ns || tc == sw || tc == se
		}
		if direction == down {
			return tc == ns || tc == ne || tc == nw
		}
		return false
	}
	if fc == ew {
		if direction == right {
			return tc == ew || tc == nw || tc == sw
		}
		if direction == left {
			return tc == ew || tc == ne || tc == se
		}
		return false
	}
	if fc == ne {
		if direction == up {
			return tc == ns || tc == sw || tc == se
		}
		if direction == right {
			return tc == ew || tc == nw || tc == sw
		}
		return false
	}
	if fc == nw {
		if direction == up {
			return tc == ns || tc == sw || tc == se
		}
		if direction == left {
			return tc == ew || tc == ne || tc == se
		}
		return false
	}
	if fc == sw {
		if direction == left {
			return tc == ew || tc == ne || tc == se
		}
		if direction == down {
			return tc == ns || tc == ne || tc == nw
		}
		return false
	}
	if fc == se {
		if direction == right {
			return tc == ew || tc == nw || tc == sw
		}
		if direction == down {
			return tc == ns || tc == nw || tc == ne
		}
		return false
	}
	fmt.Println(fc, tc, direction)
	panic("unreachable code")
}

func findFurthestPoint(si, sj int, g [][]byte) int {
	visited := make([][]int, len(g))
	for i := range visited {
		visited[i] = make([]int, len(g[0]))
		for j := range visited[i] {
			visited[i][j] = -1
		}
	}
	length := 0
	queue := [][]int{{si, sj, 0}}
	for len(queue) > 0 {
		i, j, h := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]
		if visited[i][j] != -1 {
			length = max(length, visited[i][j])
			continue
		}
		visited[i][j] = h
		counter := 0
		for d, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if canGo(i, j, ni, nj, d, g) {
				queue = append(queue, []int{ni, nj, h + 1})
				counter++
			}
		}
		if counter != 2 {
			return 0
		}
	}
	return length
}

func Part1(g [][]byte) int {
	si, sj := findStart(g)

	ans := 0
	options := []byte{ns, ew, ne, nw, sw, se}
	for _, option := range options {
		g[si][sj] = option
		ans = max(ans, findFurthestPoint(si, sj, g))
	}

	return ans
}

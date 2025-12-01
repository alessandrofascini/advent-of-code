package day202310

import "fmt"

const (
	outside = 'O'
	pipe    = 'P'
)

func replaceWithPipe(si, sj int, g [][]byte) [][]byte {
	queue := [][]int{{si, sj}}

	visited := make([][]byte, len(g))
	for i := range visited {
		visited[i] = make([]byte, len(g[i]))
		for j := range visited[i] {
			visited[i][j] = ground
		}
	}

	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]
		if visited[i][j] != ground {
			continue
		}
		for d, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if canGo(i, j, ni, nj, d, g) {
				queue = append(queue, []int{ni, nj})
			}
		}
		visited[i][j] = g[i][j]
	}

	return visited
}

func countInside(g [][]byte) int {
	inside := 0
	for _, r := range g {
		fmt.Print(string(r))
		for _, v := range r {
			if v == ground {
				inside++
			}
		}
		fmt.Println(" ", inside)
	}
	return inside >> 1
}

func findOutside(g [][]byte) {
	n, m := len(g), len(g[0])
	var rec func(i, j int)
	rec = func(i, j int) {
		if !isInside(i, j, g) || g[i][j] != ground {
			return
		}
		g[i][j] = outside
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			rec(ni, nj)
		}
	}
	// horizontal
	for i := 0; i < m; i++ {
		f, s := g[0][i], g[n-1][i]
		if f == ground {
			rec(0, i)
		}
		if s == ground {
			rec(n-1, i)
		}
	}
	// vertical
	for i := 1; i < n-1; i++ {
		f, s := g[i][0], g[i][m-1]
		if f == ground {
			rec(i, 0)
		}
		if s == ground {
			rec(i, m-1)
		}
	}
}

func expandGrid(g [][]byte) [][]byte {
	n, m := len(g), len(g[0])
	expandedGrid := make([][]byte, n<<1)
	for i := range expandedGrid {
		expandedGrid[i] = make([]byte, m<<1)
		for j := range expandedGrid[i] {
			expandedGrid[i][j] = ground
		}
	}

	for i, r := range g {
		di := i << 1
		for j := range r {
			dj := j << 1
			expandedGrid[di][dj] = g[i][j]
		}
	}

	for i, r := range expandedGrid {
		for j, v := range r {
			switch v {
			case '7':
				expandedGrid[i][j-1] = '-'
				expandedGrid[i+1][j] = '|'
			case 'F':
				expandedGrid[i][j+1] = '-'
				expandedGrid[i+1][j] = '|'
			case 'J':
				expandedGrid[i][j-1] = '-'
				expandedGrid[i-1][j] = '|'
			case 'L':
				expandedGrid[i][j+1] = '-'
				expandedGrid[i-1][j] = '|'
			}
		}
	}

	for i, row := range expandedGrid {
		for j := 1; j < len(row)-1; j++ {
			if row[j] == ground && row[j-1] == '-' && row[j+1] == '-' {
				expandedGrid[i][j] = '-'
			}
		}
	}

	for j := 1; j < len(expandedGrid[0])-1; j++ {
		for i := 1; i < len(expandedGrid)-1; i++ {
			if expandedGrid[i][j] == '.' && expandedGrid[i-1][j] == '|' && expandedGrid[i+1][j] == '|' {
				expandedGrid[i][j] = '|'
			}
		}
	}

	return expandedGrid
}

func tilesEnclosed(si, sj int, g [][]byte) int {
	// replace values with pipe char
	g = replaceWithPipe(si, sj, g)
	for _, r := range g {
		fmt.Println(string(r))
	}
	fmt.Println()
	// expandGrid
	g = expandGrid(g)
	// exclude outside
	findOutside(g)
	// count inside
	return countInside(g)
}

func Part2(g [][]byte) int {
	si, sj := findStart(g)

	options := []byte{ns, ew, ne, nw, sw, se}
	for _, option := range options {
		g[si][sj] = option
		if pipeLength := findFurthestPoint(si, sj, g); pipeLength > 0 {
			return tilesEnclosed(si, sj, g)
		}
	}
	panic("unreachable")
}

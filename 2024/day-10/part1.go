package day10

import "fmt"

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func inside(i, j int, g [][]byte) bool {
	return -1 < i && i < len(g) && -1 < j && j < len(g[0])
}

func findScore(si, sj int, m [][]byte) int {
	coords := map[string]struct{}{}
	queue := [][]int{{si, sj}}
	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]
		if !inside(i, j, m) {
			continue
		}
		if m[i][j] == '9' {
			key := fmt.Sprintf("%d,%d", i, j)
			coords[key] = struct{}{}
			continue
		}
		sc := m[i][j] + 1
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if inside(ni, nj, m) && m[ni][nj] == sc {
				queue = append(queue, []int{ni, nj})
			}
		}
	}
	return len(coords)
}

func Part1(m [][]byte) int {
	score := 0
	for i, r := range m {
		for j, c := range r {
			if c == '0' {
				score += findScore(i, j, m)
			}
		}
	}
	return score
}

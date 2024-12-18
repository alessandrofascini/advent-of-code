package day18

import (
	"math"
)

const (
	empty     = '.'
	corrupted = '#'
)

func isInside(i, j, n int) bool {
	return -1 < i && i < n && -1 < j && j < n
}

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func Part1(n int, points [][]int) int {
	n = n + 1
	memory := make([][]byte, n)
	for i := range memory {
		memory[i] = make([]byte, n)
		for j := range memory[i] {
			memory[i][j] = empty
		}
	}

	for _, point := range points {
		i, j := point[0], point[1]
		memory[i][j] = corrupted
	}

	distance := make([][]int, n)
	for i := range distance {
		distance[i] = make([]int, n)
		for j := range distance[i] {
			distance[i][j] = math.MaxInt
		}
	}

	queue := [][]int{{0, 0, 0}}

	for len(queue) > 0 {
		i, j, w := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]
		if w >= distance[i][j] {
			continue
		}
		distance[i][j] = w
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if isInside(ni, nj, n) && memory[ni][nj] != corrupted {
				queue = append(queue, []int{ni, nj, w + 1})
			}
		}
	}

	n = n - 1
	return distance[n][n]
}

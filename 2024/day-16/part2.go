package day16

import (
	"fmt"
	"math"
)

func Part2(g [][]byte) int {
	si, sj := searchStartTile(g)
	ei, ej := searchEndTile(g)
	answer := math.MaxInt
	visited := map[string][]int{}
	queue := [][]int{{si, sj, 0, right, -1, -1}}
	for len(queue) > 0 {
		i, j, w, d := queue[0][0], queue[0][1], queue[0][2], queue[0][3]
		pi, pj := queue[0][4], queue[0][5]
		queue = queue[1:]

		key := makeKey(i, j)
		if i == ei && j == ej {
			if w < answer {
				answer = w
			}
			//if w == answer {
			//}
			visited[key] = []int{w, pi, pj, d}
			continue
		}
		if v, ok := visited[key]; g[i][j] == wall || (ok && w > v[0]) {
			continue
		}
		visited[key] = []int{w, pi, pj, d}
		for nd, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if d&1 == nd&1 {
				queue = append(queue, []int{ni, nj, w + 1, nd, i, j})
			} else {
				// clockwise or counterclockwise rotation
				queue = append(queue, []int{ni, nj, w + 1001, nd, i, j})
			}
		}
	}

	var dfs func(i, j, h int)
	dfs = func(i, j, h int) {
		if i == si && j == sj {
			fmt.Println(h + 1)
			return
		}
		key := makeKey(i, j)
		if _, ok := visited[key]; !ok {
			fmt.Println("should not be here")
			return
		}
		fmt.Println(key)
		v := visited[key]
		pi, pj := v[1], v[2]
		dfs(pi, pj, h+1)
	}
	dfs(ei, ej, 1)

	return 0
}

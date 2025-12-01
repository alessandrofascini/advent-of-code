package day_20

import (
	"fmt"
)

const (
	S     = 'S'
	E     = 'E'
	track = '.'
	wall  = '#'
)

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func findCharacter(g [][]byte, c byte) (int, int) {
	for i, r := range g {
		for j, v := range r {
			if v == c {
				return i, j
			}
		}
	}
	return -1, -1
}

func makeKey(i, j int) string {
	return fmt.Sprintf("(%d;%d)", i, j)
}

func isInside(i, j int, g [][]byte) bool {
	return -1 < i && i < len(g) && -1 < j && j < len(g[0])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// todo: change me
func findCandidates(fi, fj int, direction []int, g [][]byte, distances map[string]int) [][]int {
	c := make([][]int, 0, 4)
	ci, cj := fi+direction[0], fj+direction[1]
	ni, nj := ci+direction[0], cj+direction[1]
	if _, ok := distances[makeKey(ni, nj)]; !ok && isInside(ni, nj, g) && !(ni == fi && nj == fj) && g[ni][nj] == track {
		c = append(c, []int{fi, fj, ni, nj})
	}
	return c
}

func dfs(si, sj, sd int, g [][]byte, distances map[string]int) map[string][][]int {
	candidates := make(map[string][][]int)

	stack := [][]int{{si, sj, sd}}
	for len(stack) > 0 {
		last := len(stack) - 1
		i, j, d := stack[last][0], stack[last][1], stack[last][2]
		stack = stack[:last]

		key := makeKey(i, j)
		if _, ok := distances[key]; ok {
			continue
		}
		distances[key] = d

		for w := len(dirs) - 1; w > -1; w-- {
			ni, nj := i+dirs[w][0], j+dirs[w][1]
			if g[ni][nj] == track {
				stack = append(stack, []int{ni, nj, d + 1})
			} else if g[ni][nj] == wall {
				for _, c := range findCandidates(i, j, dirs[w], g, distances) {
					k := makeKey(i, j)
					candidates[k] = append(candidates[k], c)
				}
			}

		}
	}

	return candidates
}

func Part1(grid [][]byte, r int) int {
	si, sj := findCharacter(grid, S)
	if si == -1 {
		panic("Not found Start")
	}
	grid[si][sj] = track
	ei, ej := findCharacter(grid, E)
	if ei == -1 {
		panic("Not found End")
	}
	grid[ei][ej] = track

	distances := make(map[string]int)
	candidates := dfs(si, sj, 0, grid, distances)

	values := map[int]int{}

	for _, c1 := range candidates {
		for _, c := range c1 {
			fi, fj, ni, nj := c[0], c[1], c[2], c[3]
			fKey := makeKey(fi, fj)
			if _, ok := distances[fKey]; !ok {
				panic("not found 1")
			}
			nKey := makeKey(ni, nj)
			if _, ok := distances[nKey]; !ok {
				panic("not found 2")
			}
			delta := abs(distances[nKey]-distances[fKey]) - 2
			if delta > 0 {
				values[delta]++
			}
		}
	}

	ans := 0
	for k, ps := range values {
		if k >= r {
			ans += ps
		}
	}
	return ans
}

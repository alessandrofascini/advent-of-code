package day12

import "fmt"

var dirs = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

type region struct {
	symbol    byte
	area      int
	perimeter int
	sides     int
}

func (r *region) String() string {
	return fmt.Sprintf("%s: (%d; %d; %d)", string(r.symbol), r.area, r.perimeter, r.sides)
}

func makeKey(i, j int) string {
	return fmt.Sprintf("%d;%d", i, j)
}

func isInside(i, j int, g [][]byte) bool {
	return -1 < i && i < len(g) && -1 < j && j < len(g[0])
}

func newRegion(si, sj int, g [][]byte, visited map[string]struct{}) *region {
	reg := &region{}
	reg.symbol = g[si][sj]

	queue := [][]int{{si, sj}}
	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]
		key := makeKey(i, j)
		if _, ok := visited[key]; !isInside(i, j, g) || ok {
			continue
		}
		visited[key] = struct{}{}

		reg.area++
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if !isInside(ni, nj, g) || reg.symbol != g[ni][nj] {
				reg.perimeter++
			} else if reg.symbol == g[ni][nj] {
				queue = append(queue, []int{ni, nj})
			}
		}
	}

	return reg
}

func Part1(g [][]byte) int {
	visited := map[string]struct{}{}
	cost := 0
	for i, row := range g {
		for j, _ := range row {
			key := makeKey(i, j)
			if _, ok := visited[key]; !ok {
				r := newRegion(i, j, g, visited)
				fmt.Println(r)
				cost += r.area * r.perimeter
			}
		}
	}
	return cost
}

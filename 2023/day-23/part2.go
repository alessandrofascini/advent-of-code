package day23

import "fmt"

type Vertex struct {
	i, j int
}

func NewVertex(i int, j int) *Vertex {
	return &Vertex{i, j}
}

func (v *Vertex) String() string {
	return fmt.Sprintf("(%d, %d)", v.i, v.j)
}

type Edge struct {
	from, to *Vertex
	weight   int
}

func NewEdge(from, to *Vertex, weight int) *Edge {
	return &Edge{from, to, weight}
}

type Graph struct {
	edges map[string]*Edge
}

func NewGraph() (*Graph, int) {
	return nil, 0
}

func findIntersactionPoint(si, sj int, g [][]byte) [][]int {
	points := [][]int{{si, sj}}

	visited := map[string]struct{}{}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		key := makeKey(i, j)
		if _, ok := visited[key]; ok || !isInside(i, j, g) {
			return
		}
		visited[key] = struct{}{}

		connectedWith := 0
		nexts := make([][]int, 0, 4)
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if isInside(ni, nj, g) && g[ni][nj] == '.' {
				connectedWith++
				nexts = append(nexts, []int{ni, nj})
			}
		}

		if connectedWith >= 3 {
			points = append(points, []int{i, j})
		}

		for _, next := range nexts {
			dfs(next[0], next[1])
		}
	}

	dfs(si, sj)
	return points
}

func shortestPath(g *Graph) int {
	return 0
}

func existsPathWithAlmostWeight(g *Graph, requiredWeight int) int {
	return -1
}

func findLongestPath2(grid [][]byte) int {
	si, sj := findStartPoint(grid)
	ei, ej := findEndPoint(grid)

	// construct graph
	graph, ub := NewGraph()
	lb := shortestPath(graph)

	answer := lb
	for lb < ub {
		middle := (ub + lb) >> 1
		if w := existsPathWithAlmostWeight(graph, middle); w != -1 {
			lb = middle
			answer = w
		} else {
			ub = middle
		}
	}
	return answer
}

func Part2(grid [][]byte) int {
	// replace
	for i, r := range grid {
		for j, b := range r {
			if b != tree && b != path {
				grid[i][j] = path
			}
		}
	}
	return findLongestPath2(grid)
}

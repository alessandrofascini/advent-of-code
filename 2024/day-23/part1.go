package day23

import (
	"strings"
)

func createEdges(connections []string) [][]string {
	ans := make([][]string, len(connections))
	for i, connection := range connections {
		ans[i] = strings.Split(connection, "-")
	}
	return ans
}
func createEdge(g map[string][]string, u, v string) {
	g[u] = append(g[u], v)
}

func createGraph(edges [][]string) map[string][]string {
	graph := make(map[string][]string)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		createEdge(graph, u, v)
		createEdge(graph, v, u)
	}
	return graph
}

func calculatePossibilities(g map[string][]string, v string) int {
	ans := len(g[v]) - 1
	for _, u := range g[v] {
		if u != v {
			ans += len(g[u]) - 2
		}
	}
	return 2 * ans
}

func Part1(connections []string) int {
	graph := createGraph(createEdges(connections))

	ans := 0
	for key := range graph {
		if key[0] == 't' {
			ans += calculatePossibilities(graph, key)
		}
	}

	return ans
}

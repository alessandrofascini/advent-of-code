package day23

func findLongestPath1(grid [][]byte) int {
	si, sj := findStartPoint(grid)
	ei, ej := findEndPoint(grid)

	visited := map[string]struct{}{}

	var dfs func(i, j, w int) int
	dfs = func(i, j, w int) int {
		if ei == i && ej == j {
			return w
		}
		key := makeKey(i, j)
		if _, ok := visited[key]; ok || !isInside(i, j, grid) {
			return 0
		}
		visited[key] = struct{}{}

		answer := 0
		w = w + 1

		b := grid[i][j]
		isPath := b == path
		if isPath || b == upSlopes {
			ni, nj := i+dirs[up][0], j+dirs[up][1]
			answer = dfs(ni, nj, w)
		}
		if isPath || b == rightSlopes {
			ni, nj := i+dirs[right][0], j+dirs[right][1]
			answer = max(answer, dfs(ni, nj, w))
		}
		if isPath || b == downSlopes {
			ni, nj := i+dirs[down][0], j+dirs[down][1]
			answer = max(answer, dfs(ni, nj, w))
		}
		if isPath || b == leftSlopes {
			ni, nj := i+dirs[left][0], j+dirs[left][1]
			answer = max(answer, dfs(ni, nj, w))
		}
		delete(visited, key)

		return answer
	}

	return dfs(si, sj, 0)
}

func Part1(grid [][]byte) int {
	return findLongestPath1(grid)
}

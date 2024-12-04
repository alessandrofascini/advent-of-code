package day_04

func insideGrid(i, j int, grid [][]byte) bool {
	n, m := len(grid), len(grid[0])
	return -1 < i && i < n && -1 < j && j < m
}

func find(grid [][]byte, i, j int) int {
	if !insideGrid(i, j, grid) {
		return 0
	}
	dirs := [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	founded := 0
	for _, dir := range dirs {
		curr := make([]byte, 3)
		for w := 1; w < 4; w++ {
			ni, nj := i+w*dir[0], j+w*dir[1]
			if insideGrid(ni, nj, grid) {
				curr[w-1] = grid[ni][nj]
			}
		}
		if string(curr) == "MAS" {
			founded++
		}
	}
	return founded
}

func Part1(grid [][]byte) int {
	ans := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'X' {
				ans += find(grid, i, j)
			}
		}
	}
	return ans
}

func xFind(grid [][]byte, i, j int) bool {
	if !(insideGrid(i-1, j-1, grid) && insideGrid(i+1, j+1, grid)) {
		return false
	}
	if !((grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S') || (grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M')) {
		return false
	}
	return (grid[i+1][j-1] == 'M' && grid[i-1][j+1] == 'S') || (grid[i+1][j-1] == 'S' && grid[i-1][j+1] == 'M')
}

func Part2(grid [][]byte) int {
	n, m := len(grid)-1, len(grid[0])-1
	ans := 0
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if grid[i][j] == 'A' && xFind(grid, i, j) {
				ans++
			}
		}
	}
	return ans
}

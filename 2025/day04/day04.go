package day04

const NOTHING = '.'
const ROOL_OF_PAPER = '@'

type Row = []rune

var directions = [][]int{
	{-1,-1},
	{-1, 0},
	{-1, 1},
	{ 0,-1},
	{ 0, 1},
	{ 1,-1},
	{ 1, 0},
	{ 1, 1},
}

func canAccess(i, j int, grid []Row) int {
	n, m := len(grid), len(grid[0])
	count := 0
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if 0 <= ni && ni < n && 0 <= nj && nj < m && grid[ni][nj] == ROOL_OF_PAPER {
			count++
		}
	}

	return count
}


func Part1(grid []Row) int64 {
	ans := int64(0)

	for i, row := range grid {
		for j, cell := range row {
			if cell == ROOL_OF_PAPER && canAccess(i, j, grid) < 4 {
				ans++
			}	
		}
	}

	return ans
}

func Part2(grid []Row) int64 {
	ans := int64(0)

	for {
		t := int64(0)
		for i, row := range grid {
			for j, cell := range row {
				if cell == ROOL_OF_PAPER && canAccess(i, j, grid) < 4 {
					t++
					grid[i][j] = rune(NOTHING)
				}	
			}
		}

		if t == 0 {
			break
		}
		ans += t
	}

	return ans
}
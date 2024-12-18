package day18

import "fmt"

// optimize this solution with union find

func makeKey(i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

const (
	borderNorth = 0
	borderEast  = 1
	borderSouth = 2
	borderWest  = 3

	noBorderReached = 4
)

func parseBorder(i, j, n int) int {
	if (i == 0 && j == 0) || (i == 0 && j == n) || (i == n && j == 0) || (i == n && j == n) {
		return noBorderReached
	}
	if i == 0 {
		return borderNorth
	}
	if j == n {
		return borderEast
	}
	if j == 0 {
		return borderWest
	}
	if i == n {
		return borderSouth
	}
	panic("not belongs to the border")
}

func separateMemory(s, e int) bool {
	if s == borderNorth || s == borderEast {
		return e == borderWest || e == borderSouth
	}
	return e == borderNorth || e == borderEast
}

var allDirections = [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}

func foundBorder(curr []int, corruptedPoints map[string]struct{}, visited map[string]struct{}, n int) int {
	i, j := curr[0], curr[1]
	key := makeKey(i, j)
	if _, ok := visited[key]; ok {
		return noBorderReached
	}
	visited[key] = struct{}{}
	for _, dir := range allDirections {
		ni, nj := i+dir[0], j+dir[1]
		nk := makeKey(ni, nj)
		if _, ok := corruptedPoints[nk]; ok {
			if _, ok := visited[nk]; !ok && (ni == 0 || ni == n || nj == 0 || nj == n) {
				check := ni == 0 && (nj == 0 || nj == n)
				check = check || ni == n && (nj == 0 || nj == n)
				check = check || nj == 0 && ni == n
				check = check || nj == n && (ni == 0 || ni == n)
				if !check {
					return parseBorder(ni, nj, n)
				}
			}
			b := foundBorder([]int{ni, nj}, corruptedPoints, visited, n)
			if b != noBorderReached {
				return b
			}
		}
	}
	return noBorderReached
}

func canAchieveEnd(n int, borderPoints [][]int, corruptedPoints map[string]struct{}) bool {
	visited := make(map[string]struct{})
	for _, point := range borderPoints {
		sb := parseBorder(point[0], point[1], n)
		eb := foundBorder(point, corruptedPoints, visited, n)
		if separateMemory(sb, eb) {
			return true
		}
	}
	return false
}

func Part2(n int, points [][]int) string {
	borderPoints := make([][]int, 0)
	corruptedPoints := map[string]struct{}{}
	for _, point := range points {
		i, j := point[0], point[1]
		corruptedPoints[makeKey(i, j)] = struct{}{}
		if i == 0 || j == 0 || i == n || j == n {
			if i == n && j == n {
				return makeKey(j, i)
			}
			// on border
			borderPoints = append(borderPoints, []int{i, j})
		}
		if len(borderPoints) > 0 && canAchieveEnd(n, borderPoints, corruptedPoints) {
			return makeKey(j, i)
		}
	}
	return "\"not found\""
}

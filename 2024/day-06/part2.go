package day06

const (
	// cnv = 0 // cell not visited
	av  = 'X'
	cup = 1 // cell visited up dir
	crg = 2 // cell visited right dir
	cdw = 4 // cell visited down dir
	clf = 8 // cell visited left dir
)

func dirToFlag(d int) int {
	switch d {
	case 0:
		return cup
	case 1:
		return crg
	case 2:
		return cdw
	case 3:
		return clf
	}
	panic("invalid direction")
}

func visitedInTheSameDirection(v, d int) bool {
	return v&d == d
}

func writeInMap(m map[int]map[int]int, a, b, v int) {
	if _, ok := m[a]; !ok {
		m[a] = make(map[int]int)
	}
	if _, ok := m[b]; !ok {
		m[b] = make(map[int]int)
	}
	m[a][b] |= v
}

func canLoop(i, j, d int, g [][]byte, visited map[int]map[int]int) int {
	tempVisited := map[int]map[int]int{}

	for {
		ni, nj := i+dirs[d][0], j+dirs[d][1]
		if !isInside(ni, nj, g) {
			break
		}

		if g[ni][nj] == wall {
			// if I've already hit the wall in the same direction
			f := dirToFlag(d)
			if visitedInTheSameDirection(visited[ni][nj], f) || visitedInTheSameDirection(tempVisited[ni][nj], f) {
				return 1
			}
			writeInMap(tempVisited, ni, nj, f)
			d = (d + 1) % 4
			continue
		}
		i, j = ni, nj
	}
	return 0
}

func Part2(g [][]byte) int {
	visited := map[int]map[int]int{}

	gi, gj := findGuardPosition(g)
	i, j := gi, gj
	move := getDir(i, j, g)

	ans := 0
	for {
		ni, nj := i+dirs[move][0], j+dirs[move][1]
		if !isInside(ni, nj, g) {
			break
		}
		if g[ni][nj] == wall {
			writeInMap(visited, ni, nj, dirToFlag(move))
			move = (move + 1) % 4
			continue
		}

		if !(ni == gi && nj == gj) && g[ni][nj] != av {
			// simulate a wall in position (ni, nj)
			g[ni][nj] = wall
			ans += canLoop(i, j, move, g, visited)
			g[ni][nj] = av
		}
		i, j = ni, nj
	}

	return ans
}

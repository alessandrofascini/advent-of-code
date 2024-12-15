package day15

const (
	boxOpen  = '['
	boxClose = ']'
)

func resizeMap(g [][]byte) [][]byte {
	newMap := make([][]byte, len(g))
	for i, row := range g {
		newMap[i] = make([]byte, len(row)<<1)
		for j, c := range row {
			nj := j << 1
			switch c {
			case wall:
				newMap[i][nj], newMap[i][nj+1] = wall, wall
			case empty:
				newMap[i][nj], newMap[i][nj+1] = empty, empty
			case box:
				newMap[i][nj], newMap[i][nj+1] = boxOpen, boxClose
			case robot:
				newMap[i][nj], newMap[i][nj+1] = robot, empty
			}
		}
	}
	return newMap
}

func boxesToMove(i, l, r int, dir []int, g [][]byte) [][]int {
	points := make([][]int, 0)
	if r-l != 1 {
		panic("r - l != 1")
	}
	for {
		if g[i][l] == boxOpen {
			if g[i][r] != boxClose {
				panic("invalid input: " + string(g[i]))
			}
			points = append(points, []int{i, l})
		} else if g[i][l] == wall || g[i][r] == wall {
			// can't move nothing
			return nil
		} else if g[i][l] == empty || g[i][r] == empty {
			if g[i][l] == empty && g[i][r] == empty {
				return points
			}
			if g[i][l] != empty {
				if g[i][l] != boxClose {
					panic("invalid input: " + string(g[i]))
				}
				if lp := boxesToMove(i, l-1, l, dir, g); lp != nil {
					return mergeArrays(points, lp, [][]int{})
				}
				return nil
			}
			if g[i][r] != boxOpen {
				panic("invalid input: " + string(g[i]))
			}
			if rp := boxesToMove(i, r, r+1, dir, g); rp != nil {
				return mergeArrays(points, rp, [][]int{})
			}
			return nil
		} else {
			lp, rp := boxesToMove(i, l-1, l, dir, g), boxesToMove(i, r, r+1, dir, g)
			if lp == nil || rp == nil {
				return nil
			}
			return mergeArrays(lp, points, rp)
		}
		i = i + dir[0]
	}
}

func moveVertical(i, j int, dir []int, g [][]byte) (int, int) {
	l, r := j, j+1
	if g[i][l] == boxClose {
		l, r = l-1, l
	}
	boxes := boxesToMove(i, l, r, dir, g)
	if boxes == nil {
		// cannot move any box
		return -1, -1
	}
	for _, b := range boxes {
		w, l, r := b[0], b[1], b[1]+1
		g[w][l], g[w][r] = empty, empty
	}

	for _, b := range boxes {
		w, l, r := b[0]+dir[0], b[1], b[1]+1
		g[w][l], g[w][r] = boxOpen, boxClose
	}

	return i, j
}

// start from current position, return the new position
func tryToMove2(i, j int, move byte, g [][]byte) (int, int) {
	dir := moveToDir(move)
	ni, nj := i+dir[0], j+dir[1]
	symbol := g[ni][nj]
	if symbol == wall {
		return i, j
	}
	if symbol == empty {
		return ni, nj
	}
	// founded box
	if move == up || move == down {
		ni, nj = moveVertical(ni, nj, dir, g)
		if ni == -1 {
			return i, j
		}
		return ni, nj
	}
	// move is right or left
	for {
		ni, nj = ni+dir[0], nj+dir[1]
		symbol = g[ni][nj]
		if symbol == wall || symbol == empty {
			break
		}
	}
	if symbol == wall {
		return i, j
	}
	// symbol empty, can move
	for !(ni == i && nj == j) {
		pi, pj := ni-dir[0], nj-dir[1]
		g[ni][nj], g[pi][pj] = g[pi][pj], g[ni][nj]
		ni, nj = pi, pj
	}
	return i + dir[0], j + dir[1]
}

func Part2(g [][]byte, movements []byte) int {
	g = resizeMap(g)

	i, j := findRobot(g)
	g[i][j] = empty

	for _, move := range movements {
		i, j = tryToMove2(i, j, move, g)
	}

	return gps(g)
}

func mergeArrays(a, b, c [][]int) [][]int {
	n, m := len(a), len(b)
	newArray := make([][]int, n+m+len(c))
	for i, v := range a {
		newArray[i] = v
	}
	for i, v := range b {
		newArray[i+n] = v
	}
	n = n + m
	for i, v := range c {
		newArray[i+n] = v
	}
	return newArray
}

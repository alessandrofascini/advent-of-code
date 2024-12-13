package day08

func findAntiNodes(coll []*coords, p, q *coords, n, m int) []*coords {
	if q.x < p.x || q.y < p.y {
		p, q = q, p
	}
	dx, dy := q.x-p.x, q.y-p.y
	i, j := p.x, p.y
	for -1 < i && -1 < j {
		coll = append(coll, &coords{i, j})
		i, j = i-dx, j-dy
	}
	i, j = q.x, q.y
	for i < n && j < m {
		coll = append(coll, &coords{i, j})
		i, j = i+dx, j+dy
	}
	return coll
}

func getAntiNodes2(antCoords []*coords, n, m int) []*coords {
	antiNodes := make([]*coords, 0)

	for i, p := range antCoords {
		for _, q := range antCoords[i+1:] {
			antiNodes = findAntiNodes(antiNodes, p, q, n, m)
		}
	}

	return antiNodes
}

func Part2(g [][]byte) int {
	antennas := map[byte][]*coords{}

	for i, row := range g {
		for j, c := range row {
			if c != empty {
				antennas[c] = append(antennas[c], &coords{i, j})
			}
		}
	}

	ans := 0
	for _, antennaCoordinates := range antennas {
		candidates := getAntiNodes2(antennaCoordinates, len(g), len(g[0]))
		for _, candidate := range candidates {
			i, j := candidate.x, candidate.y
			if insideGrid(i, j, g) && g[i][j] != antiNode {
				g[i][j] = antiNode
				ans++
			}
		}
	}

	return ans
}

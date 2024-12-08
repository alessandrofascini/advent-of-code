package day10

import "fmt"

const (
	empty    = '.'
	antiNode = '#'
)

type coords struct {
	x, y int
}

func (c *coords) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func insideGrid(i, j int, g [][]byte) bool {
	return -1 < i && i < len(g) && -1 < j && j < len(g[0])
}

func calculateAntiNode(p, q *coords) (*coords, *coords) {
	dx, dy := q.x-p.x, q.y-p.y
	first := &coords{q.x + dx, q.y + dy}
	second := &coords{p.x - dx, p.y - dy}
	return first, second
}

func getAntiNodes(antCoords []*coords) []*coords {
	n := len(antCoords)
	antiNodes := make([]*coords, 0, (n*(n-1))>>1)

	for i, p := range antCoords {
		for _, q := range antCoords[i+1:] {
			f, s := calculateAntiNode(p, q)
			antiNodes = append(antiNodes, f, s)
		}
	}

	return antiNodes
}

func Part1(g [][]byte) int {
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
		candidates := getAntiNodes(antennaCoordinates)
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

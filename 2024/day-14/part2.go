package day14

import "fmt"

const (
	empty = ' '
	full  = '#'
)

func findEasterEgg(robots []*Robot, g [][]byte) {
	for _, robot := range robots {
		x, y := robot.pos.x, robot.pos.y
		g[x][y] = full
	}

	for _, r := range g {
		fmt.Println(string(r))
	}

	for _, robot := range robots {
		x, y := robot.pos.x, robot.pos.y
		g[x][y] = empty
	}
}

func VisualizePart2(robots []*Robot, bounds *Coordinate, until int) {
	grid := make([][]byte, bounds.x)
	for i := range grid {
		grid[i] = make([]byte, bounds.y)
		for j := range grid[i] {
			grid[i][j] = empty
		}
	}

	origin := NewCoordinate(0, 0)
	for i := 0; i < until; i++ {
		findEasterEgg(robots, grid)
		for _, robot := range robots {
			robot.updatePosition(origin, bounds)
		}
	}
}

func (c *Coordinate) Adjacent(diagonal bool) []*Coordinate {
	l := 4
	if diagonal {
		l += 4
	}
	adj := make([]*Coordinate, l)
	adj[0] = NewCoordinate(c.x-1, c.y)
	adj[1] = NewCoordinate(c.x+1, c.y)
	adj[2] = NewCoordinate(c.x, c.y-1)
	adj[3] = NewCoordinate(c.x, c.y+1)
	if diagonal {
		adj[4] = NewCoordinate(c.x-1, c.y-1)
		adj[5] = NewCoordinate(c.x-1, c.y+1)
		adj[6] = NewCoordinate(c.x+1, c.y-1)
		adj[7] = NewCoordinate(c.x+1, c.y+1)
	}
	return adj
}

func countLargestArea(robots []*Robot) int {
	coordinates := map[string]struct{}{}
	for _, robot := range robots {
		coordinates[robot.pos.String()] = struct{}{}
	}
	uf := NewUnionFind()

	for _, robot := range robots {
		adj := robot.pos.Adjacent(false)
		u := robot.pos.String()
		for _, coord := range adj {
			v := coord.String()
			if _, ok := coordinates[v]; ok {
				uf.Union(u, v)
			}
		}
	}

	return uf.SizeOfLargestSet()
}

func Part2(robots []*Robot, bounds *Coordinate) int {
	origin := NewCoordinate(0, 0)
	for i := 0; i <= 10000; i++ {
		largest := countLargestArea(robots)
		if largest > 50 {
			VisualizePart2(robots, bounds, 1)
			return i
		}
		// update position
		for _, robot := range robots {
			robot.updatePosition(origin, bounds)
		}
	}
	return 0
}

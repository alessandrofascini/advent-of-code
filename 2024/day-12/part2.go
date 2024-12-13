package day12

func newRegionWithSide(si, sj int, g [][]byte, visited map[string]struct{}) *region {
	symbol := g[si][sj]
	area, perimeter := 0, 0
	queue := [][]int{{si, sj}}

	uf := NewUnionFind()

	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]
		key := makeKey(i, j)
		if _, ok := visited[key]; !isInside(i, j, g) || ok {
			continue
		}
		visited[key] = struct{}{}

		area++
		for d, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if !isInside(ni, nj, g) || symbol != g[ni][nj] {
				perimeter++

				el := NewElement(i, j, d)
				if !uf.Belongs(el.String()) {
					el = uf.AddElement(el)
				}
				key := el.String()

				idx := (d + 1) % 4
				ni, nj = i+dirs[idx][0], j+dirs[idx][1]
				other := NewElement(ni, nj, d)
				if _, ok := uf.Find(other); ok {
					uf.Union(other.String(), key)
				}

				idx = (d + 3) % 4
				ni, nj = i+dirs[idx][0], j+dirs[idx][1]
				other = NewElement(ni, nj, d)
				if _, ok := uf.Find(other); ok {
					uf.Union(other.String(), key)
				}
			} else if symbol == g[ni][nj] {
				queue = append(queue, []int{ni, nj})
			}
		}
	}

	return &region{symbol, area, perimeter, uf.NumberOfPartitions()}
}

func Part2(g [][]byte) int {
	visited := map[string]struct{}{}
	cost := 0
	for i, row := range g {
		for j := range row {
			key := makeKey(i, j)
			if _, ok := visited[key]; !ok {
				r := newRegionWithSide(i, j, g, visited)
				//fmt.Println(r)
				cost += r.area * r.sides
			}
		}
	}
	return cost
}

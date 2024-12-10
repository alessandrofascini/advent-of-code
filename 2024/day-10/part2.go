package day10

func rate(si, sj int, m [][]byte) int {
	r := 0
	queue := [][]int{{si, sj}}
	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]
		if !inside(i, j, m) {
			continue
		}
		if m[i][j] == '9' {
			r++
			continue
		}
		sc := m[i][j] + 1
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if inside(ni, nj, m) && m[ni][nj] == sc {
				queue = append(queue, []int{ni, nj})
			}
		}
	}
	return r
}

func Part2(m [][]byte) int {
	rating := 0
	for i, r := range m {
		for j, c := range r {
			if c == '0' {
				rating += rate(i, j, m)
			}
		}
	}
	return rating
}

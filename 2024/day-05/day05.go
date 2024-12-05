package day05

func isCorrect(adj map[int]map[int]struct{}, update []int) bool {
	visited := make(map[int]struct{})
	for _, u := range update {
		visited[u] = struct{}{}
		for v := range adj[u] {
			if _, ok := visited[v]; ok {
				return false
			}
		}
	}
	return true
}

func Part1(adj map[int]map[int]struct{}, updates [][]int) int {
	ans := 0
	for _, update := range updates {
		if isCorrect(adj, update) {
			ans += update[len(update)>>1]
		}
	}
	return ans
}

func orderAns(ans []int, i int, adj map[int]map[int]struct{}) {
	candidateIndex := i
	u := ans[candidateIndex]
	nextValues := adj[u]
	for j := len(ans) - 2; j >= 0; j-- {
		v := ans[j]
		if _, ok := nextValues[v]; ok {
			candidateIndex = j
		}
	}
	if candidateIndex == i {
		return
	}
	// insert at index i
	for candidateIndex < i {
		ans[i] = ans[i-1]
		i--
	}
	ans[i] = u
}

func fixUpdate(adj map[int]map[int]struct{}, update []int) []int {
	ans := make([]int, 0, len(update))
	for i, u := range update {
		ans = append(ans, u)
		orderAns(ans, i, adj)
	}
	return ans
}

func Part2(adj map[int]map[int]struct{}, updates [][]int) int {
	ans := 0
	for _, update := range updates {
		if !isCorrect(adj, update) {
			update = fixUpdate(adj, update)
			ans += update[len(update)>>1]
		}
	}
	return ans
}

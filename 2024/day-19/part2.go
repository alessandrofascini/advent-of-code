package day19

func rec(i int, s string, root *Trie, dp []int) int {
	type item struct {
		idx int
		t   *Trie
	}
	queue := []item{{i, root}}
	n := len(s)
	ans := 0
	for len(queue) > 0 {
		i, t := queue[0].idx, queue[0].t
		queue = queue[1:]
		b := s[i]
		nt := t.children[b]
		if nt == nil {
			// no child
			continue
		}
		if nt.isEnd {
			if i+1 == n {
				ans++
				continue
			}
			if dp[i+1] > 0 {
				ans += dp[i+1]
			}
		}
		if i+1 < n {
			queue = append(queue, item{i + 1, nt})
		}
	}
	return ans
}

func waysToBuild(s string, root *Trie) int {
	dp := make([]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = rec(i, s, root, dp)
	}
	return dp[0]
}

func Part2(patterns []string, designs []string) int {
	root := NewTrie()
	for _, p := range patterns {
		root.Insert(p)
	}

	available := 0
	for _, d := range designs {
		ways := waysToBuild(d, root)
		available += ways
	}
	return available
}

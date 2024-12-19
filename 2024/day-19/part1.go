package day19

//const (
//	white = 'w'
//	black = 'b'
//	blue  = 'u'
//	red   = 'r'
//	green = 'g'
//)

func searchValue(i int, s string, root *Trie, dp []bool) bool {
	type item struct {
		idx int
		t   *Trie
	}
	queue := []item{{i, root}}
	n := len(s)
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
				return true
			}
			if dp[i+1] {
				return true
			}
		}
		if i+1 < n {
			queue = append(queue, item{i + 1, nt})
		}
	}
	return false
}

func canMake(s string, root *Trie) bool {
	dp := make([]bool, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = searchValue(i, s, root, dp)
	}
	return dp[0]
}

func Part1(patterns []string, designs []string) int {
	root := NewTrie()
	for _, p := range patterns {
		root.Insert(p)
	}

	available := 0
	for _, d := range designs {
		//fmt.Print(d)
		if canMake(d, root) {
			available += 1
			//fmt.Print(" OK")
		}
		//fmt.Println()
	}
	return available
}

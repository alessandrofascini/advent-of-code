package day19

type Trie struct {
	children map[byte]*Trie
	isEnd    bool
}

func NewTrie() *Trie {
	return &Trie{map[byte]*Trie{}, false}
}

func (t *Trie) Insert(s string) {
	curr := t
	for i := range s {
		b := s[i]
		if curr.children[b] == nil {
			curr.children[b] = NewTrie()
		}
		curr = curr.children[b]
	}
	curr.isEnd = true
}

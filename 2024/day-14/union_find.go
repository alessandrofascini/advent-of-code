package day14

type UnionFind struct {
	parents map[string]string
}

func NewUnionFind() *UnionFind {
	return &UnionFind{map[string]string{}}
}

func (uf *UnionFind) MakeSet(s string) string {
	uf.parents[s] = s
	return s
}

func (uf *UnionFind) Find(s string) string {
	for s != uf.parents[s] {
		s = uf.parents[s]
	}
	return s
}

func (uf *UnionFind) Union(u, v string) {
	if _, ok := uf.parents[u]; !ok {
		uf.MakeSet(u)
	}
	if _, ok := uf.parents[v]; !ok {
		uf.MakeSet(v)
	}

	uf.parents[v] = uf.parents[u]
}

func (uf *UnionFind) SizeOfLargestSet() int {
	p := map[string]int{}
	for key := range uf.parents {
		value := uf.Find(key)
		p[value]++
	}

	m := 0
	for _, v := range p {
		m = max(m, v)
	}
	return m
}

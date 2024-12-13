package day12

import (
	"fmt"
)

type Direction = int

type Element struct {
	i, j   int
	dir    Direction
	parent *Element
}

func NewElement(i, j int, dir Direction) *Element {
	return &Element{i, j, dir, nil}
}

func (e *Element) String() string {
	return fmt.Sprintf("%d;%d;%d", e.i, e.j, e.dir)
}

// UnionFind balanced UnionFind with path compression
type UnionFind struct {
	partitions int
	forest     map[string]*Element
}

func NewUnionFind() *UnionFind {
	return &UnionFind{0, map[string]*Element{}}
}

func (uf *UnionFind) MakeSet(i, j int, dir Direction) *Element {
	uf.partitions++
	element := NewElement(i, j, dir)
	uf.forest[element.String()] = element
	return element
}

func (uf *UnionFind) AddElement(e *Element) *Element {
	return uf.MakeSet(e.i, e.j, e.dir)
}

func (uf *UnionFind) Find(e *Element) (string, bool) {
	key := e.String()
	element, belongsTo := uf.forest[key]
	if !belongsTo {
		return "", false
	}
	if element.parent == nil {
		return key, true
	}
	// parent cannot be nil
	parent := element.parent

	elements := make([]*Element, 0)
	for parent.parent != nil {
		elements = append(elements, parent)
		element = parent
		parent = parent.parent
	}

	for _, el := range elements {
		el.parent = parent
	}

	return parent.String(), true
}

func (uf *UnionFind) alreadyRelated(a, b string) bool {
	elementA, present := uf.forest[a]
	if !present {
		return false
	}
	elementB, present := uf.forest[b]
	if !present {
		return false
	}
	// exists parent
	parentOfA, present := uf.Find(elementA)
	if !present {
		panic("parent cannot be present")
	}
	parentOfB, present := uf.Find(elementB)
	if !present {
		panic("parent cannot be present")
	}
	return parentOfA == parentOfB
}

func (uf *UnionFind) Union(a, b string) {
	// if two elements are not related, then reduce the number of partitions
	if uf.alreadyRelated(a, b) {
		return
	}
	// a and b are presents
	uf.partitions--
	parent := uf.forest[a]
	child := uf.forest[b]

	child.parent = parent
}

func (uf *UnionFind) Belongs(k string) bool {
	_, ok := uf.forest[k]
	return ok
}

func (uf *UnionFind) NumberOfPartitions() int {
	return uf.partitions
}

package day12

import (
	"testing"
)

func assert(t *testing.T, actual, expected int, s string) {
	if actual != expected {
		t.Fatalf("%s: expected %d, got %d", s, expected, actual)
	}
}

func TestUnionFind_NumberOfPartitions(t *testing.T) {
	uf := NewUnionFind()
	assert(t, uf.NumberOfPartitions(), 0, "empty union")

	const up = 0

	keys := make([]string, 0)

	el := uf.MakeSet(0, 1, up)
	keys = append(keys, el.String())

	el = uf.MakeSet(0, 0, up)
	keys = append(keys, el.String())

	el = uf.MakeSet(0, 2, up)
	keys = append(keys, el.String())

	for i := 1; i < len(keys); i++ {
		uf.Union(keys[i-1], keys[i])
	}

	assert(t, uf.NumberOfPartitions(), 1, "empty union")
}

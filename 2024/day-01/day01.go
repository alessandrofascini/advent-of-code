package day_01

import (
	"container/heap"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Part1(input [][]int) int {
	leftHeap, rightHeap := &IntHeap{}, &IntHeap{}
	for _, n := range input {
		l, r := n[0], n[1]
		heap.Push(leftHeap, l)
		heap.Push(rightHeap, r)
	}
	answer := 0
	for leftHeap.Len() > 0 {
		left := heap.Pop(leftHeap).(int)
		right := heap.Pop(rightHeap).(int)
		answer += abs(left - right)
	}
	return answer
}

func Part2(input [][]int) int {
	occ := map[int]int{}
	for _, n := range input {
		r := n[1]
		occ[r]++
	}
	score := 0
	for _, n := range input {
		l := n[0]
		score += occ[l] * l
	}
	return score
}

package day16

import (
	"container/heap"
	"fmt"
)

func Part2(lines [][]byte) int {
	si, sj := searchStartTile(lines)
	startItem := &Item{si, sj, 0, 0, 0}
	//parents := map[string]int{}

	queue := make(PriorityQueue, 0)
	heap.Push(&queue, startItem)

	distances := map[string]int{}
	distances[startItem.Key()] = 0

	items := map[string]*Item{}

	for queue.Len() > 0 {
		u := heap.Pop(&queue).(*Item)

		for nd, dir := range dirs {
			ni, nj := u.i+dir[0], u.j+dir[1]
			w := distances[u.Key()] + 1
			if u.direction&1 != nd&1 {
				w += 1000
			}
			v := &Item{ni, nj, nd, w, -1}
			key := v.Key()
			if dsv, ok := distances[key]; ok && w < dsv {
				v = items[key]
				queue.update(v, nd, w)
			} else {
				heap.Push(&queue, v)
				distances[key] = w
				items[key] = v
			}
		}

	}

	ei, ej := searchEndTile(lines)
	fmt.Println(distances[(&Item{ei, ej, 0, 0, 0}).Key()])

	return 0
}

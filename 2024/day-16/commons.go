package day16

import "fmt"

func searchSymbol(g [][]byte, symbol byte) (int, int) {
	for i, r := range g {
		for j, c := range r {
			if c == symbol {
				return i, j
			}
		}
	}
	return -1, -1
}

func searchStartTile(g [][]byte) (int, int) {
	return searchSymbol(g, startTile)
}

func searchEndTile(g [][]byte) (int, int) {
	return searchSymbol(g, endTile)
}

func makeKey(i, j int) string {
	return fmt.Sprintf("%d;%d", i, j)
}

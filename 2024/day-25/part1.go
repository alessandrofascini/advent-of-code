package day25

import (
	"strconv"
	"strings"
)

const (
	filled = '#'
)

type Schematic = [][]byte

func makeKey(scheme Schematic) []int {
	n := len(scheme[0])
	key := make([]int, n)
	for _, r := range scheme[1 : len(scheme)-1] {
		for j, c := range r {
			if c == filled {
				key[j]++
			}
		}
	}
	return key
}

func reverseKey(key []int, m int) []int {
	ans := make([]int, len(key))
	for i, k := range key {
		ans[i] = m - k
	}
	return ans
}

func keyToString(k []int) string {
	strs := make([]string, len(k))
	for i, c := range k {
		strs[i] = strconv.Itoa(c)
	}
	return strings.Join(strs, ",")
}

func Part1(schematics []Schematic) int {
	m := len(schematics[0]) - 2
	ans := 0
	visited := map[string]struct{}{}
	for _, scheme := range schematics {
		key := makeKey(scheme)
		keyStr := keyToString(key)
		if _, ok := visited[keyStr]; ok {
			ans++
		}
		reversed := reverseKey(key, m)
		revKeyStr := keyToString(reversed)
		visited[revKeyStr] = struct{}{}
	}
	return ans
}

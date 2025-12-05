package day05

import "sort"

type IngredientId = int

type Interval [2]int

func NewInterval(start, end int) Interval {
	if start > end {
		return Interval{end, start}
	}
	return Interval{start, end}
}

func Merge(a, b Interval) Interval {
	return NewInterval(min(a[0], b[0]), max(a[1], b[1]))
}

func (i Interval) Compare(other Interval) int {
	a, b := i[0], i[1]
	c, d := other[0], other[1]
	switch {
	case d < a:
		return -1
	case b < c:
		return 1
	default:
		return 0
	}
}

func (i Interval) IsInside(v int) bool {
	return i[0] <= v && v <= i[1]
}

func belongs(ranges []Interval, id IngredientId) bool {
	for _, interval := range ranges {
		if interval.IsInside(id) {
			return true
		}
	}
	return false
}

func Part1(ranges []Interval, ingredients []IngredientId) int {
	ans := 0
	for _, ingredient := range ingredients {
		if belongs(ranges, ingredient) {
			ans++
		}
	}
	return ans
}

func Part2(ranges []Interval) int {
	// merge ranges
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Compare(ranges[j]) > 0
	})

	newRanges := make([]Interval, 0, len(ranges))
	for _, interval := range ranges {
		i := len(newRanges) - 1
		if i == -1 {
			newRanges = append(newRanges, interval)
			continue
		}
		curr := interval
		for i > -1 && newRanges[i].Compare(curr) == 0 {
			curr = Merge(newRanges[i], curr)
			i--
		}
		newRanges = append(newRanges[:i+1], curr)
	}

	ans := 0
	for _, interval := range newRanges {
		ans += interval[1] - interval[0] + 1
	}
	return ans
}

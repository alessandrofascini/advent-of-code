package day02

import (
	"strconv"
)

type Range struct {
	Lower int64
	Upper int64
}

func SplitEvery(s string, m int) []string {
    if m <= 0 {
        return []string{s}
    }

    var chunks []string
    for i := 0; i < len(s); i += m {
        end := i + m
        if end > len(s) {
            end = len(s)
        }
        chunks = append(chunks, s[i:end])
    }
    return chunks
}

func isInvalid(n int64) bool {
	s := strconv.FormatInt(n, 10)
	l := len(s) >> 1
	return s[l:] == s[:l]
}

func invalidIds(r Range, isInvalid func (int64) bool) int64 {
	ans := int64(0)
	for i := r.Lower; i <= r.Upper; i++ {
		if(isInvalid(i)) {
			ans += i
		}
	}
	return ans
}

func Part1(ranges []Range) int64 {
	ans := int64(0)
	for _, r := range ranges {
		ans += invalidIds(r, isInvalid)
	}
	return ans
}

func isInvalid2(n int64) bool {
	s := strconv.FormatInt(n, 10)
	l := len(s) >> 1 + 1

	check := func (s []string) bool {
		if(len(s) < 2) {
			return false
		}
		t := s[0]
		for _, v := range s {
			if v != t {
				return false
			}
		}
		return true
	}

	for i := l; i > 0; i-- {
		if(len(s) % i != 0) {
			continue
		}
		splitted := SplitEvery(s, i)
		if(check(splitted)) {
			return true
		}
	}
	return false
}

func Part2(ranges []Range) int64 {
	ans := int64(0)
	for _, r := range ranges {
		ans += invalidIds(r, isInvalid2)
	}
	return ans
}
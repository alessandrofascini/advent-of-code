package day_03

import (
	"strconv"
	"strings"
)

func mul(s string, i int) int {
	// at least 8 characters length
	if len(s)-i < 8 {
		return 0
	}
	if s[i:i+4] != "mul(" {
		return 0
	}
	i += 4
	commaIdx := strings.Index(s[i:], ",")
	if commaIdx == -1 {
		return 0
	}
	firstNumber, err := strconv.Atoi(s[i : i+commaIdx])
	if err != nil {
		return 0
	}
	i += commaIdx + 1
	parIdx := strings.Index(s[i:], ")")
	if parIdx == -1 {
		return 0
	}
	secondNumber, err := strconv.Atoi(s[i : i+parIdx])
	if err != nil {
		return 0
	}
	return firstNumber * secondNumber
}

func Part1(s string) int {
	ans := 0
	for i := range s {
		if s[i] == 'm' {
			ans += mul(s, i)
		}
	}
	return ans
}

func checkEnable(s string, i int) (bool, bool) {
	if len(s)-i < 4 {
		return false, false
	}
	if s[i:i+4] == "do()" {
		return true, true
	}
	if s[i:i+7] == "don't()" {
		return false, true
	}
	return false, false
}

func Part2(s string) int {
	ans := 0
	enable := true
	for i := range s {
		if enable && s[i] == 'm' {
			ans += mul(s, i)
		} else if s[i] == 'd' {
			if t, ok := checkEnable(s, i); ok {
				enable = t
			}
		}
	}
	return ans
}

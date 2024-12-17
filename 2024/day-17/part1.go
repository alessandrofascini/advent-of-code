package day17

import (
	"strconv"
	"strings"
)

const (
	adv = 0
	bxl = 1
	bst = 2
	jnz = 3
	bxc = 4
	out = 5
	bdv = 6
	cdv = 7
)

func comboOperator(A, B, C, op int) int {
	switch op {
	case 0, 1, 2, 3:
		return op
	case 4:
		return A
	case 5:
		return B
	case 6:
		return C
	case 7:
		fallthrough
	default:
		panic("invalid operation")
	}
}

func exeAdv(A, B, C, op int) int {
	num := A
	den := comboOperator(A, B, C, op)
	den = 1 << den
	return num / den
}

func exeBxl(_, B, _ int, op int) int {
	// XOR
	return B ^ op
}

func exeBst(A, B, C, op int) int {
	return comboOperator(A, B, C, op) % 8
}

func exeJnz(A, _, _, op int) int {
	if A == 0 {
		return -1
	}
	return op - 2
}

func exeBxc(_, B, C, _ int) int {
	return B ^ C
}

func exeOut(A, B, C, op int) int {
	return comboOperator(A, B, C, op) % 8
}

func Part1(A, B, C int, program []int) string {
	P := len(program)
	outputs := make([]string, 0)
	for ip := 0; ip < P; ip = ip + 2 {
		cmd, op := program[ip], program[ip+1]
		switch cmd {
		case adv:
			A = exeAdv(A, B, C, op)
		case bxl:
			B = exeBxl(A, B, C, op)
		case bst:
			B = exeBst(A, B, C, op)
		case jnz:
			if j := exeJnz(A, B, C, op); j != -1 {
				ip = j
			}
		case bxc:
			B = exeBxc(A, B, C, op)
		case out:
			result := exeOut(A, B, C, op)
			// out this value
			s := strconv.Itoa(result)
			str := strings.Split(s, "")
			for i := len(str) - 1; i >= 0; i-- {
				outputs = append(outputs, str[i])
			}
		case bdv:
			B = exeAdv(A, B, C, op)
		case cdv:
			C = exeAdv(A, B, C, op)
		default:
			panic("invalid command")
		}
	}
	return strings.Join(outputs, ",")
}

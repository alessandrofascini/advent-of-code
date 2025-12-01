package day01

type Row struct {
	Dir byte
	Val int
}

const Left = 'L'

const Right = 'R'

const Module = 100

func Part1(input []Row) int {
	ans, curr := 0, 50
	for _, row := range input {
		if row.Dir == Left {
			curr = (curr - row.Val + Module) % Module
		} else if row.Dir == Right {
			curr = (curr + row.Val) % Module
		} else {
			panic("invalid dir")
		}
		if curr == 0 {
			ans++
		}
	}
	return ans
}

func Part2(input []Row) int {
	ans, curr := 0, 50
	for _, row := range input {
		ans += row.Val / Module
		val := row.Val % Module
		if row.Dir == Left {
			if val > curr {
				ans++
			}
			curr = (curr - val + Module) % Module
		} else if row.Dir == Right {
			if curr+val >= Module {
				ans++
			}
			curr = (curr + val) % Module
		} else {
			panic("invalid dir")
		}
	}
	return ans
}

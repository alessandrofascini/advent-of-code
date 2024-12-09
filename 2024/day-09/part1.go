package day09

const (
	free = -1
)

func transformBlock(input []byte) []int {
	length := 0
	for _, b := range input {
		length += int(b - '0')
	}
	output := make([]int, 0, length)
	freeSpace := false
	id := 0
	for _, b := range input {
		charToPrint := free
		if !freeSpace {
			charToPrint = id
			id++
		}
		n := int(b - '0')
		for i := 0; i < n; i++ {
			output = append(output, charToPrint)
		}
		freeSpace = !freeSpace
	}
	return output
}

func reduceSpace(input []int) {
	left, right := 0, len(input)-1
	for left < right {
		if input[left] != free {
			left++
			continue
		}
		if input[right] == free {
			right--
			continue
		}
		input[left], input[right] = input[right], input[left]
		left++
		right--
	}
}

func checksum(disk []int) int {
	sum := 0
	for i, v := range disk {
		if v == free {
			continue
		}
		sum += i * v
	}
	return sum
}

func Part1(input []byte) int {
	disk := transformBlock(input)
	reduceSpace(disk)
	return checksum(disk)
}

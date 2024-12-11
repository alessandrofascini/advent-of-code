package day09

//type file struct {
//	id, from, to int
//}
//
//func (f *file) String() string {
//	return fmt.Sprintf("file{id: %d, from: %d, to: %d}", f.id, f.from, f.to)
//}
//
//func (f *file) length() int {
//	return f.to - f.from
//}
//
//func createDisk(input []byte) ([]*file, map[int][]*file) {
//	disk := make([]*file, 0, len(input))
//	fs := map[int][]*file{}
//	id := 0
//	freeSpace := false
//	j := 0
//	for _, b := range input {
//		n := int(b - '0')
//		to := j + n
//		if freeSpace {
//			t := &file{-1, j, to}
//			key := t.length()
//			fs[key] = append(fs[key], t)
//		} else {
//			disk = append(disk, &file{id, j, to})
//			id++
//		}
//		freeSpace = !freeSpace
//		j = to
//	}
//	return disk, fs
//}
//
//func checksum2(disk []*file) int {
//	sum := 0
//	for _, f := range disk {
//		v := f.to * (f.to - 1)
//		v -= f.from * (f.from - 1)
//		v = v >> 1
//		sum += f.id * v
//	}
//	return sum
//}
//
//func compact(disk []*file, fs map[int][]*file) []*file {
//	maxSpaceAvailable := 0
//	for key := range fs {
//		maxSpaceAvailable = max(maxSpaceAvailable, key)
//	}
//
//	for i := len(disk) - 1; i > -1; i-- {
//		// leftmost position
//		f := disk[i]
//		spaceRequired := f.length()
//		for spaceRequired <= maxSpaceAvailable {
//			if freeSpaces, ok := fs[spaceRequired]; ok {
//				freeSpace := freeSpaces[0]
//				if !(freeSpace.from < f.from) {
//					spaceRequired++
//					continue
//				}
//				/// pop from map
//				fs[spaceRequired] = fs[spaceRequired][1:]
//				f.from, f.to = freeSpace.from, freeSpace.from+f.length()
//				freeSpace.from += f.length()
//
//				if key := freeSpace.length(); key > 0 {
//					fs[key] = append(fs[key], freeSpace)
//					// keep sorted
//					arr := fs[key]
//					sort.Slice(arr, func(i, j int) bool {
//						return arr[i].from < arr[j].from
//					})
//					fs[key] = arr
//				}
//				///
//				if len(fs[spaceRequired]) == 0 {
//					delete(fs, spaceRequired)
//				}
//				break
//			}
//			spaceRequired++
//		}
//	}
//
//	return disk
//}
//
//func Part2(input []byte) int {
//	disk, fs := createDisk(input)
//	disk = compact(disk, fs)
//	return checksum2(disk)
//}

func compact(disk []int) {
	right := len(disk) - 1
	for right > -1 {
		// search the index of last element
		if disk[right] == free {
			right--
			continue
		}
		// disk[left] = free
		// disk[right] != free
		fsidx := right
		for fsidx > -1 && disk[fsidx] == disk[right] {
			fsidx--
		}
		fsidx++
		right++

		n := right - fsidx
		// search for empty space with n size
		lsid := firstEmptyStateWithNSize(disk, n)
		if lsid != -1 && lsid < fsidx {
			for i := 0; i < n; i++ {
				disk[lsid+i], disk[fsidx+i] = disk[fsidx+i], disk[lsid+i]
			}
		}
		right = fsidx - 1
	}
}

func firstEmptyStateWithNSize(disk []int, n int) int {
	counter := 0
	idx := -1
	for i, v := range disk {
		if v == free {
			counter++
			if idx == -1 {
				idx = i
			}
		} else {
			counter = 0
			idx = -1
		}
		if counter == n {
			return idx
		}
	}
	return idx
}

func Part2(input []byte) int {
	disk := transformBlock(input)
	compact(disk)
	return checksum(disk)
}
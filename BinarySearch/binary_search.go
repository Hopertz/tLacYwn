package binarysearch

func binary_search(listnums []int, num int) bool {

	left := 0
	right := len(listnums) - 1

	for left <= right {
		idx := (left + right) / 2
		if num > listnums[idx] {
			left = idx + 1
		} else if num < listnums[idx] {
			right = idx - 1
		} else {
			return true
		}
	}

	return false
}

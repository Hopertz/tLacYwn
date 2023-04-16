package linearsearch

func linear_search(listnums []int, num int) bool {
	for _, no := range listnums {
		if no == num {
			return true
		}
	}

	return false
}

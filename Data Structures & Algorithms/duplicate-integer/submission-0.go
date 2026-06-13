func hasDuplicate(nums []int) bool {
	has := make(map[int]struct{})

	for _, value := range nums {
		if _, ok := has[value]; ok {
			return true
		}

		has[value] = struct{}{}
	}
	
	return false
}

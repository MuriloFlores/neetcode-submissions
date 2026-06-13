func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		compl := target - num

		if indexCompl, ok := seen[compl]; ok {
			return []int{indexCompl, i}
		}

		seen[num] = i
	}

	return nil
}
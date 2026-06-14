func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	left := sortArray(nums[:mid])
	rigth := sortArray(nums[mid:])

	return merge(left, rigth)
}

func merge(left, rigth []int) []int {
	result := make([]int, 0, len(left) + len(rigth))

	i, j := 0, 0

	for i < len(left) && j < len(rigth) {
		if left[i] <= rigth[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, rigth[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, rigth[j:]...)

	return result
}

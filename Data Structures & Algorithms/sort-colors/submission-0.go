func sortColors(nums []int) {
    low, mid, high := 0, 0, len(nums) - 1

    for mid <= high {
        if nums[mid] == 0 {
            nums[low], nums[mid] = nums[mid], nums[low]
            low += 1
            mid += 1
        } else if nums[mid] == 1 {
            mid += 1
        } else {
            nums[mid], nums[high] = nums[high], nums[mid]
            high -= 1
        }
    }
}

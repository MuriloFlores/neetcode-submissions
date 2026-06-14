func majorityElement(nums []int) int {
    counter := make(map[int]int)

    for _, value := range nums {
        counter[value]++

        if counter[value] > len(nums) / 2 {
            return value
        }
    }

    return 0
}

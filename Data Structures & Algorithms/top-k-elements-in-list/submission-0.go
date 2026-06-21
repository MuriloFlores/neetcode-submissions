func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }

    buckets := make([][]int, len(nums) + 1)

    for num, count := range freq {
        buckets[count] = append(buckets[count], num)
    }

    result := make([]int, 0, k)

    for i := len(buckets) - 1; i >= 0; i-- {
        if len(buckets[i]) > 0 {
            result = append(result, buckets[i]...)

            if len(result) >= k {
                break
            }
        }
    }

    return result
}

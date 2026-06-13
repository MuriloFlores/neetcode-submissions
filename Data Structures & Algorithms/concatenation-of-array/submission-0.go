func getConcatenation(nums []int) []int {
   for _, value := range nums {
    nums = append(nums, value)
   }
   return nums
}

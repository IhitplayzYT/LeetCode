func maxAlternatingSum(nums []int) int64 {
    // Square and sort then interview largest and smallest alternating 
	n := len(nums)
	for i := 0; i < n; i++ { nums[i] = nums[i] * nums[i]}
	sort.Ints(nums)
	mid := n / 2
	var ret int64
	for i := n - 1; i >= 0; i-- {
		if i >= mid {
			ret += int64(nums[i])
		} else {
			ret -= int64(nums[i])
		}
	}
	return ret
}
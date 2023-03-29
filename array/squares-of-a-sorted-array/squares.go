package squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	size := len(nums)
	result := make([]int, size, size)
	k, l := size-1, 0
	r := size - 1
	for l <= r {
		lm, rm := nums[l]*nums[l], nums[r]*nums[r]
		if lm > rm {
			result[k] = lm
			l++
		} else {
			result[k] = rm
			r--
		}
		k--
	}
	return result
}

package minimum_size_subarray_sum

func minSubArrayLen(target int, nums []int) int {
	i, sum := 0, 0
	size := len(nums)
	maxLen := size + 1 // 初始化返回长度为 size+1，目的是为了判断 “不存在符合条件的子数组，返回0” 的情况
	result := maxLen
	for j, v := range nums {
		sum += v
		for sum >= target {
			subL := j - i + 1
			if subL < result {
				result = subL
			}
			sum = sum - nums[i] // 减去首位元素，继续判断。
			i++
		}
	}
	if result == maxLen {
		return 0
	}
	return result
}

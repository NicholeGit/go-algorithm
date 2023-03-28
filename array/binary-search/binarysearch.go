package binary_search

//使用 [l,r] 不变式
func search1(nums []int, target int) int {
	//[l,r]
	l, r := 0, len(nums)-1
	for l <= r { // [1:1]
		m := (l + r) >> 1
		switch {
		case nums[m] > target:
			r = m - 1
		case nums[m] < target:
			l = m + 1
		case nums[m] == target:
			return m
		}
	}
	return -1
}

//	使用 [l,r) 不变式
func search2(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r { // [1:2)
		m := (l + r) >> 1
		switch {
		case nums[m] > target:
			r = m
		case nums[m] < target:
			l = m + 1
		case nums[m] == target:
			return m
		}
	}
	return -1
}

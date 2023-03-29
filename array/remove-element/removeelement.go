package remove_element

func removeElement(nums []int, val int) int {
	slow := 0
	for _, v := range nums {
		if v != val {
			nums[slow] = v
			slow++
		}
	}
	return slow
}

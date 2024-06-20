package remove_element

func removeElement(nums []int, val int) int {
	count := 0
	for _, v := range nums {
		if v != val {
			nums[count] = v
			count++
		}
	}
	return count
}

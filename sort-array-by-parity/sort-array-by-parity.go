package sort_array_by_parity

func sortArrayByParity(nums []int) []int {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

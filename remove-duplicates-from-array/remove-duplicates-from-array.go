package remove_duplicates_from_array

func removeDuplicates(nums []int) int {
	unique := 0
	for index := range nums {
		if nums[index] != nums[unique] {
			nums[index], nums[unique+1] = nums[unique+1], nums[index]
			unique++
		}
	}
	return unique + 1
}

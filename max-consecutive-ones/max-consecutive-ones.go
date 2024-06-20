package max_consecutive_ones

func findMaxConsecutiveOnes(nums []int) int {
	var max int
	var count int
	for index := range nums {
		if nums[index] == 1 {
			count++
		} else {
			count = 0
		}
		if count >= max {
			max = count
		}
	}
	return max
}

package running_sum_of_1d_array

func runningSum(nums []int) []int {
	output := make([]int, len(nums))
	var sum int
	for i := range nums {
		sum += nums[i]
		output[i] += sum
	}
	return output
}

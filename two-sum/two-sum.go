package two_sum

func twoSum(nums []int, target int) []int {
	var answer []int

	for i := range nums {
		for j := range nums {
			if nums[i]+nums[j] == target && i != j {
				answer = []int{i, j}
				return answer
			}
		}
	}
	return nil
}

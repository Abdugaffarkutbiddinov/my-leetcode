package find_numbers_with_even_numbers_of_digits

func findNumbers(nums []int) int {
	even := 0
	for _, value := range nums {
		digits := 0
		for value != 0 {
			value /= 10
			digits++
		}
		if digits%2 == 0 {
			even++
		}
	}
	return even
}

package kids_with_greatest_number_of_candies

func max(candies []int) (result int) {

	for i := range candies {
		if candies[i] >= result {
			result = candies[i]
		}
	}
	return result
}
func kidsWithCandies(candies []int, extraCandies int) (result []bool) {
	var max = max(candies)
	for i := range candies {
		if candies[i]+extraCandies >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}

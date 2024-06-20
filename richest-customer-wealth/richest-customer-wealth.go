package richest_customer_wealth

func maximumWealth(accounts [][]int) int {
	var max int
	for _, customer := range accounts {
		var sum int
		for _, amount := range customer {
			sum += amount
		}
		if sum >= max {
			max = sum
		}
	}
	return max
}

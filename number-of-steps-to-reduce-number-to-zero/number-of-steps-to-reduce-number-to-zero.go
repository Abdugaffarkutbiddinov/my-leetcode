package number_of_steps_to_reduce_number_to_zero

func numberOfSteps(num int) int {
	var step int
	for num != 0 {
		step++
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
	}
	return step
}

package check_if_n_and_its_double_exist

func checkIfExist(arr []int) bool {
	m := map[int]bool{}

	for _, n := range arr {
		if m[n*2] || (n%2 == 0 && m[n/2]) {
			return true
		}
		m[n] = true
	}

	return false
}

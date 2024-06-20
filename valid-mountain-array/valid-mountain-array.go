package validmountainarray

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	i, max := 0, len(arr)-1

	for ; i < max && arr[i] < arr[i+1]; i++ {
	}

	if i == 0 || i == max {
		return false
	}

	for ; i < max && arr[i] > arr[i+1]; i++ {
	}

	return i == max
}

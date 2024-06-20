package duplicate_zeros

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func duplicateZeros(arr []int) {
	var queue []int
	for i := range arr {
		if arr[i] == 0 {
			queue = append(queue, 0, 0)
			arr[i] = queue[0]
			queue = RemoveIndex(queue, 0)
		} else {
			queue = append(queue, arr[i])
			arr[i] = queue[0]
			queue = RemoveIndex(queue, 0)
		}
	}

}

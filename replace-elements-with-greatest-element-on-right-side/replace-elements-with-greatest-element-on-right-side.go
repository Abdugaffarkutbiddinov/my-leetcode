package replace_elements_with_greatest_element_on_right_side

func replaceElements(arr []int) []int {
	maxSoFar := -1
	for i := len(arr) - 1; i >= 0; i-- {
		currElement := arr[i]
		arr[i] = maxSoFar
		if currElement > maxSoFar {
			maxSoFar = currElement
		}
	}
	return arr
}

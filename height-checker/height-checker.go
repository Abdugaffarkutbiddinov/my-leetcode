package height_checker

import "sort"

func heightChecker(heights []int) int {
	var count int
	var expected []int
	for _, height := range heights {
		expected = append(expected, height)
	}
	sort.Ints(expected)

	for i := range expected {
		if expected[i] != heights[i] {
			count++
		}
	}

	return count
}

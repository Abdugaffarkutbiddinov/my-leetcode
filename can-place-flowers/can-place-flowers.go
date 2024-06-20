package can_place_flowers

// Time complexity:  O(n)
// Space complexity: O(1)
func canPlaceFlowers(flowerbed []int, n int) bool {
	// plant a flower in the center plot if the right and
	// left plots are empty
	left, center, right := true, true, false
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		left, center, right = center, right, flowerbed[i] == 1
		if !left && !center && !right {
			center, n = true, n-1
		}
	}
	if !center && !right {
		n--
	}
	return n <= 0
}

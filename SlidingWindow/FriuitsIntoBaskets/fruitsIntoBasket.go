package fruitsintobasket

// Input: {"A", "B", "C", "A", "C"}
// Output: 3

func fruitsIntoBasket(fruits []string) int {
	endWindow := 0
	startWindow := 0
	maxLength := 0
	charFrequeny := make(map[string]int)

	for endWindow = 0; endWindow < len(fruits); endWindow++ {
		rightChar := fruits[endWindow]
		charFrequeny[rightChar] += 1

		for len(charFrequeny) > 2 {
			leftChar := fruits[startWindow]

			charFrequeny[leftChar] -= 1

			if charFrequeny[leftChar] == 0 {
				delete(charFrequeny, leftChar)
			}

			startWindow += 1
		}
		maxLength = max(maxLength, endWindow-startWindow+1)
	}
	return maxLength
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

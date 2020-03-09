package largestSubstringKDistinct

// Input  -> String :{"A", "R", "A", "A", "C", "I"}, K: 2
// Output -> 3
// String :{"A", "A", "R", "A", "A", "C", "I"}
func largestSubstringKDistinct(str []string, k int) int {
	endWindow := 0
	startWindow := 0
	maxLength := 0
	charFrequeny := make(map[string]int)

	for endWindow = 0; endWindow < len(str); endWindow++ {
		rightChar := str[endWindow]
		charFrequeny[rightChar] += 1

		for len(charFrequeny) > k {
			leftChar := str[startWindow]

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

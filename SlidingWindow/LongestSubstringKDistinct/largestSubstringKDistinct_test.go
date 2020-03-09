package largestSubstringKDistinct

import (
	"testing"
)

/*
  Given a string, find the length of the longest substring in it with no
  more than K distinct characters.
*/

func TestLargestSubstringKDistinct(t *testing.T) {

	input := []string{"A", "R", "A", "A", "C", "I"}
	result := largestSubstringKDistinct(input, 2)

	if result != 4 {
		t.Errorf("Expected 4, but it was %v instead.", false)
	}
}

func TestLargestSubstringKDistinct2(t *testing.T) {

	input := []string{"A", "R", "A", "A", "C", "I"}
	result := largestSubstringKDistinct(input, 1)

	if result != 2 {
		t.Errorf("Expected 2, but it was %v instead.", false)
	}
}

func TestLargestSubstringKDistinct3(t *testing.T) {

	input := []string{"C", "B", "B", "E", "B", "I"}
	result := largestSubstringKDistinct(input, 3)

	if result != 5 {
		t.Errorf("Expected 5, but it was %v instead.", false)
	}
}

package findKLargestNumbers

import (
	"testing"
)

func TestFindKLargestNumbers(t *testing.T) {

	input := []int{3, 1, 5, 12, 2, 11}
	result := findKLargestNumbers(input, 3)

	if result != nil {
		t.Errorf("Expected 4, but it was %v instead.", false)
	}
}

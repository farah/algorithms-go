package wordSearch

import (
	"testing"
)

func TestWordSearch(t *testing.T) {

	matrix := [][]string{
		{"A", "C", "A", "T"},
		{"A", "A", "T", "K"},
		{"A", "M", "Z", "N"},
		{"Z", "A", "O", "H"},
		{"Q", "Z", "O", "N"},
	}

	words := []string{"amazon", "cat"}
	result := wordSearch(matrix, words)

	if result != 4 {
		t.Errorf("Expected 4, but it was %v instead.", false)
	}
}

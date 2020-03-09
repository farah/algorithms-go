package permutationInAString

import (
	"testing"
)

/*
  Given an array of characters where each character represents a fruit tree,
  you are given two baskets and your goal is to put maximum number of fruits in each basket.
  The only restriction is that each basket can have only one type of fruit.

  You can start with any tree, but once you have started you can’t skip a tree.
  You will pick one fruit from each tree until you cannot, i.e., you will stop
  when you have to pick from a third fruit type.

  Write a function to return the maximum number of fruits in both the baskets.
*/

func TestPermutationInAString(t *testing.T) {

	input := []string{"A", "B", "C", "A", "C"}
	result := permutationInAString(input)

	if result != 3 {
		t.Errorf("Expected 3, but it was %v instead.", false)
	}
}
package mergeintervals

import (
	"testing"
)

func TestMergeIntervals(t *testing.T) {

	intervals := []Interval{
		{Start: 1, End: 4},
		{Start: 2, End: 5},
		{Start: 2, End: 4},
	}
	result := mergeIntervals(intervals)

	if result != nil {
		t.Errorf("Expected 4, but it was %v instead.", false)
	}
}

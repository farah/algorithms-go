package mergeintervals

type Interval struct {
	Start int
	End   int
}

/*
	intervals := []Interval{
		{Start: 1, End: 4},
		{Start: 2, End: 5},
		{Start: 7, End: 9},
  }
*/
func mergeIntervals(intervals []Interval) []Interval {
	var mergedIntervals []Interval

	start := intervals[0].Start
	end := intervals[0].End

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval.Start <= end {
			end = max(end, interval.End)
		} else {
			mergedIntervals = append(mergedIntervals, Interval{start, end})
			start = interval.Start
			end = interval.End

		}
	}

	mergedIntervals = append(mergedIntervals, Interval{start, end})

	return mergedIntervals
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func merge(intervals [][]int) (ret [][]int) {
	if len(intervals) <= 1 {
		return intervals
	}

	// make sure intervals are sorted by 'start' ascending
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		a := intervals[i-1]
		b := intervals[i]

		// overlaps
		if a[1] >= b[0] && a[0] <= b[1] {
			b[0] = min(a[0], b[0])
			b[1] = max(a[1], b[1])
		} else {
			ret = append(ret, a)
		}
	}

	// should take care of the last element
	ret = append(ret, intervals[len(intervals)-1])

	return ret
}

// Leetcode #56: Merge Intervals
// https://leetcode.com/problems/merge-intervals/submissions/
func main() {
	var intervals = [][]int{
		// {1, 3}, {2, 6}, {8, 10}, {15, 18},
		{1, 4}, {1, 4},
	}

	fmt.Printf("%v", merge(intervals))
}

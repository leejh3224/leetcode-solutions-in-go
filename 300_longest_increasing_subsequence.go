package main

import (
	"fmt"
	"sort"
)

/*
 * solution using Dynamic Programming
 */
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))

	// initialize
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		// lenght of previous longest sequence
		len := 0
		for j := 0; j < i; j++ {
			// if we can extend longest sequence
			if dp[j] > len && nums[j] < nums[i] {
				len = dp[j]
			}
		}
		// add self
		dp[i] = len + 1
	}

	// to get max
	sort.Ints(dp)

	return dp[len(nums)-1]
}

func main() {
	// var input = []int{
	// 	10, 9, 2, 5, 3, 7, 101, 18,
	// }
	var input = []int{
		-2, -1,
	}

	fmt.Printf("%v", lengthOfLIS(input))
}

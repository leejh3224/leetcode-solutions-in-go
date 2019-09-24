package main

import (
	"fmt"
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

/*
 * this can be solved using dynamic programming
 * it can be tricky cause you have to consider negative values
 */
func maxProduct(nums []int) (ret int) {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	/*
	 * dp slice holds both min and max values.
	 * The reason why dp should keep track of min values is because min value can be negative
	 * and if negative value meets negative value, it may turn to max value.
	 */
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]

	for i := 1; i < len(nums); i++ {
		cur := nums[i]

		// should compare min value as well
		dp[i][0] = max(cur, max(cur*dp[i-1][0], cur*dp[i-1][1]))
		dp[i][1] = min(cur, min(cur*dp[i-1][0], cur*dp[i-1][1]))
	}

	ret = dp[0][0]

	// take max from "max product slice"
	for _, v := range dp {
		if v[0] > ret {
			ret = v[0]
		}
	}

	return ret
}

// Leetcode #152: Maximum Product Subarray
// https://leetcode.com/problems/maximum-product-subarray/submissions/
func main() {
	var input = []int{2, 3, -2, 4}
	var input2 = []int{-2, 3, -4}

	fmt.Println(maxProduct(input))
	fmt.Println(maxProduct(input2))
}

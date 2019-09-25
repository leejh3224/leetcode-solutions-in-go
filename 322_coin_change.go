package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
 * Dynamic Programming problem
 * solution in Java: https://www.programcreek.com/2015/04/leetcode-coin-change-java/
 */
func coinChange(coins []int, amount int) (ret int) {
	if amount == 0 {
		return -1
	}

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			// we can get here by 1 step
			if i == coin {
				dp[i] = 1
			}

			if i > coin {
				// unreachable
				if dp[i-coin] == math.MaxInt64 {
					continue
				}
				// take min of prev step + 1 and value at i
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}

	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

// Leetcode #322: Coin Change
// https://leetcode.com/problems/coin-change/
func main() {
	var (
		coins  = []int{1, 2, 5}
		amount = 11
	)

	fmt.Printf("%v", coinChange(coins, amount))
}

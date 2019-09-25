package main

import (
	"fmt"
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
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}

	if dp[amount] == amount+1 {
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

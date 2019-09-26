package main

import "fmt"

/*
 * Dynamic Programming
 * If obstacle is at grid[i][j], there is no way you can get there so it's 0.
 * And as we can move to right or down, so if a tile has a neighboring tile in left or up,
 * Then we plus 1 to dp[i][j].
 */
func uniquePathsWithObstacles(grid [][]int) int {
	lenX := len(grid[0])
	lenY := len(grid)

	dp := make([][]int, lenX)
	for i := range dp {
		dp[i] = make([]int, lenY)
	}
	dp[0][0] = 1

	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			if grid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				if i > 0 {
					dp[i][j] += dp[i-1][j]
				}
				if j > 0 {
					dp[i][j] += dp[i][j-1]
				}
			}
		}
	}

	return dp[lenX-1][lenY-1]
}

// Leetcode #63: Unique Paths II
// https://leetcode.com/problems/unique-paths-ii/
func main() {
	var grid = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	fmt.Printf("%v", uniquePathsWithObstacles(grid))
}

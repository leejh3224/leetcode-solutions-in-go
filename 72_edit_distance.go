package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
 * You can use dynamic programming.
 *
 * Let's see how this can be solved via DP.
 * Assume two strings are "horse" and "ros".
 * It becomes clear if you compare edit distance of two strings in table format.
 *
 * "" -> empty string
 *
 *    "" r o s
 * "" 0  1 2 3
 * h  1  1 2 3
 * o  2  2 1 2
 * r  3  2 2 2
 * s  4  3 3 2
 * e  5  4 4 3
 *
 * Do you see the pattern?
 * if char at i and j are same, then it means it doesn't cost more to add 1 more character.
 * So the distance[i][j] equals to distance[i-1][j-1] (ex (2,2) and (3,3)).
 *
 * What if char at i and j are different?
 * You can either replace or insert from substring at [i-1][i-1], [i-1][j] and [i][j-1].
 * Of course we take min among those and add 1(cost of operation).
 */
func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)

	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	for i := 0; i <= len1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}

	fmt.Printf("%v", dp)

	return dp[len1][len2]
}

// Leetcode #72: Edit Distance
// https://leetcode.com/problems/edit-distance/
func main() {
	var (
		word1 = "horse"
		word2 = "ros"
	)

	fmt.Println(minDistance(word1, word2))
}

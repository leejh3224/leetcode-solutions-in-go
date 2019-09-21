package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
 * Try to solve this by hand first.
 * I'm sure you'll be doing something like below.
 * Given [2, 5, 2, 1, 2] and target = 5, you know you can start making sum from either 1, 2 or 5.
 * then start to connect elements with line.
 *
 * 1 - 2 - 2
 * 1 - 2 - 5 oops! over 5
 * 1 - 2 - 2 ok but it's duplicate...
 * 2 - 2 - 1 also duplicate
 * ...
 *
 * How can we implement it in algorithms?
 * You can use "Backtracking" (read more: https://www.geeksforgeeks.org/backtracking-algorithms/).
 * In short, it's a technique similar to finding paths in maze,
 * meaning "keep going further until you find dead end and if so, go back and repeat it"
 *
 * For this kind of problem, it's really important to find base case and edge cases.
 * Base case is easy to find; if sum of the combination is equal to target then it's all good.
 * What about edge cases?
 * For this problem, not allowing duplicates is hard.
 * You have to do 2 things.
 * 1. sort 'candidates' so you're starting from the smallest number
 * 2. check if you've already checked certain starting elements
 *
 * Below is the code and if you run this, you can see nested lines of function calls.
 * It's there to illustrate the whole process.
 */
func combinationSum2(candidates []int, target int) (ret [][]int) {
	// not to include duplicates
	sort.Ints(candidates)
	helper(candidates, nil, target, 0, 0, &ret, 0)
	return ret
}

/*
 * depth is there to to visualize recursion
 */
func helper(candidates, cur []int, target, sum, start int, ret *[][]int, depth int) {
	// base case
	if sum == target {
		*ret = append(*ret, cur)
		return
	}

	for i := start; i < len(candidates) && (sum+candidates[i]) <= target; i++ {
		fmt.Printf("%scur: %v, remain: %v, to test: %v (at: %v)\n", strings.Repeat("	", depth), cur, candidates[start:], candidates[i], i)

		// we've already done it so skip
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		helper(candidates, append(cur, candidates[i]), target, sum+candidates[i], i+1, ret, depth+1)
	}
}

// Leetcode #40: Combination Sum II
// https://leetcode.com/problems/combination-sum-ii/
func main() {
	const target = 5
	var (
		candidates = []int{2, 5, 2, 1, 2}
		ret        = combinationSum2(candidates, target)
	)

	fmt.Printf("%v", ret)
}

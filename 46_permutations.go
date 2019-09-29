package main

import (
	"fmt"
	"strings"
)

/*
 * Solution using Backtracking
 * original code in Java: https://leetcode.com/problems/permutations/discuss/18239/A-general-approach-to-backtracking-questions-in-Java-(Subsets-Permutations-Combination-Sum-Palindrome-Partioning
 */
func permute(nums []int) (ret [][]int) {
	perm := make([]int, 0)
	visited := make([]bool, len(nums))
	depth := 0
	backtrack(&ret, &visited, &perm, nums, &depth)
	return ret
}

func backtrack(ret *[][]int, visited *[]bool, perm *[]int, nums []int, depth *int) {
	if len(*perm) == len(nums) {
		/*
		 * need to copy temp permutation
		 * if not all temp permutation will point to last permutation
		 *
		 * append([]int{}, (*perm)...) will work too
		 */
		permCopied := make([]int, len(*perm))
		copy(permCopied, *perm)
		*ret = append(*ret, permCopied)
		return
	}

	for i := range nums {
		if !(*visited)[i] {
			*perm = append(*perm, nums[i]) // add

			// visualize callstack
			*depth++
			fmt.Printf("%sin: %v\n", strings.Repeat("    ", *depth), *perm)

			(*visited)[i] = true
			backtrack(ret, visited, perm, nums, depth)
			(*visited)[i] = false
			*perm = (*perm)[:len(*perm)-1] // remove last

			// to visualize callstack
			*depth--
			fmt.Printf("%sout: %v\n", strings.Repeat("    ", *depth), *perm)
		}
	}
}

func main() {
	var input = []int{5, 4, 6, 2}
	fmt.Printf("%v", permute(input))
}

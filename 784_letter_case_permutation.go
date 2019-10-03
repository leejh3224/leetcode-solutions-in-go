package main

import (
	"fmt"
	"strings"
)

/*
 * Backtracking Problem
 */
func letterCasePermutation(S string) []string {
	perms := make([]string, 0)
	arr := strings.Split(S, "")
	helper(&perms, &arr, 0)
	return perms
}

func helper(perms *[]string, arr *[]string, index int) {
	if index == len(*arr) {
		*perms = append(*perms, strings.Join(*arr, ""))
		return
	}

	char := (*arr)[index]

	// you can convert string into rune slice or vice versa
	r := []rune(char)[0]

	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		(*arr)[index] = strings.ToLower(string((*arr)[index]))
		helper(perms, arr, index+1)
		(*arr)[index] = strings.ToUpper(string((*arr)[index]))
		helper(perms, arr, index+1)
	} else {
		helper(perms, arr, index+1)
	}
}

// Leetcode #784: Letter Case Permutation
// https://leetcode.com/problems/letter-case-permutation/submissions/
func main() {
	var input = "a1b2"
	fmt.Printf("%v\n", letterCasePermutation(input))
}

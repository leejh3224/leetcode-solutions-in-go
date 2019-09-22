package main

import (
	"fmt"
)

/*
 * You have to prove that left side of the peak is strictly less than all the elements from right side
 * and strictly greater than elements from left side.
 * You can achieve this by moving pointer from left to right
 * and test if pointer moves in the way we want it to move.
 */
func validMountainArray(A []int) bool {
	n := len(A)
	i := 0

	if n < 3 {
		return false
	}

	// climbing up from left side
	for i+1 < n && A[i] < A[i+1] {
		i++
	}

	// monotonically increasing or decreasing
	if i == 0 || i == n-1 {
		return false
	}

	// climbing down to right side
	for i+1 < n && A[i] > A[i+1] {
		i++
	}

	// if i is at the end then it's all good
	return i == n-1
}

// Leetcode #941: Valid Mountain Array
// https://leetcode.com/articles/valid-mountain-array/
func main() {
	var arr = []int{0, 3, 2, 1}
	var arr2 = []int{3, 5, 5}

	fmt.Printf("%v", validMountainArray(arr))
	fmt.Printf("%v", validMountainArray(arr2))
}

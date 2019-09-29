package main

import "fmt"

func majorityElement(nums []int) (ret int) {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	for k, v := range counter {
		if v > len(nums)/2 {
			ret = k
		}
	}

	return ret
}

// Leetcode #169: Majority Element
// https://leetcode.com/problems/majority-element/
func main() {
	var input = []int{3, 2, 3}
	fmt.Printf("%v\n", majorityElement(input))
}

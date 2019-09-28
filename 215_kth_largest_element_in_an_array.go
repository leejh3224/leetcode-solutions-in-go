package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
 * solution using sorting
 */
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

/*
 * solution using max heap
 */
func findKthLargest2(nums []int, k int) (ret int) {
	maxheap := make(MaxHeap, 0)
	heap.Init(&maxheap)

	for _, val := range nums {
		heap.Push(&maxheap, val)
	}

	for i := 0; i < k; i++ {
		ret = heap.Pop(&maxheap).(int)
	}

	return ret
}

// MaxHeap Implementation

// MaxHeap ...
type MaxHeap []int

// Len ...
func (h MaxHeap) Len() int {
	return len(h)
}

// Less ...
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Pop ...
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	*h = (*h)[:n-1]
	return el
}

// Push ...
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	var input = []int{
		3, 2, 3, 1, 2, 4, 5, 5, 6,
	}

	fmt.Printf("%v", findKthLargest(input, 4))
	fmt.Printf("%v", findKthLargest2(input, 4))
}

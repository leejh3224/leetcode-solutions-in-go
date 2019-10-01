package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * Solution using Floyd's Tortoise and Hare
 * reference: https://en.wikipedia.org/wiki/Cycle_detection#Floyd's_Tortoise_and_Hare
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	p1, p2 := head, head.Next

	for p2 != nil && p2.Next != nil && p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	return p1 == p2
}

// Leetcode #141 Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/
func main() {
	var node = &ListNode{
		Val: 3,
	}
	var node1 = &ListNode{
		Val: 2,
	}
	var node2 = &ListNode{
		Val: 0,
	}
	var node3 = &ListNode{
		Val: 4,
	}
	node.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1

	fmt.Printf("%v\n", hasCycle(node))
}

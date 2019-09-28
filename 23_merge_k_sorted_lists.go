package main

import (
	"container/heap"
	"fmt"
)

/*
 * solution using PriorityQueue
 */
func mergeKLists(lists []*ListNode) (ret *ListNode) {
	pq := make(PriorityQueue, 0)

	for _, node := range lists {
		if node != nil {
			pq = append(pq, node)
		}
	}

	if len(pq) == 0 {
		return nil
	}

	heap.Init(&pq)

	var head *ListNode = &ListNode{}
	ret = head

	for len(pq) > 0 {
		// pull min node from pq
		node := heap.Pop(&pq).(*ListNode)

		// push next node if node has one
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}

		// moving pointers
		head.Next = node
		head = head.Next // set head to latest
	}

	// ret is 'dummy head'
	return ret.Next
}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

/* PriorityQueue implementation */

// PriorityQueue ...
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

// Pop ...
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[:n-1]
	return node
}

// Push ...
func (pq *PriorityQueue) Push(node interface{}) {
	*pq = append(*pq, node.(*ListNode))
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// utility
func printList(root *ListNode) {
	var node *ListNode = root

	for node != nil {
		sep := "->"
		if node.Next == nil {
			sep = ""
		}

		fmt.Printf("%v%s", node.Val, sep)
		node = node.Next
	}
}

// Leetcode #23: Merge k Sorted Lists
// https://leetcode.com/problems/merge-k-sorted-lists/
func main() {
	var list1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	var list2 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	var list3 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	}

	printList(
		mergeKLists([]*ListNode{
			list1, list2, list3,
		}),
	)
}

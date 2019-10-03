package main

import "fmt"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Traversing Tree in order gives you sorted list of nodes
func traverse(node *TreeNode, arr *[]*TreeNode) {
	if node == nil {
		return
	}

	traverse(node.Left, arr)
	*arr = append(*arr, node)
	traverse(node.Right, arr)
}

func kthSmallest(root *TreeNode, k int) int {
	arr := make([]*TreeNode, 0)
	traverse(root, &arr)

	// arr is sorted in ascending order
	return arr[k-1].Val
}

// Leetcode #230: Kth Smallest Element in a BST
// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
func main() {
	var tree = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	fmt.Printf("%v\n", kthSmallest(tree, 1))
}

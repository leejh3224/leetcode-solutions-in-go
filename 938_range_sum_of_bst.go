package main

import "fmt"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(node *TreeNode, L, R int, sum *int) {
	if node == nil {
		return
	}

	if node.Val >= L && node.Val <= R {
		*sum += node.Val
	}

	traverse(node.Left, L, R, sum)
	traverse(node.Right, L, R, sum)
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	traverse(root, L, R, &sum)
	return sum
}

// Leetcode #938: Range Sum of BST
// https://leetcode.com/problems/range-sum-of-bst/
func main() {
	var tree = &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}
	fmt.Printf("%v\n", rangeSumBST(tree, 7, 15))
}

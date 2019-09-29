package main

import "fmt"

/*
 * solution using BFS
 * It's almost same as BFS but we reverse the insertion order when we start from right to left
 */
func zigzagLevelOrder(root *TreeNode) (ret [][]int) {
	if root == nil {
		return ret
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	reverse := false

	for len(queue) > 0 {
		level, n := make([]int, 0), len(queue)

		for i := 0; i < n; i++ {
			var head *TreeNode
			head, queue = queue[0], queue[1:]

			if reverse {
				level = append([]int{head.Val}, level...)
			} else {
				level = append(level, head.Val)
			}

			if head.Left != nil {
				queue = append(queue, head.Left)
			}

			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}

		ret = append(ret, level)
		reverse = !reverse
	}

	return ret
}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Leetcode #103: Binary Tree Zigzag Level Order Traversal
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
func main() {
	var root = TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Printf("%v", zigzagLevelOrder(&root))
}

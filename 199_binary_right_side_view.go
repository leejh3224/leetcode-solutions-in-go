package main

import "fmt"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * You can solve this using BFS.
 * The idea is to grab the last element of given depth.
 * Except that it's just plain BFS.
 */
func rightSideView(root *TreeNode) (ret []int) {
	// edge case
	if root == nil {
		return ret
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		var n = len(q)

		for i := 0; i < n; i++ {
			var cur *TreeNode
			cur, q = q[0], q[1:]

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}

			// node is the last element of given depth
			if i == n-1 {
				ret = append(ret, cur.Val)
			}
		}
	}

	return ret
}

// Leetcode #199: Binary Tree Right Side View
// https://leetcode.com/problems/binary-tree-right-side-view/
func main() {
	var root = TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	fmt.Printf("%v", rightSideView(&root))
}

package main

import "fmt"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * You can use either BFS or DFS.
 * Just find 'left leaves' and keep track of the sum of its value
 */
func sumOfLeftLeaves(root *TreeNode) (ret int) {
	dfs(root, &ret)
	return ret
}

func dfs(node *TreeNode, sum *int) {
	if node == nil {
		return
	}
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		*sum += node.Left.Val
	}
	dfs(node.Left, sum)
	dfs(node.Right, sum)
}

// Leetcode #404: Sum of Left Leaves
// https://leetcode.com/problems/sum-of-left-leaves/
func main() {
	var tree = &TreeNode{
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

	fmt.Printf("%v", sumOfLeftLeaves(tree))
}

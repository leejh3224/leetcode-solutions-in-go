package main

import "fmt"

/*
 * Should remember the definition of BST.
 * Remember that given tree is a BST iff value of "all" the left subtree is less than parent and
 * value of "all" the right subtree is bigger than parent.
 *
 * Thus in order to check BST invariant, we shuld know upperbound and lowerbound
 */
const (
	uintMax = ^uint(0)
	intMax  = int(uintMax >> 1)
	intMin  = -intMax - 1
)

func isValidBST(root *TreeNode) bool {
	return isValid(root, intMin, intMax)
}

func isValid(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return isValid(root.Left, lower, root.Val) &&
		isValid(root.Right, root.Val, upper)
}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Leetcode #98: Validate Binary Search Tree
// https://leetcode.com/problems/validate-binary-search-tree/
func main() {
	var tree = &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}

	fmt.Printf("%v", isValidBST(tree))
}

package main

import "fmt"

// Stack ...
type Stack []*TreeNode

// Push ...
func (s *Stack) Push(x *TreeNode) {
	*s = append(*s, x)
}

// Pop ...
func (s *Stack) Pop() *TreeNode {
	x := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return x
}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return root == nil || compare(root.Left, root.Right)
}

/*
 * Recursive solution
 */
func compare(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil || node2 == nil {
		return node1 == node2
	}

	if node1.Val != node2.Val {
		return false
	}

	return compare(node1.Left, node2.Right) && compare(node1.Right, node2.Left)
}

/*
 * Non recursive solution using Stack
 */
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := Stack{}

	stack.Push(root.Right)
	stack.Push(root.Left)

	for len(stack) > 0 {
		left := stack.Pop()
		right := stack.Pop()

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		stack.Push(left.Left)
		stack.Push(right.Right)
		stack.Push(left.Right)
		stack.Push(right.Left)
	}

	return true
}

// Leetcode #101: Symmetric Tree
// https://leetcode.com/problems/symmetric-tree/
func main() {
	var tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Printf("%v\n", isSymmetric(tree))
	fmt.Printf("%v\n", isSymmetric2(tree))
}

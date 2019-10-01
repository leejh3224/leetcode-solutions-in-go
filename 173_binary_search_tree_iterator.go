package main

import "fmt"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Stack ...
type Stack []*TreeNode

// Pop ...
func (s *Stack) Pop() *TreeNode {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

// Push ...
func (s *Stack) Push(x *TreeNode) {
	*s = append(*s, x)
}

// BSTIterator ...
type BSTIterator struct {
	stack *Stack
}

// Inorder ...
func (this *BSTIterator) Inorder(root *TreeNode) {
	for root != nil {
		this.stack.Push(root)
		root = root.Left
	}
}

// Constructor ...
/*
 * The idea is if you do inorder traversal for given tree,
 * then you'll get sorted list.
 * You can do dfs for tree traversal but sadly it can't be done in O(1) time and O(h) space when h = height of the tree.
 * Instead you can do it with stack.
 */
func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{&Stack{}}
	iterator.Inorder(root)
	return iterator
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := this.stack.Pop()

	if node.Right != nil {
		this.Inorder(node.Right)
	}

	return node.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(*this.stack) != 0
}

// Leetcode #173 Binary Search Tree Iterator
// https://leetcode.com/problems/binary-search-tree-iterator/
func main() {
	var tree = &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	var iterator = Constructor(tree)
	fmt.Printf("%v\n", iterator.Next())
	fmt.Printf("%v\n", iterator.Next())
	fmt.Printf("%v\n", iterator.HasNext())
}

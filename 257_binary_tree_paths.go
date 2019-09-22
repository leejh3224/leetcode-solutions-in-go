package main

import (
	"fmt"
	"strconv"
)

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * You can solve this problem by using "backtracking"(dfs)
 * The idea is to hold the string of path and when you reach the leaf(node with no children),
 * then you completed single path.
 * There is not so many edge cases to handle thus difficulty is "easy",
 * if you've never encountered backtracking problem, it can be little hard.
 */
func binaryTreePaths(root *TreeNode) (ret []string) {
	// edge case
	if root == nil {
		return []string{}
	}

	helper(root, fmt.Sprintf("%d", root.Val), &ret)
	return ret
}

func helper(node *TreeNode, path string, ret *[]string) {
	if node.Left == nil && node.Right == nil {
		*ret = append(*ret, path)
		return
	}

	/*
	 * concat strings in Go: https://goruncode.com/how-to-concatenate-strings-in-go/
	 * convert int to string in Go: https://www.admfactory.com/how-to-convert-an-int-type-to-a-string-type-in-golang/
	 */
	if node.Left != nil {
		helper(node.Left, fmt.Sprintf("%s->%s", path, strconv.Itoa(node.Left.Val)), ret)
	}

	if node.Right != nil {
		helper(node.Right, fmt.Sprintf("%s->%s", path, strconv.Itoa(node.Right.Val)), ret)
	}
}

// Leetcode #257: Binary Tree Paths
// https://leetcode.com/problems/binary-tree-paths/
func main() {
	var (
		node1 TreeNode = TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}
		node2 TreeNode = TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		}
		node3 TreeNode = TreeNode{
			Val:   2,
			Left:  nil,
			Right: &node2,
		}
		root TreeNode = TreeNode{
			Val:   1,
			Left:  &node3,
			Right: &node1,
		}
		path []string
	)

	path = binaryTreePaths(&root)

	fmt.Printf("%v", path)
}

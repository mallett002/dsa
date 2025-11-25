package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left)
	dfs(root.Right)

	return
}

// Binary search -> loops w/ pointers
// Depth First Search -> recursion

package main

import "fmt"

func MaxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// get max of left subtree
	left := MaxDepth(node.Left)
	// get max of right subtree
	right := MaxDepth(node.Right)

	fmt.Printf("left %d; right %d\n", left, right)

	// pick whichever is deeper
	return max(left, right) + 1 // 1 is self node
}

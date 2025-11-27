package main

func MaxValue(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// leaf - just return its value
	if node.Left == nil && node.Right == nil {
		return node.Value
	}

	// recursively get left and right
	left := MaxValue(node.Left)
	right := MaxValue(node.Right)

	// return the largest
	return max(node.Value, left, right)
}

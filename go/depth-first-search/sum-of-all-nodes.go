package main

// Left off: https://www.hellointerview.com/learn/code/depth-first-search/return-values - Solving Problems with Recursion
func SumOfAllNodes(node *TreeNode) int {
	// empty node (parent didn't have a child here, but had one on the other side (ex: no left, but had a right)
	if node == nil {
		return 0
	}

	// leaf node
	if node.Left == nil && node.Right == nil {
		return node.Value
	}

	return SumOfAllNodes(node.Left) + SumOfAllNodes(node.Right) + node.Value
}

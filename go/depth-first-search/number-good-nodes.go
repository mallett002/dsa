package main

import "math"

func helperCountGoodNodes(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}

	// count # of good nodes
	count := 0

	// good, set new max val and increment count
	if node.Value >= maxVal {
		maxVal = node.Value
		count++
	}

	// get left and right recursively
	left := helperCountGoodNodes(node.Left, maxVal)
	right := helperCountGoodNodes(node.Right, maxVal)

	return count + left + right
}

func NumberGoodNodes(node *TreeNode) int {
	return helperCountGoodNodes(node, math.MinInt)
}

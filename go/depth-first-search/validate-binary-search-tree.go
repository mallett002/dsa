package main

import "math"

/*
Valid if
  - Every value to left is less than node
  - Every value to right is greater than node
  - Left and right sub-trees must meet this too
  - Think of it like a window min .... max. Needs to fit in the window

		6
	 2	   8
   1   3  7  9
*/

func doValidateBST(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}

	// needs to be greater than min, and less than max
	if node.Value <= min || node.Value >= max {
		return false
	}

	// left has to be less than max (current node's value)
	isLeftValid := doValidateBST(node.Left, min, node.Value) // max is current node's val

	// right has to be greater than min (current nodes' value)
	isRightValid := doValidateBST(node.Right, node.Value, max) // min is current node's val

	return isLeftValid && isRightValid
}

func ValidateBST(node *TreeNode) bool {
	return doValidateBST(node, math.MinInt64, math.MaxInt64) // huge window to get things started
}

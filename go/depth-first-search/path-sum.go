package main

func HasPathToTarget(node *TreeNode, target int) bool {
	if node == nil {
		return false
	}

	// at leaf, see if target matches leaf's val
	if node.Left == nil && node.Right == nil {
		return node.Value == target
	}

	// if not leaf, decrement target
	target -= node.Value

	return HasPathToTarget(node.Left, target) || HasPathToTarget(node.Right, target)
}

/*
How this works

	1
2		3

target 4

node: 1; left: 2; right: 3
not leaf; decrement target; target 3

node: 2; left: nil; right: nil
is leaf: leaf matches target? leaf 2; target 3; NO return false

node: nil; return false

node: 3; left: nil; right: nil
is leaf: leaf matches target? leaf 3; target 3; YES return true


*/

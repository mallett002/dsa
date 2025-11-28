package main

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	/*
			4
		  2   7
		1  3 6  9

		Target 13

		If tree has path from root to leaf that sums to target return true
		else false
	*/

	var nine TreeNode = TreeNode{Value: 9}
	var six TreeNode = TreeNode{Value: 6}
	var seven TreeNode = TreeNode{Left: &six, Right: &nine, Value: 7}
	var one TreeNode = TreeNode{Value: 1}
	var three TreeNode = TreeNode{Value: 3}
	var two TreeNode = TreeNode{Left: &one, Right: &three, Value: 2}
	var root TreeNode = TreeNode{Left: &two, Right: &seven, Value: 4}

	var meetsTarget bool = HasPathToTarget(&root, 17)
	expected := true

	if meetsTarget != expected {
		t.Fatalf("FAIL: Expected %v; Received %v", expected, meetsTarget)
	}
}

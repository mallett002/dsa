package main

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
	/*
			4
		  2	  7
		1  3    9
		6

		How tall is the tree? or "How many node high is the tree?
		Expected 4 nodes tall

	*/

	var six TreeNode = TreeNode{Value: 6}
	var nine TreeNode = TreeNode{Value: 9}
	var seven TreeNode = TreeNode{Right: &nine, Value: 7}
	var one TreeNode = TreeNode{Value: 1, Left: &six}
	var three TreeNode = TreeNode{Value: 3}
	var two TreeNode = TreeNode{Left: &one, Right: &three, Value: 2}
	var root TreeNode = TreeNode{Left: &two, Right: &seven, Value: 4}

	var total int = MaxDepth(&root)
	expected := 4

	if total != expected {
		t.Fatalf("FAIL: Expected %d; Received %d", expected, total)
	}
}

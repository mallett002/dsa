package main

import (
	"testing"
)

func TestMaxValue(t *testing.T) {
	/*
			4
		  2	  7
		1  3    9

	*/
	var nine TreeNode = TreeNode{Value: 9}
	var seven TreeNode = TreeNode{Right: &nine, Value: 7}
	var one TreeNode = TreeNode{Value: 1}
	var three TreeNode = TreeNode{Value: 3}
	var two TreeNode = TreeNode{Left: &one, Right: &three, Value: 2}
	var root TreeNode = TreeNode{Left: &two, Right: &seven, Value: 4}

	var total int = MaxValue(&root)
	expected := 9

	if total != expected {
		t.Fatalf("FAIL: Expected %d; Received %d", expected, total)
	}
}

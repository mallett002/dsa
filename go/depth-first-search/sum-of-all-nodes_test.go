package main

import (
	"testing"
)

func TestSumOfAllNodes(t *testing.T) {
	/*
			4
		  2	  5
		1  3    2
	*/
	var secondTwo TreeNode = TreeNode{Value: 2}
	var five TreeNode = TreeNode{Right: &secondTwo, Value: 5}
	var one TreeNode = TreeNode{Value: 1}
	var three TreeNode = TreeNode{Value: 3}
	var two TreeNode = TreeNode{Left: &one, Right: &three, Value: 2}
	var root TreeNode = TreeNode{Left: &two, Right: &five, Value: 4}

	var total int = SumOfAllNodes(&root)
	expected := 17

	if total != expected {
		t.Fatalf("FAIL: Expected %d; Received %d", expected, total)
	}
}

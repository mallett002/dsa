package main

import (
	"testing"
)

func TestValidateBSTTrue(t *testing.T) {
	/*
			4
		  2	  7
		1  3 6  9
	*/
	// given
	var six TreeNode = TreeNode{Value: 6}
	var nine TreeNode = TreeNode{Value: 9}
	var seven TreeNode = TreeNode{Right: &nine, Left: &six, Value: 7}
	var one TreeNode = TreeNode{Value: 1}
	var three TreeNode = TreeNode{Value: 3}
	var two TreeNode = TreeNode{Left: &one, Right: &three, Value: 2}
	var root TreeNode = TreeNode{Left: &two, Right: &seven, Value: 4}

	// when
	var isValid bool = ValidateBST(&root)

	// then
	if isValid != true {
		t.Errorf("FAIL: Expected %v; Received %v", true, isValid)
	}
}

func TestValidateBSTFalse(t *testing.T) {
	/*
			4
		  2	  7
		1  3 10  9 --> invalid, 10 is greater than 7
	*/
	// given
	var ten TreeNode = TreeNode{Value: 10}
	var nine TreeNode = TreeNode{Value: 9}
	var seven TreeNode = TreeNode{Right: &nine, Left: &ten, Value: 7}
	var one TreeNode = TreeNode{Value: 1}
	var three TreeNode = TreeNode{Value: 3}
	var two TreeNode = TreeNode{Left: &one, Right: &three, Value: 2}
	var root TreeNode = TreeNode{Left: &two, Right: &seven, Value: 4}

	// when
	var isValid bool = ValidateBST(&root)

	// then
	if isValid != false {
		t.Errorf("FAIL: Expected %v; Received %v", false, isValid)
	}
}

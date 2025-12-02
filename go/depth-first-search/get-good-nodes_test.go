package main

import (
	"reflect"
	"testing"
)

func TestGetGoodNodes(t *testing.T) {
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
	var nodes []*TreeNode = GetGoodNodes(&root)

	// then
	var expected []*TreeNode = []*TreeNode{
		&root,
		&seven,
		&nine,
	}

	if !reflect.DeepEqual(nodes, expected) {
		t.Errorf("FAIL: Expected %v; Received %v", expected, nodes)
	}
}

package main

import (
	"fmt"
	"testing"
)

func createList(count int) *ListNode {
	head := NewListNode(1)
	current := head

	for i := 2; i <= count; i++ {
		current.Next = NewListNode(i)
		current = current.Next
	}

	return head
}

func TestFindMiddleNode(t *testing.T) {
	headTestOne := createList(5)
	headTestTwo := createList(3)

	fmt.Printf("List one %s", PrintList(headTestOne))
	fmt.Println("")
	fmt.Printf("List two %s", PrintList(headTestTwo))
	fmt.Println("")

	var tests = []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{
			"5 nodes should return node with 3",
			headTestOne,
			headTestOne.Next.Next,
		},
		{
			"3 nodes should return node with 2",
			headTestTwo,
			headTestTwo.Next,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var middle *ListNode = FindMiddleNode(tt.head)

			if middle.Val != tt.expected.Val {
				t.Errorf("FAIL: Expected %d; Received %d;", tt.expected.Val, middle.Val)
			}
		})
	}
}

package main

import (
	"fmt"
	"testing"
)

func createList(count int) *ListNode {
	head := NewListNode(1)
	current := head

	for i := 2; i <= count; i++ {
		fmt.Println(i)
		current.Next = NewListNode(i)
		current = current.Next
	}

	return head
}

func TestFindMiddleNode(t *testing.T) {
	headWithCycle := NewListNode(1)
	headWithCycle.Next = NewListNode(2)
	headWithCycle.Next.Next = NewListNode(3)
	headWithCycle.Next.Next.Next = headWithCycle.Next

	headNoCycle := createList(4)

	var tests = []struct {
		name     string
		head     *ListNode
		expected bool
	}{
		{
			"5 nodes should return node with 3",
			headWithCycle,
			true,
		},
		{
			"3 nodes should return node with 2",
			headNoCycle,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var hasCycle bool = DetectCycle(tt.head)

			if hasCycle != tt.expected {
				t.Errorf("FAIL: Expected %v; Received %v;", tt.expected, hasCycle)
			}
		})
	}
}

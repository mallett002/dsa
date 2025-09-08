package main

import (
	// "fmt"
	"testing"
)

func createList() *ListNode {
	head := NewListNode(1)
	head.Next = NewListNode(2)
	head.Next.Next = NewListNode(3)
	head.Next.Next.Next = NewListNode(4)

	return head
}

func TestCountNode(t *testing.T) {
	// given
	head := createList()

	// when
	result := CountNodes(head)

	// then
	if result != 4 {
		t.Errorf("Incorrect. Received: %d Expected: %d ", result, 4)
	}
}

func TestDeleteNode(t *testing.T) {
	// given
	head := createList()

	// when
	newHead := DeleteNode(head, 2)
	count := CountNodes(newHead)
	printedList := PrintList(newHead)

	// then
	if count != 3 {
		t.Errorf("Incorrect. Expected %d but got %d", 3, count)
	}

	exptectedList := "1 -> 3 -> 4"

	if printedList != exptectedList {
		t.Errorf("Incorrect. Expected %s but got %s", exptectedList, printedList)
	}
}

func TestDeleteNode2(t *testing.T) {
	var tests = []struct {
		name     string
		input    int
		expected string
	}{
		{
			"removing first node",
			1,
			"2 -> 3 -> 4",
		},
		{
			"removing second node",
			2,
			"1 -> 3 -> 4",
		},
		{
			"removing last node",
			4,
			"1 -> 2 -> 3",
		},
	}

	for _, tt := range tests {
		head := createList()

		t.Run(tt.name, func(t *testing.T) {
			newHead := DeleteNode(head, tt.input)
			count := CountNodes(newHead)
			printedList := PrintList(newHead)

			if count != 3 {
				t.Errorf("Incorrect. Expected %d but got %d", 3, count)
			}

			if printedList != tt.expected {
				t.Errorf("Expected %s; Received %s", tt.expected, printedList)
			}
		})
	}
}

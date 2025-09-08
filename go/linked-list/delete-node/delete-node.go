package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

func PrintList(head *ListNode) string {
	curr := head
	var strBuilder strings.Builder

	for curr != nil {
		var arrow string = ""

		if curr.Next != nil {
			arrow = " -> "
		}

		strBuilder.WriteString(fmt.Sprintf("%d%s", curr.Val, arrow))

		curr = curr.Next
	}

	return strBuilder.String()
}

func CountNodes(head *ListNode) int {
	count := 0
	curr := head

	for curr != nil {
		count++
		curr = curr.Next
	}

	return count
}

// 1 -> 2 -> 3 -> 4
// 1 -> 3 -> 4 (rm 2)
func DeleteNode(head *ListNode, target int) *ListNode {
	if head.Val == target {
		return head.Next
	}

	var prev *ListNode = nil
	curr := head

	for curr != nil {
		if curr.Val == target {
			prev.Next = curr.Next
			break
		}

		// move along
		prev = curr
		curr = curr.Next
	}

	return head
}

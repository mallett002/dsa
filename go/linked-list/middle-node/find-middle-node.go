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

// 1 2 3 4
// 1,2      2,4
func FindMiddleNode(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

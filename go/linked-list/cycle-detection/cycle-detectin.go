package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

// 1->2->3->4
//		  <-

// fast 2, 4, 4, 4
// slow 1, 2, 3, 4
func DetectCycle(node *ListNode) bool {
	fast := node.Next
	slow := node

	for fast != nil && fast.Next != nil {
		if fast.Val == slow.Val {
			return true
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}

package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

/*
	BFS:
	Uses queue to keep track of nodes it needs to visit
*/

func bfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	// Start by putting the first root in the queue
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:] // remove item at the beginning of the queue

		result = append(result, currNode.Val)

		if currNode.Left != nil {
			queue = append(queue, currNode.Left)
		}

		if currNode.Right != nil {
			queue = append(queue, currNode.Right)
		}
	}

	return result
}

package main

import (
	"container/heap"
)

func KthLargest(nums []int, k int) int {
	// Start a heap
	h := &IntHeap{}

	// add first k items on heap
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}

	// heapify
	heap.Init(h)

	// iterate through remaining elements
	// if current item > root, replace and heapify
	for i := k; i < len(nums); i++ {
		if nums[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}

	}

	// return root at end
	return (*h)[0]
}

/*
[7, 10, 4, 3, 20, 15]  k:3

add first k items and heapify: [7, 10, 4] => [4, 7, 10]
add 3 no (3 not greater than root)
add 20 yes "20 > root (4)" [7, 10, 20]
add 15 yes "15 greater than root (7) [10, 15, 20]

return root (3rd largest number) 10
*/

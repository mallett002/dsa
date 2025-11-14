package main

import (
	"container/heap"
)

// Uses heap setup from int-heap.go

func TopThreeElements(nums []int) []int {
	h := &IntHeap{}

	// put first 3 items in heap
	for i := 0; i < 3; i++ {
		heap.Push(h, nums[i])
	}

	heap.Init(h)

	// iterate through remaining elements
	for i := 3; i < len(nums); i++ {

		// if curr item greater than top, pop top and add curr
		if nums[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}

	result := make([]int, len(*h))
	for i, val := range *h {
		result[i] = val
	}

	return result
}

/*
How this works:

start with this: [9, 3, 7, 1, -2, 6, 8]
Return the top 3 largest

make min heap with first 3 items:
			3
	  7			  9

iterate through rest. if curr > root, pop root and add curr

...iterating:
curr: 1, -2, 6 (6 larger)

result in valid heap:
			6
	  7			  9

curr: 8 (8 larger)
			7
	  9			  8

result: [7, 9, 8]



*/

func TopKElements(nums []int, k int) []int {
	h := &IntHeap{}

	// put first 3 items in heap
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}

	heap.Init(h)

	// iterate through remaining elements
	for i := k; i < len(nums); i++ {

		// if curr item greater than top, pop top and add curr
		if nums[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}

	result := make([]int, len(*h))
	for i, val := range *h {
		result[i] = val
	}

	return result
}

package main

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j] // max heap's less ">". min heap's is "<"
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push and Pop use pointer receivers because they modify the slice's length, not just its contents.
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int)) // assert that x is an int at runtime (panics if isn't)
}

func (h *IntHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]     // Get the last value to return at the end
	*h = (*h)[0 : n-1] // shrink the heap, removing last item

	return x // return that last val
}

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

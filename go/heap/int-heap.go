package main

import (
	"container/heap"
	"fmt"
)

type Heap[T int] struct {
	items []int
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push and Pop use pointer receivers because they modify the slice's length, not just its contents.
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int)) // assert that x is an int at runtime (panics if isn't)
}

func (h *IntHeap) Pop() any {
	curr := *h // set curr to be the current heap
	n := len(curr)

	x := curr[n-1] // Get the last value to return at the end

	*h = curr[0 : n-1] // shrink the heap, removing last item

	return x // return that last val
}

func main() {
	h := &IntHeap{2, 1, 5}

	heap.Init(h)
	heap.Push(h, 3)

	fmt.Printf("root %d\n", (*h)[0])

	for h.Len() > 0 {
		fmt.Printf("%d\n", heap.Pop(h))
	}

	fmt.Println("done")
}

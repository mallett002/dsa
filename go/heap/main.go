package main

import (
	"fmt"
	// "math"
)

/*
    Heap - array with smallest item at index 0
    Can also think of it as binary tree

    Top "K" problems - Find "k" smallest/largest items in collection

    Heap:
    [1, 2, 4, 5, 8, 6, 9]

    visualize like this:
                    1
                2       4
               5 8     6 9

    1 -> 2, 4
    2 ->

{
    1: {
        2: {
            5: {},
            8: {},
        },
        4: {
            6: {},
            9: {},
        },
    },
}


Left:   2 * i + 1
Right:  2 * i + 2
Parent  [(i - 1) / 2] floor division

i = 2
parent: 1 / 2       -> 0 (floor)
left:   2 * i + 1   -> 5
right:  2 * i + 2   -> 6

push: add to heap
pop: remove root element from heap
peak: look at root element w/o removing it
heapify([elements]): convert array to heap in place

*/

// type Heap struct {
//     Items int[]
// }

// func Heapify(list []int) map[int] {
//
// }

// Organizes into a correct min heap
// if child isn't larger than parent, swaps with parent
func BubbleUp(items []int, index int) {
	// if reached the root, return, it's good
	if index == 0 {
		return
	}

	parentIndex := (index - 1) / 2
	parent := items[parentIndex]

	// If child greater than parent, return, it's good
	if items[index] > parent {
		fmt.Println("si senor")
		return
	}

	// swap
	temp := items[parentIndex]
	items[parentIndex] = items[index]
	items[index] = temp

	// keep going
	BubbleUp(items, parentIndex)
}

// Put methods on the heap interface:
type MinHeap struct {
	items []int
}

func (h *MinHeap) Push(item int) {
	h.items = append(h.items, item)
	h.bubbleUp(len(h.items) - 1)
}

func (h *MinHeap) bubbleUp(index int) {
	// if reached the root, return, it's good
	if index == 0 {
		return
	}

	parentIndex := (index - 1) / 2

	// If child greater than parent, return, it's good
	if h.items[index] >= h.items[parentIndex] {
		return
	}

	// swap
	h.items[parentIndex], h.items[index] = h.items[index], h.items[parentIndex] // in go you don't need a temp

	// keep going
	h.bubbleUp(parentIndex)
}

func (h *MinHeap) Pop() {
	if len(h.items) == 0 {
		panic("cannot pop from empty heap")
	}

	// Move the last element to the root and shrink the slice
	lastIndex := len(h.items) - 1
	h.items[0] = h.items[lastIndex] // move the last item to the root
	h.items = h.items[:lastIndex]   // remove last item

	// Bubble down to restore heap property
	h.bubbleDown(0)
}

func (h *MinHeap) bubbleDown(parentIndex int) {
	left := 2*parentIndex + 1
	right := 2*parentIndex + 2

	// compare parent & 2 children to determine which is smallest
	// if the smallest is a child, swap parent with that child
	smallest := parentIndex

	// Compare with left child
	if left < len(h.items) && h.items[left] < h.items[smallest] {
		smallest = left
	}

	// Compare with right child
	if right < len(h.items) && h.items[right] < h.items[smallest] {
		smallest = right
	}

	// If the smallest is parentIndex, we're done
	if smallest == parentIndex {
		return
	}

	// parent wasn't the smallest, swap with smallest and keep bubbling down
	h.items[parentIndex], h.items[smallest] = h.items[smallest], h.items[parentIndex]
	h.bubbleDown(smallest)
}

func (h *MinHeap) Peek() int {
	return h.items[0]
}

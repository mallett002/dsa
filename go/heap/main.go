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

func Push(item int, items []int) {
	// add item to end of list
	// min heap: if less than parent, swap it with parent. keep doing until greater than parent or reach root node
	items = append(items, item)

	BubbleUp(items, len(items)-1)
}

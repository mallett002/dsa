package main

import (
	"reflect"
	"testing"
)

func TestBubbleUp(t *testing.T) {
	items := []int{1, 2, 5, 4, 8, 9}
	newItem := 3

	// Add item for bubble up
	items = append(items, newItem)
	newItemIndex := len(items) - 1

	BubbleUp(items, newItemIndex)

	expected := []int{1, 2, 3, 4, 8, 9, 5}

	if !reflect.DeepEqual(items, expected) {
		t.Errorf("FAIL: Expected %v; Received %v", expected, items)
	}
}

func TestPush(t *testing.T) {
	var heap MinHeap = MinHeap{items: []int{1, 2, 5, 4, 8, 9}}

	heap.Push(3)

	expected := []int{1, 2, 3, 4, 8, 9, 5}

	if !reflect.DeepEqual(heap.items, expected) {
		t.Errorf("FAIL: Expected %v; Received %v", expected, heap.items)
	}
}

func TestPop(t *testing.T) {
	var heap MinHeap = MinHeap{items: []int{1, 3, 2, 7, 8, 4, 9}}

	/*
		        start:
				                1
				            3       2
				          7   8   4   9

				pop: (replace start with end val)
				                9
				            3       2
				          7   8   4

				bubble: (swap 9 and 2)
				                2
				            3       9
				          7   8   4

				bubble: (swap 9 and 4)
				                2
				            3       4
				          7   8   9

	*/

	heap.Pop()

	expected := []int{2, 3, 4, 7, 8, 9}

	if !reflect.DeepEqual(heap.items, expected) {
		t.Errorf("FAIL: Expected %v; Received %v", expected, heap.items)
	}
}

func TestPeek(t *testing.T) {
	var heap MinHeap = MinHeap{items: []int{1, 3, 2, 7, 8, 4, 9}}

	result := heap.Peek()

	expected := 1

	if result != expected {
		t.Errorf("FAIL: Expected %d; Received %d", expected, result)
	}
}

func TestHeapify(t *testing.T) {
	/*
	   Start with an invalid heap:
	           4
	       6       9
	    3    2  8    3

	   step 1: handle last non-leaf node (swap 9 & 3)
	           4
	       6       3
	    3    2  8    9

	   step 2: move to previous non-leaf node and repeat until root of tree reached
	   - replace 6 & 2
	           4
	       2       3
	    3    6   8    9

	   - replace 2 & 4
	           2
	       4       3
	    3    6   8    9

	   step 3: ensure valid heap (replace 4 & 3)
	           2
	       3       3
	    4    6   8    9

	*/

}

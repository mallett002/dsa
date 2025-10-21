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

	heap.PopAI()

	expected := []int{2, 3, 4, 7, 8, 9}

	if !reflect.DeepEqual(heap.items, expected) {
		t.Errorf("FAIL: Expected %v; Received %v", expected, heap.items)
	}
}

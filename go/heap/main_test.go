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

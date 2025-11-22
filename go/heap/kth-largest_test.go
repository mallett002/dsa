package main

import (
	"testing"
)

func TestKthLargest(t *testing.T) {
	arr := []int{5, 3, 2, 1, 4}

	res := KthLargest(arr, 4)
	expected := 2

	if res != expected {
		t.Errorf("FAIL: Expected %d; Received %d", expected, res)
	}
}

func TestKthLargestAgain(t *testing.T) {
	arr := []int{7, 10, 4, 3, 20, 15}

	res := KthLargest(arr, 3)
	expected := 10

	// 3, 4, 7, 10, 15, 20 -> 10 is 3rd largest

	if res != expected {
		t.Errorf("FAIL: Expected %d; Received %d", expected, res)
	}
}

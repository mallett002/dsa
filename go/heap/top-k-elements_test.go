package main

import (
	"reflect"
	"testing"
)

func TestTopThreeElements(t *testing.T) {
	arr := []int{9, 3, 7, 1, -2, 6, 8}

	res := TopThreeElements(arr)
	expected := []int{7, 9, 8} // or [7, 8, 9] or [9, 7, 8]

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("FAIL: Expected %v; Received %v", expected, res)
	}
}

func TestTopKElements(t *testing.T) {
	arr := []int{9, 3, 7, 1, -2, 6, 8}

	res := TopKElements(arr, 4)
	expected := []int{6, 8, 7, 9}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("FAIL: Expected %v; Received %v", expected, res)
	}
}

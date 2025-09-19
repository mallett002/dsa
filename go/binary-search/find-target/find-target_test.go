package main

import (
	"testing"
)

func TestFindTarget(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		{
			"finds index when target present",
			[]int{-1, 0, 3, 5, 9, 12},
			9,
			4,
		},
		{
			"returns -1 when target not in list",
			[]int{-1, 0, 3, 5, 9, 12},
			7,
			-1,
		},
		{
			"returns -1 if empty input",
			[]int{},
			7,
			-1,
		},
		{
			"returns -1 if only 1 item, but not target",
			[]int{5},
			7,
			-1,
		},
		{
			"returns 0 if only 1 item, and is target",
			[]int{5},
			5,
			0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result int = FindIndexOfTarget(test.input, test.target)

			if result != test.expected {
				t.Errorf("FAIL - Expected %d; Received %d", test.expected, result)
			}
		})
	}
}

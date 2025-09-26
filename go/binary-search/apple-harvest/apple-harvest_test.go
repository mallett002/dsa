package main

import "testing"

func TestAppleHarvest(t *testing.T) {
	var cases = []struct {
		name  string
		input struct {
			trees []int
			hours int
		}
		expected int
	}{
		{
			name: "with {3, 6, 7} in 8 hours, Bobby can do 3 apples/hr",
			input: struct {
				trees []int
				hours int
			}{
				trees: []int{3, 6, 7},
				hours: 8,
			},
			expected: 3,
		},
		{
			name: "with {25, 9, 23, 8, 3} in 5 hours, Bobby can do 25 apples/hr",
			input: struct {
				trees []int
				hours int
			}{
				trees: []int{25, 9, 23, 8, 3},
				hours: 5,
			},
			expected: 25,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := AppleHarvest(tc.input.trees, tc.input.hours)

			if result != tc.expected {
				t.Errorf("FAIL; Expected %d; Received %d", tc.expected, result)
			}
		})
	}

}

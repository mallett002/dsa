package main

import "fmt"

func main() {
	var input [7]uint8 = [7]uint8{1, 3, 4, 6, 8, 10, 13}
	var target uint8 = 13

	res := twoSum(input[:], target)

	fmt.Printf("\n%v", res)
}

func twoSum(input []uint8, target uint8) bool {
	var left uint8 = 0
	var right uint8 = uint8(len(input) - 1)

	for left != right {
		leftNum := input[left]
		rightNum := input[right]

		sum := leftNum + rightNum

		if sum == target {
			return true
		}

		if sum > target {
			right--
		} else if sum < target {
			left++
		}

		if left == right {
			break
		}
	}

	return false
}

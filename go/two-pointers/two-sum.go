package main

import "fmt"

func main() {
	var nums [7]uint8 = [7]uint8{1, 3, 4, 6, 8, 10, 13}
	var target uint8 = 13

	res := twoSum(nums[:], target)

	fmt.Printf("\n%v", res)
}

func twoSum(sortedNums []uint8, target uint8) bool {
	left := 0
	right := len(sortedNums) - 1

	for left < right {
		sum := sortedNums[left] + sortedNums[right]
		fmt.Printf("sum %v", sum)

		if sum == target {
			return true
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return false
}

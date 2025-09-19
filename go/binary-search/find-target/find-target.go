package main

// Binary search for target - Return its index
func FindIndexOfTarget(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// if item is less than target - it's in the right half.
		// move left pointer up to one past the mid to get the right half
		if nums[mid] < target {
			left = mid + 1
		} else { // in left half, move right pointer back to one past mid
			right = mid - 1
		}
	}

	return -1
}

package main

import (
	"fmt"
)

/*
Array where each item is a type of fruit
Find longest subarray of fruits with at most 2 different types of fruits in it.
Must stop when you encounter 3 different types of fruits

-- How works:
state {}
start 0
end 0
max 0

loop over end
	set state: amount of this type of fruit

	while not valid (more than 2 types of fruits in state)
		bring up start
		  - handle state for start
		  - if start is 0 delete it from state

	if count in state > max, set new max

return max

*/

func main() {
	fruits := [7]int{3, 3, 2, 1, 2, 1, 0}

	res := fruitIntoBaskets(fruits[:])

	fmt.Printf("\n%d", res)
}

func fruitIntoBaskets(fruits []int) int {
	basket := make(map[int]int)
	start := 0
	maxFruit := 0

	for end := 0; end < len(fruits); end++ {
		basket[fruits[end]]++ // leverage 0 int value of golang. simply increment, if was 0, will be 1 now

		for len(basket) > 2 { // while basket has more than 2 different types of fruits (not valid)
			basket[fruits[start]]--

			if basket[fruits[start]] == 0 {
				delete(basket, fruits[start])
			}

			start++
		}

		basketSize := (end - start) + 1

		if basketSize > maxFruit {
			maxFruit = basketSize
		}
	}

	return maxFruit
}

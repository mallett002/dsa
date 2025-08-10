package main

import (
	"fmt"
)

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

		for len(basket) > 2 {
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

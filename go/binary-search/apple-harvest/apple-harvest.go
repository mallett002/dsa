package main

import (
	// "fmt"
	"math"
)

/*
Wants to complete fixed # apples/hr
Can only do one tree/hr

	if finish tree b4 hour, still have to wait until hr is up

Ex:

	3 apples on a tree; does 1 apple/hr; takes 3 hrs
	3 apples on a tree; does 2 apple/hr; takes 2 hrs
	    [hr1: 2 apples] [hr2: finishes the 1 apple]

Input: applesPerTree, hours
Find slowest rate of apples/hr Bobby can harvest and still complete all trees in "hours" hours

Ex: [3,6,7] 8

1? 3,6,7: 16hrs
  - [tree1: 3 apples, 3hrs], [tree2: 6 apples, 6hrs]

2? 2,3,4: 9hrs
  - [tree1: 3 apples, 2hrs], [tree2: 6 apples, 3hrs], [tree3: 7 apples, 4hrs]

3? 1,2,3: 6hrs [VALID]

ceil( 7/2 ) = 4

	(7 + 2-1) / 2
	  8 / 2 = 4
*/

func ceilDivide(num int, div int) int {
	return int(math.Ceil(float64(num) / float64(div)))
}

// ------- BRUTE FORCE -----------------------------
// see how long takes to do this rate
func timeTakes(trees []int, rate int) int {
	hours := 0

	for i := 0; i < len(trees); i++ {
		// ceil division (tree + rate-1) / rate
		//	ex: ceil( 6/2 ) ->	(6 + (2 - 1)) / 2	->	(6+1) / 2	-> 7 / 2 -> 3 in golang (keeps int)
		//hours += (trees[i] + rate - 1) / rate // ceil division
		hours += ceilDivide(trees[i], rate)
	}

	// 9/4 -> (9+4-1)/4 -> (9+3)/4 -> 12/4 -> 3
	// 9/4 -> 2.25 round up -> 3
	return hours
}

func AppleHarvest(apples []int, hours int) int {
	rate := 1

	for timeTakes(apples, rate) > hours {
		rate++
	}

	return rate
}

// Binary search approach -------------------------
// Ex: [3,6,7] 8

// apples/hr:
// 1, 2, 3, 4, 5, 6, 7 (most is 7, larget tree)

// 1:
// left: 0;
// right: len - 1;
// mid: index 4 (val 4)
//		     x
// [1, 2, 3, 4, 5, 6, 7]
// 4app/hr -> 1, 2, 2: 5hrs VALID
// eliminate all greater (right mid + 1)
// [1, 2, 3, 4, 5]

// 2: new mid
//        x
// [1, 2, 3, 4, 5]
// 3app/hr -> 1, 2, 3: (5hrs VALID)
// eliminate all greater (right mid + 1)
// [1, 2, 3 ]

// 3: new mid
//     x
// [1, 2, 3]
// 2app/hr -> 2, 3, 4: (9 hrs: not valid)
// eliminate mid and less (left mid + 1)

// new mid
// [3]

// if left === right return apples[left]

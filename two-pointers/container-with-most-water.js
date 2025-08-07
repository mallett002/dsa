function getMaxArea(heights) {
	let left = 0;
	let right = heights.length - 1;

	let maxArea = 0;

	while (left < right) {
		// get smallest num from heights[left] vs heights[right]
		// get area (dist apart indexes, right - left)
		//	smallest * dist = area

		const smallest = Math.min(heights[left], heights[right]);
		
		const dist = right - left;
		const area = smallest * dist;

		if (area > maxArea) {
			maxArea = area;
		}

		// if left greater than right, move right back else move left up
		// always moving the one that's smaller to try and find one larger
		if (heights[left] > heights[right]) {
			right--;
		} else {
			left++;
		}
	}
	
	return maxArea;

}

/*
ending at right:
l  r  d
3, 2, 8 => 16
4, 2, 7 => 14
1, 2, 6 => 6
2, 2, 5 => 10
2, 2, 4 => 8
4, 2, 3 => 6

ending at left
l  r  d
3, 2, 8 => 16
3, 3, 7 => 21
3, 1, 6 => 6
3, 2, 5 => 10
3, 2, 4 => 10
3, 1, 3 => 3

*/

const result = getMaxArea([3,4,1,2,2,4,1,3,2]);

console.log(`expected 21; result: ${result}`);

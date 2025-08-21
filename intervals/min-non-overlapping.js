const theVals = [[1,5],[2,6],[3,7],[4,8]];  // 2

// Find the minimum # of intervals to remove so there are no overlapping intervals:
function findMinNumIntervalsToRemove(intervals) {
	const sorterdByEndTime = intervals.toSorted((a, b) => a[1] - b[1]); // by start time
    
	let countNonOverlapping = 1;

	for (let i = 1; i < sorterdByEndTime.length; i++) {
		const curr = sorterdByEndTime[i];
		
		if (curr[0] >= sorterdByEndTime[i-1][1]) { // no overlapp
			countNonOverlapping++;	
		}
	}

	const res = sorterdByEndTime.length - countNonOverlapping;

	return res;
}


// Their solution:
function nonOverlappingIntervals(intervals) {
    if (intervals.length === 0) {
        return 0;
    }

    intervals.sort((a, b) => a[1] - b[1]);
    let end = intervals[0][1];
    let count = 1;

    for (let i = 1; i < intervals.length; i++) {
        // Non-overlapping interval found
        if (intervals[i][0] >= end) {
            end = intervals[i][1];
            count++;
        }
    }

    return intervals.length - count;
}


function findOnesThatOverlapInstead(intervals) {
    intervals.sort((a, b) => a[1] - b[1]);

	let count = 0;

	for (let i = 1; i < intervals.length; i++) {
		if (intervals[i][0] < intervals[i-1][1]) { // overlap
			count++;			
		}
		
	}

	return count;
}

console.log(findMinNumIntervalsToRemove(theVals))
console.log(nonOverlappingIntervals(theVals))
console.log(findOnesThatOverlapInstead(theVals))


// Data:
// ...x-x.............	4,6
// ..........x.....x..	11,17
// .x...............x.	2,18
// ......x..x.........	7,10

// By starttime:
// .x...............x.	2,18
// ...x-x.............	4,6
// ......x..x.........	7,10
// ..........x.....x..	11,17

// By endtime:
// ...x-x.............	4,6
// ......x..x.........	7,10
// ..........x.....x..	11,17
// .x...............x.	2,18

const cases = [
	{input: [[1,2],[2,3],[3,4],[1,3]], expected: 1},
	{input: [[1,2],[1,2],[1,2]], expected: 2},
	{input: [[1,2],[2,3]], expected: 0},
	{input: [], expected: 0},
	{input: [[1,5],[2,6],[3,7],[4,8]], expected: 3},
	{input: [[1,2],[1,3],[2,2],[2,3],[3,4]], expected: 2},
	{input: [[1,4],[5,6],[6,8],[7,9],[10,12],[11,13]], expected: 2},
	{input: [[1,10],[2,3],[4,5],[6,7],[8,9]], expected: 1},
	{input: [[5,6],[1,2],[3,4],[2,3],[7,8]], expected: 0},
	{input: [[2,2],[2,2],[2,3]], expected: 2},
];


for (const {input, expected} of cases) {
	console.log(`findMinNumIntervalsToRemove: ${findMinNumIntervalsToRemove(input)}; expected: ${expected}`)	
	console.log(`nonOverlappingIntervals: ${nonOverlappingIntervals(input)}; expected: ${expected}`)	
	console.log(`findOnesThatOverlapInstead: ${findOnesThatOverlapInstead(input)}; expected: ${expected}`)	
	console.log('\n')
}


//| Test Case               | Input                                       | Expected Removals | Comments                        |
//| ----------------------- | ------------------------------------------- | ----------------- | ------------------------------- |
//| Basic example           | `[[1,2],[2,3],[3,4],[1,3]]`                 | 1                 | Overlap only \[1,3]             |
//| All identical intervals | `[[1,2],[1,2],[1,2]]`                       | 2                 | Two must go                     |
//| Just touching endpoints | `[[1,2],[2,3]]`                             | 0                 | No overlap                      |
//| Empty list              | `[]`                                        | 0                 | Trivial                         |
//| Chained overlaps        | `[[1,5],[2,6],[3,7],[4,8]]`                 | 3                 | Remove all but one              |
//| Mixed overlaps/gaps     | `[[1,2],[1,3],[2,2],[2,3],[3,4]]`           | 2                 | Need to drop two specific ones  |
//| Gaps with overlaps      | `[[1,4],[5,6],[6,8],[7,9],[10,12],[11,13]]` | 2                 | Two separate overlap zones      |
//| Nested intervals        | `[[1,10],[2,3],[4,5],[6,7],[8,9]]`          | 1                 | Drop big encompassing one       |
//| Unsorted input          | `[[5,6],[1,2],[3,4],[2,3],[7,8]]`           | 0                 | Sorting helps reveal no overlap |
//
//| Zero-length intervals   | `[[2,2],[2,2],[2,3]]`                       | 2                 | Edge-case handling needed       |

// [[1,5],[3,6],[8,10],[15,18]]
// [[1,5],[3,9],[8,9],[15,18]]

// Sort by start
// keep track of merged intervals
// loop over each interval
// if current start is before the last end, merge 
// 	update last merged item to be [currStart, greatest between currEnd and merge's end]
// else put the curr interval in merged list
// return merged

function mergeIntervals(intervals) {
	const sorted = intervals.sort((a, b) => a[0] - b[0]);

	const merged = [];

	for (const interval of sorted) {
		const currStart = interval[0];

		// if prev end is before this one's start ~ no overlap
		if (!merged.length || merged[merged.length - 1][1] < currStart) {
			merged.push(interval);			
		} else { // overlap ~ this one's start if before last one's end time
			// adjust the last interval in merged	
			// set it to be 
			const mergedStart = merged[merged.length - 1][0];
			const newMergedEnd = Math.max(merged[merged.length - 1][1], interval[1]);

			merged[merged.length - 1] = [mergedStart, newMergedEnd];
		}
	}

	return merged;
}


const intervals = [[1,5],[3,9],[8,9],[15,18]]; // [[1,9],[15,18]]
console.log(JSON.stringify( mergeIntervals(intervals), null, 2 ));


// [x...x] // [ 1,5 ]
// [...x.....x] // overlap [ 3,9 ]
// [..............x...x] [ 15,18 ]


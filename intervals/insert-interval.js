/* 
Task: insert the new interval and merge with others as needed

intervals = [[1,3],[6,9]]
newInterval = [2,5]
outputs: [[1,5],[6,9]]

Approach:
1. Add all the intervals starting before newInterval to merged.
2. Merge all overlapping intervals with newInterval and add that merged interval to merged.
3. Add all the intervals starting after newInterval to merged.


	old		x.x.........	[1,3]
	new		.x..x.......	[2,5] overlap: start < prevEnd
	old		.....x...x..	[6,9]

*/

function insertInterval(sortedIntervals, newInterval) {
	const merged = [];
	let i = 0;

	// 1. Add all the intervals starting before newInterval to merged.
	// while currEnd < newStart
	while (i < sortedIntervals.length && sortedIntervals[i][1] < newInterval[0]) {
		merged.push(sortedIntervals[i]);
		i++;
	}
	
	// 2. Merge all overlapping intervals with newInterval and add that merged interval to merged.
	// while currEnd >= newStart
	while (i < sortedIntervals.length && newInterval[1] >= sortedIntervals[i][0]) {
		newInterval[0] = Math.min(sortedIntervals[i][0], newInterval[0]);
		newInterval[1] = Math.max(sortedIntervals[i][1], newInterval[1]);

		i++;
	}
	merged.push(newInterval);
	
	// 3. Add all the intervals starting after newInterval to merged.
	// while newEnd < currStart
	while (i < sortedIntervals.length && newInterval[0] < sortedIntervals[i][0]) {
		merged.push(sortedIntervals[i]);
		i++;
	}
	
	return merged;
}



console.log(`Expected: ${JSON.stringify( [[1,3],[4,10],[11,15]] )}; Got ${JSON.stringify(insertInterval([[1,3],[4,6],[6,7],[8,10],[11,15]], [5,8]))}`)


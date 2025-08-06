const nums = [1,3,4,6,8,10,13];
const target = 6;

function twoSum(nums, target) {
	let left = 0;
	let right = nums.length - 1;

	while (left < right) {
		const sum = nums[right] + nums[left];

		if (sum === target) {
			return true;
		} else if (sum > target) {
			right--;	
		} else {
			left++;
		}
	}

	return false;
}

const res = twoSum(nums, target);

console.log(`nums: ${nums}; target: ${target}`);
console.log('Expected false:', res);


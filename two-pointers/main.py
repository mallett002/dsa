# nums = [1,3,4,6,8,10,13], target = 13
def two_sum(nums, target):
    left = 0
    right = len(nums) - 1

    while left < right:
        sum = nums[left] + nums[right]

        if sum == target:
            return True

        if sum > target:
            right -= right
            continue

        if sum < target:
            left += left
            continue

    return False

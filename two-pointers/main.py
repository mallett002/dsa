# nums = [1,3,4,6,8,10,13], target = 13
def two_sum(nums, target):
    # pointers on the outside indexes
    left = 0
    right = len(nums) - 1

    # if curr sum > target: all with current "right" will be greater. move "right" back one
    # if curr < target: all with left will be smaller. move "left" up one

    while left < right: # while they haven't met
        sum = nums[left] + nums[right]

        if sum == target:
            return True

        if sum > target:
            right -= right;
            continue;
        
        if sum < target:
            left += left;
            continue;
    
    return False

# Write a function to count the number of triplets in an integer array nums that could form the sides of a triangle.
# The triplets do not need to be unique.
def triangle_numbers(nums):
    nums.sort()
    count = 0

    # loop backwards, from end to 3rd from beginning (index 2) - leave room for left and right pointers
    # [2, 9, 10, 12]
    for i in range(len(nums) - 1, 1, -1):
        left = 0
        right = i - 1

        while left < right:
            if nums[left] + nums[right] > nums[i]:
                count += right - left # all nums from left to right are valid
                right -= 1 # decrement right to see if this one also works
            else:
                left += 1 # increment left

    return count

# i: 4
# l: 0
# r: 3

# count = 6

# [4, 6, 9, 11, 15, 18]
#  0  1  2   3   4   5

# l   r   i
# ---------
# 4 + 15 > 18: yes; count += 4; right -= 1;
# 4 + 11 > 18: no; left += 1  
# 6 + 11 > 18: yes; count += 2; right -= 1;
# 6 + 9 > 18: no; left += 1; left == right; break

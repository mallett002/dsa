# Write a function to count the number of triplets in an integer array nums that could form the sides of a triangle.
# The triplets do not need to be unique.
def triangle_numbers(nums):
    nums.sort()

    count = 0

    for i in range(len(nums) - 2):
        left = i + 1
        right = len(nums) - 1

        while left < right:
            # want a + b > c (valid triangle)
            # a + b < c     --> 
            # move left up (try and get bigger number)
            if nums[i] + nums[left] < nums[right]:
                left += 1

            if nums[i] + nums[left] > nums[right]:
                count += 1
                left += 1
    
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

# 4 + 11 > 15: yes... 

def triangle_numbers2(nums):
    nums.sort()

    count = 0

    start = len(nums) - 1
    stop = 1
    step = -1

    for i in range(start, stop, step):
        left = 0
        right = i - 1

        while left < right:
            if nums[left] + nums[right] > nums[i]: # found triangle
                count += right - left # all nums in-between right and left will also be triangles (bc sorted)
                right -= 1 # move right back (see if next # is big enough to create triangle)
            
            else: # no triangle
                # all nums btw left and right are also not triangles
                # [4, 6, 13]  
                left += 1 # move left up to try and find a bigger number to use
    
    return count


input = [11, 4, 9, 6, 15, 18]
# [4, 6, 9, 11, 15, 18]
#  0  1  2   3   4   5
print(triangle_numbers2(input))
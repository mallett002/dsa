# find combinations of 3 numbers that sum to 0
# ex: 
#   input: nums = [-1,0,1,2,-1,-1]
#   output: [[-1,-1,2],[-1,0,1]]  note: ([-1,0,1] occurs 2x, so removed duplicate)

# solution:
# iterate through each element in the array
# find other 2 items that sum to the negative of current number (means they net to 0)
# maybe easier to just think that all 3 nums net to 0 instead of ^^
# use two-pointer technique to solve
def three_sum(nums):
    nums.sort()

    result = []

    for i in range(len(nums) - 2):
        if (i > 0) and nums[i] == nums[i - 1]:
            continue

        left = i + 1
        right = len(nums) - 1

        while left < right:
            net = nums[i] + nums[left] + nums[right]

            if net < 0:
                left += 1

            elif net > 0:
                right -= 1

            else:
                result.append([nums[i], nums[left], nums[right]])

                while left < right and nums[left] == nums[left + 1]:
                    left += 1

                while left < right and nums[right] == nums[right - 1]:
                    right -= 1

                left += 1
                right -= 1

    return result


print(three_sum([-1, 0, 1, 2, -1, -1]))

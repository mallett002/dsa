def two_sum(nums, target):
    left = 0
    right = len(nums) - 1

    while left < right:
        sum = nums[left] + nums[right]
        print(f" left: {left}; right: {right}; sum: {sum};")

        if sum == target:
            return True

        if sum < target:
            left += 1
        else:
            right -= 1

    return False


nums = [1, 3, 4, 6, 8, 10, 13]
target = 6

res = two_sum(nums, target)
print(res)

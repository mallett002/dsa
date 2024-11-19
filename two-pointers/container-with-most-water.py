# Given an integer input array heights representing the heights of vertical lines,
# write a function that returns the maximum area of water that can be contained by 
# two of the lines (and the x-axis). 
# The function should take in an array of integers and return an integer.

# My language:
# heights is an array of ints
# each int represents a wall or barrier that contains water
# each index in the array represents x-axis
# using height walls, and the x-axis find the 2 heights that, along w/ x-axis, hold the most water (area)
# return that area (int)
def max_area(heights):
    left, right = 0, len(heights) - 1
    curr_max_area = 0

    while left < right:
        width = right - left
        height = min(heights[left], heights[right])
        area = width * height

        current_max_area = max(curr_max_area, area)

        # if left wall smaller than right wall, move left wall up.
        # Find the smaller height (since that determines the top of the water)
        # Since the smaller wall is the bottleneck, try and replace it with a taller wall
        # all other values with this smaller wall will be equal to this at best, or less
        if heights[left] < heights[right]:
            left += 1

        else: # else move right wall back
            right -= 1

    return current_max_area








# me trying to solve:
def max_area_2(heights):
    left, right = 0, len(heights) - 1
    max_area = 0

    while left < right:
        # Get the area
        height = min(heights[left], heights[right])
        width = right - left
        area = width * height

        # set new max_area
        if area > max_area:
            max_area = area
        
        # replace the smaller height (hoping to find a taller wall)
        if heights[left] < heights[right]:
            left += 1
        else:
            right -= 1

    return max_area
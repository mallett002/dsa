# https://www.hellointerview.com/learn/code/sliding-window/variable-length

# gets the size of largest subarray that has no more than 2 diff types of fruits
# fruits []int

# Naive approach
def max_fruits_naive(fruits):
    max_count = 0

    for i in range(0, len(fruits) - 1):
        for j in range(i + 1, len(fruits)):
            sub = set(fruits[i : j+1])

            if len(sub) > 2:
                break
            
            else: #valid
                sub_size = (j-i) + 1
                max_count = max(max_count, sub_size)
    
    return max_count


# max_fruits_naive([3, 3, 2, 1, 2, 1, 0])


# Sliding window technique:
#    max_count; state: {fruit: count}; start, end;
#    expand window (increment end)
#    check for max_count, update if needed
#    expand window (increment end)
#    state has more than 2 keys? shrink window (increment start)

def max_fruits_sliding_window(fruits):
    start = 0
    max_count = 0
    state = {}

    for end in range(len(fruits)):
        state[fruits[end]] = state.get(fruits[end], 0) + 1

        # while not valid subarray
        # shrink window and update state
        while len(state) > 2: 
            state[fruits[start]] -= 1

            if state[fruits[start]] == 0:
                del state[fruits[start]]
            start += 1

        # now we know it's valid, update max_count accordingly
        max_count = max(sum(state.values()), max_count)

    return max_count

result = max_fruits_sliding_window([3, 3, 2, 1, 2, 1, 0])
print(result)





# Sliding window template
def sliding_window(items):
    start = 0
    state = {}
    max_count = 0

    for end in range(len(items)):

        # while state is not valid
        # while some case that makes it not valid
            start += 1
            # update state accordingly
    
        # state is now valid
        # update max_count accordingly
        # max_count = max(max_count, end - start + 1)

    return max_count


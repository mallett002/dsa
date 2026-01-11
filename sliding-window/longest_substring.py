def longest_substring_without_repeating_chars(str):
    start = 0
    max_count = 0
    state = {}

    for end in range(len(str)):
        state[str[end]] = state.get(str[end], 0) + 1

        while state.get(str[end], 0) > 1:
            state[str[end]] -= 1
            start += 1

        max_count = max(max_count, end - start + 1)

    return max_count


res = longest_substring_without_repeating_chars("substring")

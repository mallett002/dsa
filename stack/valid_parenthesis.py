def is_valid_parens(str):
    lookup = {
        "(": ")",
        ")": "(",
        "{": "}",
        "}": "{",
        "[": "]",
        "]": "[",
    }

    stack = []

    for char in str:
        # Is the top of the stack's pair this char?
        # If so pop from stack
        # else add to stack

        if len(stack) and char == lookup[stack[0]]:
            stack.pop(0)
        else:
            stack.insert(0, char)

    return len(stack) == 0


input = "(){({})}"
print(is_valid_parens(input))

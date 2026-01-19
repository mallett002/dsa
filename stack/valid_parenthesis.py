def is_valid_parens(str):
    lookup = {")": "(", "}": "{", "]": "["}

    stack = []

    for char in str:
        if char in lookup:  # if is not an opening char
            if len(stack) == 0 or stack[len(stack) - 1] != lookup[char]:
                return False

            stack.pop()

        else:
            stack.append(char)

    return len(stack) == 0
    # or "return not stack"


input = "(){}"
print(is_valid_parens(input))

# loop over chars
#   is char opening? put in stack
#   else check stack:
#       if nothing in stack or this char isn't the match for top of stack:
#           return false
#       else is match, pop from stack
# get to end and stack empty, true

/*
Given an input string s consisting solely of the characters '(', ')', '{', '}', '[' and ']', determine whether s is a valid string.
A string is considered valid if every opening bracket is closed by a matching type of bracket and in the correct order,
and every closing bracket has a corresponding opening bracket of the same type.


Approach:
- Use stack
- loop through string
- if curr char is an opening brace, put on stack
- if not opening brace:
    - If top of stack isn't curr char's match return false (found that this string ins't valid)
    - Else found its match. Remove from the stack and continue




Example
Input: s = "(){({})}"
Output: true

Example
Input: s = "(){({}})"
Output: false

 */

// First attempt:
function isValid(str) {
    const charMap = {
        ")": "(",
        "}": "{",
        "]": "[",
    };

    const stack = [];

    for (let i = 0; i < str.length; i++) {
        const currChar = str[i];

        if (charMap[currChar]) {
            // it's a closing bracket
            if (stack.pop() !== charMap[currChar]) {
                return false;
            }
        } else {
            // it's an opening bracket
            stack.push(currChar);
        }
    }

    return !stack.length;
}



console.log(`Expected true; Got: ${isValid("(){({})}")}`);
console.log(`Expected false; Got: ${isValid("(){({}})")}`);



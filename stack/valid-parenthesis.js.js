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
- if at the end the stack is empty, we found all the matches so it's a valid string


Example
Input: s = "(){({})}"
Output: true

Example
Input: s = "(){({}})"
Output: false

 */

function isValidString(str) {
    const mapping = { ')': '(', '}': '{', ']': '[' };
    const stack = [];

    for (const char of str) {
        if (char in mapping) {
            if (!stack.length || stack[stack.length -1] !== mapping[char]) {
                return false;
            }

            // found match - rm from stack
            stack.pop();

        } else {
           stack.push(char); 
        }
        
    }

    // if we have cleared the stack we found a match for everything ~ valid string
    return stack.length === 0;
}

console.log(`Expected true; Got: ${isValidString("(){({})}")}`);
console.log(`Expected false; Got: ${isValidString("(){({}})")}`);



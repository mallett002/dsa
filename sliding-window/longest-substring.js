// state, start, longest
//
// loop end
//      update stte with new end
//
//      while not valid
//          update start state
//          increment start
//
//      update max
//
//  return state

function longestSubstrintWithoutRepeat(str) {
    let start = 0;
    let charCount = {}; // # times char appears in window
    let max = 0;

    for (let end = 0; end < str.length; end++) {
        charCount[str[end]] = (charCount[str[end]] ?? 0) + 1; // update the charCount
        
        // if not valid: move up start and handle its state
        while (charCount[str[end]] > 1) {
            // update start's state:
            charCount[str[start]]--;
            // move up start
            start++;
        }

        max = Math.max(max, (end - start) + 1);
    }

    return max;
}

console.log(longestSubstrintWithoutRepeat('eghghhgg'))

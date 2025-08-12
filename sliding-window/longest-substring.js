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

/*
  {}
*/


function longestSubstringWithoutRepeat(str) {
    let start = 0;
    let longest = 0;
    const chars = [];

    for (let end = 0; end < str.length; end++) {
        chars.push(str[end]);

        // valid if every char is unique
        let isValid = new Set(chars).size === ( end - start ) + 1;

        while (!isValid) {
            // rm from state
            chars.shift();
            start++;

            isValid = new Set(chars).size === ( end - start ) + 1;
        }

        longest = Math.max(longest, (end - start) + 1);
    }

    return longest;
}


console.log(longestSubstringWithoutRepeat('substring'));

//const it = new Set();

//it.add('b');
//
//console.log(it);
//
//it.delete('a');
//
//console.log(it)

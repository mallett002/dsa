// Sliding window framework:
// init vars
//      start, state, max
//
// loop over array (end)
//      set state
//
//      while invalid
//          decrement/remove start from state
//          increment start
//
//      update new max
// 
// return max

// Ex:
// fruits = [3, 3, 2, 1, 2, 1, 0]
//           0  1  2  3  4  5  6  (len 7)
// longest subarray with only 2 diff types of fruits
function fruitIntoBaskets(fruits) {
    let start = 0;
    let maxFruit = 0;
    let basket = {}

    for (let end = 0; end < fruits.length; end++) {
        // start by updating state (basket): current count for this fruit (or 0 if not exists) + 1
        basket[fruits[end]] = (basket[fruits[end]] ?? 0) + 1;

        while (Object.keys(basket).length > 2) {
            // decrement (or remove) start's fruit in state
            // move up start of window
            basket[fruits[start]]--;

            if (!basket[fruits[start]]) {
                delete basket[fruits[start]]
            }

            start++;
        }

        // set new max fruit
        maxFruit = Math.max(maxFruit, end - start + 1);
    }

    return maxFruit;
}

console.log('max:', fruitIntoBaskets([3, 3, 2, 1, 2, 1, 0]))

// Doing it again:
function again(fruits) {
    let start = 0;
    let max = 0;
    const basket = {};

    for (let end = 0; end < fruits.length; end++) {
        basket[fruits[end]] = (basket[fruits[end]] ?? 0) + 1;
        
        while (Object.keys(basket).length > 2) {
            basket[fruits[start]]--;

            if (basket[fruits[start]] === 0) {
               delete basket[fruits[start]]; 
            }

            start++;
        }

        max = Math.max(max, end - start + 1);
    }

    return max;
}



console.log('again:', again([3, 3, 2, 1, 2, 1, 0]))




















// longest subarray with only 2 diff types of fruits
// let start = 0; state = 0; max = 0;
//
// for let end = 0
//   increment state for fruit at state[end]      
//
//   while window not valid
//     move up start
//     handle state for start
//
//   set new max
//
// return max






















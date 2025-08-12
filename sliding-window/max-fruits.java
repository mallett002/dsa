import java.util.HashMap;
import java.util.Map;

// basket (state)
// start
// max
//
// loop fruit (end)
//   update state (end's count)
//
//   while not valid (state has more than 2 items in it)
//     handle start in state (decrement it's count)
//     if 0, delete it from state
//     increment start
//
//   update max
//
// return max

public class MaxFruits {

   public int fruitIntoBaskets(int[] fruits) {
     Map<Integer, Integer> basket = new HashMap<>();
     int start = 0; 
     int maxFruit = 0;

     for (int end = 0; end < fruits.length; end++) {
        basket.put(fruits[end], basket.getOrDefault(fruits[end], 0) + 1);

        while (basket.size() > 2) {
            int fruitAtStart = fruits[start];
            basket.put(fruitAtStart, fruitAtStart - 1);

            if (fruitAtStart == 0) {
                basket.remove(fruitAtStart);
            }

            start++;
        }

        maxFruit = Math.max((end - start) + 1, maxFruit);
     }

    return maxFruit;
   }

}

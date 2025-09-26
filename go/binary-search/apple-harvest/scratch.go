package main

/* Binary search approach -------------------------
Ex: [3,6,7] 8

apples/hr:
1, 2, 3, 4, 5, 6, 7 (most is 7, larget tree)

left: 0;
right: len - 1;
mid: index 3 (val 4)

l         m        r
[1, 2, 3, 4, 5, 6, 7]
4app/hr -> 1, 2, 2: 5hrs VALID
eliminate all greater than index 3 (right mid + 1)
l         r
[1, 2, 3, 4, 5, 6, 7]

mid = index 1 (2)
 l  m     r
[1, 2, 3, 4, 5, 6, 7]
2app/hr -> 2,3,4: 9hrs
9hrs <= 8? FALSE -> NOT VALID
not valid -> left = mid + 1

       m
       l  r
[1, 2, 3, 4, 5, 6, 7]
3app/hr -> 1,2,3: 6hrs
6hrs <= 8? TRUE -> VALID
valid -> r = m
l == r ? true -> found it


if left === right return apples[left]
*/

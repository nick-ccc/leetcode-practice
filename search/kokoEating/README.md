# Median of Two Sorted Arrays

### Prompt:


```
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.
```

### Thinking:

If there is sufficient enough time the speed k should be equal to $\lceil(\sum piles[i] )/ n\rceil$. If the number of piles equals the number of hours it should be the maximum size of one of the piles. 

Using the hint of binary search, it should be brute force setting bounds between sum / hours and max pile and doing binary search between those values.

```
Example array:
30+11+23+4+20

1) h = 5, so k = 30
2) h = 6, so k = 23 (30,7 | 23 | 20 | 11 | 4) 
3) h = 7, so k = 20 (30,10 | 23,3 | 20 | 11 | 4)
4) h = 8, k != 11 (30,19,8 | 23,12,1 | 20,9 | 11 | 4) 
    -> k != 12 (30,18,6 | 23,11 | 20,8 | 11 | 4)
    -> k = 15  (30,15 | 23,8 | 20,5 | 11 | 4)
```



### Post Solution Analysis:



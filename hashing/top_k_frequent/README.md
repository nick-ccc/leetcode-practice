# Top K Frequent Elements

### Prompt:


```
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
```


### Thinking:

Array the length of K which will store the most frequent elements, and a hashmap which holds all entries. When we increment values in the hashmap we check the corresponding lowest entry in the array and swap if needed. 

After tying for a bit on attempt 0, I couldn't get my solution to work without basically needing to resort my top k elements each time, as to gaurantee the list is sorted when I try an insertion sort. This would go against the follow up hint of Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

What stood out when looking over past problem solutions, was that each element can only occur at most n times. This seems obvious but it should be possible to set map as count values not as key values. My thought from here 
was to then do some level of insertoin sort by collecting highest frequencies but again this would be wrong since if k = n, this would be O(nlogn). It seems best strategy is to then invery the map into int (frequency) map of lists (o(n)) then grab top k values.



### Post Solution Analysis:

N/A

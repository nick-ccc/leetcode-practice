# Contains Duplicate I.

### Prompt:


```
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
```

### Thinking:

Should be achievable in O(n), the strategy would be to store each value in the list in a hashset, and then if we come across a vlaue already in the hashset than we are done. 

### Post Solution Analysis:

Despite on paper algorithim 0 being faster from a complexity standpoint (O(n)), it was much slower than the sorting solution. This may be more of an underlying statistic about the data itself, where for large n where duplicates exist in small values, sorting first gives an advantage.

# Contains Duplicate III.

```
You are given an integer array nums and two integers indexDiff and valueDiff.

Find a pair of indices (i, j) such that:

    i != j,
    abs(i - j) <= indexDiff.
    abs(nums[i] - nums[j]) <= valueDiff, and
Return true if such pair exists or false otherwise.
```
 
 ### Thinking

Would like to point out, this has no duplicates involved...

First thought is we can sort the data and maintain the original indices, then go through each item for som i,j and if abs(i - j) > valueDiff then iterate to next i. This would be O(n^2).

First attempt was pretty bad, showing also lack of experience (currently) in go.
Hint suggest complexity should be O(nlogk). First new thought is that you only
care about number whose in range of index diff as far as looking ahead. Lets say you sort the nex k values, because indexDiff and valueDiff are less than if, for every subsquesnt match of index 0 of the k subset with index i implies that all value in between would be invalid. So a sliding window, that we sort. Still doesn't seem right,
it's just a slighlty better version of last solution.

Looking at more hints, it seems a binning based strategy would do well here, where
each value is droppped into bin based on value difference. This solution would be 
O(n).

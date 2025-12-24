# Find Minimum in Rotated Sorted Array

### Prompt:

```
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

    [4,5,6,7,0,1,2] if it was rotated 4 times.
    [0,1,2,4,5,6,7] if it was rotated 7 times.

Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.
```


### Thinking:

This has the same basis as searching in the in the rotated array. Here we can search for the smallest possible number as set in the restrictions (-5000) and just return whatever. 

Or better solution is take advantage of sorted subarrays, since left value of sorted subarray is the minimum of
that subarray.


### Post Solution Analysis:



# Search in Rotated Sorted Array

### Prompt:


```
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly left rotated at an unknown index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be left rotated by 3 indices and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.
```


### Thinking:

For log(n) complexity need to do a binary search. We need to, as we search, discern the relative placement. 
If left is more than right, we are not is a sorted subarray, i.e. the rotation is affecting the search. 
If `left<center` is in order than we can gaurantee that segment is in order. Equally if `center<right` 
we can gaurantee that segment is in order. 


```
Example cases:
[4,5,6,7,8,1,2]
 ^     ^     ^
 l     c     r

[5,6,0,1,2,3,4]
 ^     ^     ^
 l     c     r
```

Thus at any point we can say (at least) one  half of the array 
is sorted between either l->c or c->r. Binary search within a sorted subarray is the same, otherwise, we can 
adjust the other pointer location since. For instance

```
if left half is sorted:
    if (left value <= target and target <= center value): 
        right = center - 1 // target is in sorted half
    else: left = center + 1 // target is not, swap search to unsorted half
```

### Post Solution Analysis:



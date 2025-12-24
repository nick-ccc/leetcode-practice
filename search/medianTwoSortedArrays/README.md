# Median of Two Sorted Arrays

### Prompt:


```
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).
```

### Thinking:

Generally I recall this problem form undergrad. It reuqires keeping the count of two partitions of each array
such that it creates subarrays whose sum have equal (or 1 off) length. If we keep track of each left partition and each right parition we want their sums to be equal, and once the middle satisifes that the left and right of opposite subarrays from the two seperate original arrays satisy `arr1_left[end] < arr2_right[start]` and `arr2_left[end] < arr1_right[start]` then we are in the middle (since `arrx_left[end] < arrx_right[start]` is always true).

Scratch pad:
```
[1,2,3,9,10,11]
[5,6,7,12,13,14]

[1,2,3,5,6,7X,X9,10,11,12,13, 14]

In this case Median is 8 

[1,2,3] [9,10,11]
[5,6,7] [12,13,14]
```

```
[1,7,9]
[2,3,4,5,6,8,10]

[1,2,3,4,5,6,7,8,9,10]
          ^ median = 5.5 

[1] [7,9] 
[2,3,4,5] [6,8,10]
```

The next difficult part is figuring out how to move the
center index using a binary search scheme. 

```
[1,2,3,4,6]
[5,7,8,9]

[1,2,3,4,5,6,7,8,9]
         ^ median

[1,2,3] [4,6] 
[5,7] [8,9]

```



### Post Solution Analysis:



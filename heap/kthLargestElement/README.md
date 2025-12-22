# 

### Prompt:


```
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?
```


### Thinking:

First and easiet solution is to just sort and then do the kth index. 

Better solution is to use a minheap (priority queue) which should hold the smallest value as the first in the tree, i.e. the root. If we keep the size of the tree consistent to the length of K by always removing the smallest (root), we can gaurantee at the end that root would be smallest of the smallest k largest elements in the array.   

PSEUDO CODE 
```
kth_largest(arr,K): 
    minHeap = minHeap()
    for val in arr: 
        minHeap.heapPush(value)
        if len(minHeap) > K:
            minHeap.pop(0)
    return minHeap[0]
```

Alternative solution from alogrithims class would be finding the kth largest by stopping quick sort early once you can gaurantee the kth largest element has been found.


### Post Solution Analysis:



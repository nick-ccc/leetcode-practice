# Minimum Size Subarray Sum

### Prompt:


```
Given an array of positive integers nums and a positive integer target, return the minimal length of a whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.
```


### Thinking:

First want to address constaints, target is at leat 1, same for lenght and the actual values in the nums array. Meaning this isn't a two sum stule of problem where the can be formed form postive and negative. 

I can first think of simple solution where sort the list, and then iterate through to see if can find some combination that works based on difference between current and possible other elements less than the current (less than to avoid recounting).

### Post Solution Analysis:



# Two Sum

### Prompt:


```
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
```


### Thinking:

First thought is we can have a hasha map where the keys are all values less than the target. At each pass we can add an index to that value in the hash map and check what the corresponding the `Target - current` and if both have a match we are done. 

First edit, this won't work because numbers less than or equal to zero are allowed.
So we really need to add all numbers we see not just those less than the targets.


### Post Solution Analysis:

Original idea had to much fluff, solution beats 100% (so really tied) all solutions for runtime, but was less efficient at memory 61%. This memory efficiency is where my idea for only using values less than target before realizing negative numebrs were allowed, meaning target could be less than one of the included values. 
# Minimum Window Substring

### Prompt:


```
Given two strings s and t of lengths m and n respectively, return the minimum window

of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.
```


### Thinking:

This seems the same as minimum size subarray, except instead of keeping track of cumulative sum we need a cumulative hasmap to keep track if we have same results as what is in t.

This worked! But it's slow, runtimte is only about 11th percentile. My problem is that we are essentially everytime checking the two hash maps against each other based on the size t, if we deonte len(s)=n and len(t)=m this means we are O(m*n) so closer to n^2 than what I think is possible of O(m+n)

better solution is to first only build a map of the t, and keep count of how many keys are satisifed. If the amount of keys is satisifed is equal to number of keys in the map of t, we have a possivble solution. Otherwise, we don't. It's also key we don't keep track of other keys, and make sure not to double count the same character.

### Post Solution Analysis:

Really happy, solved hard first try albeit not fastest solutoin. Thought on some hints and was able to come up with O(m+n) solution.

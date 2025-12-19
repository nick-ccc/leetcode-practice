# Longest Substring Without Repeating Characters

### Prompt:


```
Given a string s, find the length of the longest substring without duplicate characters.
```


### Thinking:

This should be possible in O(n), basically we keep appending to our count, we can use a map to store what we have already seen. If we get duplicate start over, but keep saved value.

Update, while working through realized that if you, say dvdf above strategy wont work unless you subtract back, this is needlessy complicated instead opt for two pointers that slide
and define the window.


### Post Solution Analysis:

I saw really short solution I found quite elegant, taking advantage of array speed vs map (in go) and also very few lines.

```go
func lengthOfLongestSubstring(s string) int {
    ans := 0
    count := make([]int, 128)

    for l, r := 0, 0; r < len(s); r++{
        count[s[r]]++
        for count[s[r]] > 1{
            count[s[l]]--
            l++
        }
        if r - l + 1 > ans{
            ans = r - l + 1
        }
    }

    return ans
}
```


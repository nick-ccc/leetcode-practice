# Two Sum

### Prompt:


```
Given two strings s and t, return true if t is an of s, and false otherwise.
```


### Thinking:

First is we need to make sure strings are the same length, otherwise they cannot possibly be anagrams (assuming whitespace does not count).
First idea is we can keep count of each character and then when done
loop through dict to assert all keys have an even count. Still linear just requires two for loops O(2*n) -> O(n).

Realizing this won't exactly work since "aa" and "bb" would return true. To get around this we instead store. 


### Post Solution Analysis:

This solution seems slower than what should be possible according to percentile.
Looking at some hints a faster idea would be counting values in only one in a array of the entire alphabet. Then you can increment through with the latter, this would allow you to use an array that is indexable by the character index.

My takeaway is that both solution have linear complexity. However, it seems like the performance of writing to the list was just much fast than a map. I think this may be worth it's own research, a quick online reference found https://kokes.github.io/blog/2020/07/20/tiny-maps-arrays-go.html.  


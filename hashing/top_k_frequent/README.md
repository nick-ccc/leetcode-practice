# Top K Frequent Elements

### Prompt:


```
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
```


### Thinking:

Array the length of K which will store the most frequent elements, and a hashmap which holds all entries. When we increment values in the hashmap we check the corresponding lowest entry in the array and swap if needed. 

After tying for a bit on attempt 0, I couldn't get my solution to work without basically needing to resort my top k elements each time, as to gaurantee the list is sorted when I try an insertion sort. This would go against the follow up hint of Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

What stood out when looking over past problem solutions, was that each element can only occur at most n times. This seems obvious but it should be possible to set map as count values not as key values. My thought from here 
was to then do some level of insertoin sort by collecting highest frequencies but again this would be wrong since if k = n, this would be O(nlogn). It seems best strategy is to then invery the map into int (frequency) map of lists (o(n)) then grab top k values.


### Post Solution Analysis:

It's worth noting official solution suggests using Hoare's selection:
    - Count frequencies with a hashmap
    - Convert the map into an array of (num, freq) pairs
    - Use Quickselect to find the element with the k-th largest frequency
    - Return all elements with frequency ≥ that threshold

This does just seem like more complex answer than just iterating again through frequencies.


FROM AI: 
```
import "math/rand"

type Pair struct {
    num  int
    freq int
}

func quickSelect(a []Pair, left, right, k int) {
    for left < right {
        pivotIdx := partition(a, left, right)
        if pivotIdx == k {
            return
        } else if pivotIdx < k {
            left = pivotIdx + 1
        } else {
            right = pivotIdx - 1
        }
    }
}

func partition(a []Pair, left, right int) int {
    pivotIdx := left + rand.Intn(right-left+1)
    pivotFreq := a[pivotIdx].freq

    // Move pivot to end
    a[pivotIdx], a[right] = a[right], a[pivotIdx]

    store := left
    for i := left; i < right; i++ {
        if a[i].freq < pivotFreq {
            a[i], a[store] = a[store], a[i]
            store++
        }
    }

    // Move pivot to its final place
    a[store], a[right] = a[right], a[store]
    return store
}


func topKFrequent(nums []int, k int) []int {
    // 1) Count frequencies
    freqMap := make(map[int]int)
    for _, n := range nums {
        freqMap[n]++
    }

    // 2) Build array of (num, freq)
    pairs := make([]Pair, 0, len(freqMap))
    for n, f := range freqMap {
        pairs = append(pairs, Pair{n, f})
    }

    // We want k largest → index len(pairs) - k
    target := len(pairs) - k

    quickSelect(pairs, 0, len(pairs)-1, target)

    // 3) Collect result
    res := make([]int, 0, k)
    for i := target; i < len(pairs); i++ {
        res = append(res, pairs[i].num)
    }

    return res
}
```

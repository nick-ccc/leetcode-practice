# Minimum Size Subarray Sum

### Prompt:


```
Given an array of positive integers nums and a positive integer target, return the minimal length of a whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.
```


### Thinking:

First want to address constaints, target is at leat 1, same for lenght and the actual values in the nums array. Meaning this isn't a two sum stule of problem where the can be formed form postive and negative. 

I can first think of simple solution where sort the list, and then iterate through to see if can find some combination that works based on difference between current and possible other elements less than the current (less than to avoid recounting).

##### UPDATE

The probelm doesn't explicitly say consecutive subarray but based on test case, however by it's definition this is implied.

> A subarray is a portion of an array that consists of consecutive elements from the original array.
> 
```
[12,28,83,4,25,26,25,2,25,25,25]

func MinSubArrayLen(target int, nums []int) int {
	_minSubArrayLen := 0
	difference := target

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for _, val := range nums {
		difference = difference - val
		_minSubArrayLen++

		fmt.Printf("Val: %d, diff: %d, len: %d\n", val, difference, _minSubArrayLen)

		if difference <= 0 {
			return _minSubArrayLen
		}
	}
	return 0
}
```

Post update:


### Post Solution Analysis:



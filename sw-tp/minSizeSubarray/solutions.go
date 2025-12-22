package minSizeSubarray

import (
	"fmt"
	"sort"
)

// Attempt 0:
// Realized it need consecutive subarray although not explicitly stated
func MinSubArrayLen0(target int, nums []int) int {
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

// Solution 1
// Stats:
// - Runtime: 100.0% precentile
// - Memory: 32.53% percentile
func MinSubArrayLen1(target int, nums []int) int {
	left := 0
	minLength := len(nums) + 1
	currentLength := 0
	currentSum := 0

	for _, val := range nums {
		currentSum = currentSum + val
		currentLength++

		if currentSum >= target {
			minLength = min(
				minLength,
				currentLength,
			)

			// Now move left most value
			// to see if we missed out on possible
			// smaller array in process
			for {
				currentSum = currentSum - nums[left]
				currentLength--
				left++

				// Check again on current sum length
				if currentSum >= target {
					minLength = min(
						minLength,
						currentLength,
					)
				} else {
					break
				}
			}
		}
	}

	if minLength == len(nums)+1 {
		minLength = 0
	}

	return minLength
}

// Solution 2 - Goal reduce memory
// Stats:
// - Runtime: 100.0% precentile
// - Memory: 75.16% percentile
func MinSubArrayLen2(target int, nums []int) int {
	left := 0
	minLength := len(nums) + 1
	currentSum := 0

	for right, val := range nums {
		currentSum += val

		for currentSum >= target {
			minLength = min(minLength, right-left+1)
			currentSum -= nums[left]
			left++
		}
	}

	if minLength == len(nums)+1 {
		minLength = 0
	}

	return minLength
}

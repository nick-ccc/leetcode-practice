package minSizeSubarray

import "sort"

func minSubArrayLen(target int, nums []int) int {
	_minSubArrayLen := 0
	difference := target

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for _, val := range nums {
		difference = difference - val
		_minSubArrayLen++

		if difference <= 0 {
			return _minSubArrayLen
		}
	}
	return 0

}

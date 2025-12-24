package longestsubstring

// solution 0
// runtime: 100.00 percentile
// memory : 55.97 percentile
func search(nums []int, target int) int {
	left := 0
	right := len(nums)

	center := len(nums) / 2

	for left != right {
		if nums[center] == target {
			return center
		} else if nums[center] < target {
			// value is less than target slide window to right half
			left = center + 1
		} else {
			// value is more than target slide window to left half
			right = center
		}

		center = (right + left) / 2
	}

	return -1
}

// solution 0
// runtime: 100.00 percentile
// memory : 70.87 percentile
func search1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	center := len(nums) / 2
	for left <= right {
		if nums[center] == target {
			return center
		} else if nums[center] < target {
			// value is less than target slide window to right half
			left = center + 1
		} else {
			// value is more than target slide window to left half
			right = center - 1
		}

		center = (right + left) / 2
	}
	return -1
}

package searchRotateArray

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		center := left + (right-left)/2
		if nums[center] == target {
			return center
		} else if nums[left] <= nums[center] {
			// this branch handles cases where left is sorted
			if nums[left] <= target && target <= nums[center] {
				// target is in left (sorted) half
				right = center - 1
			} else {
				// target is in right half
				left = center + 1
			}
		} else {
			// this branch handles cases where right is sorted
			if nums[right] >= target && target >= nums[center] {
				left = center + 1
			} else {
				right = center - 1
			}
		}
	}
	return -1
}

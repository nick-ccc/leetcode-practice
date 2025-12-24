package searchRotateArray

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		// if left and right are sorted, return the minimum
		if nums[left] <= nums[right] {
			return nums[left]
		}
		center := left + (right-left)/2
		if nums[left] <= nums[center] {
			// if the left is less than or equal to center (i.e. case where
			// you have 2 values such that center rounds down to left)
			// move up center. In this case that would then mean that the
			// left is now equal to right and that should be the minimum
			left = center + 1
		} else {
			//right becomes center and moves up
			right = center
		}
	}

	// This is never ran could be anything
	// better way would be to exit loop to on condition of equality
	// and return nums[left]
	return -1
}

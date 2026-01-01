package medianSortedArrays

import "math"

// solution 0
// runtime: 100 percentile
// memory : 96.82 percentile
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	// Ensure nums1 is the smaller array
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	low, high := 0, m-1
	half := (m + n) / 2

	var maxLeftA, minRightA, maxLeftB, minRightB float64

	for {
		indexA := (low + high) / 2
		// (indexA + 1) + (indexB + 1) = half - hence half - indexA -2
		indexB := half - indexA - 2

		if indexA < 0 {
			maxLeftA = math.Inf(-1)
		} else {
			maxLeftA = float64(nums1[indexA])
		}

		if indexA+1 >= m {
			minRightA = math.Inf(1)
		} else {
			minRightA = float64(nums1[indexA+1])
		}

		if indexB < 0 {
			maxLeftB = math.Inf(-1)
		} else {
			maxLeftB = float64(nums2[indexB])
		}

		if indexB+1 >= n {
			minRightB = math.Inf(1)
		} else {
			minRightB = float64(nums2[indexB+1])
		}

		// Correct partition
		if maxLeftA <= minRightB && maxLeftB <= minRightA {
			// If even num of elements median is average of middle
			// values
			if (m+n)%2 == 0 {
				leftMax := math.Max(maxLeftA, maxLeftB)
				rightMin := math.Min(minRightA, minRightB)
				return (leftMax + rightMin) / 2.0
			}
			// else return min
			return math.Min(minRightA, minRightB)
		} else if maxLeftA > minRightB {
			// Too far right
			high = indexA - 1
		} else {
			// Too far left
			low = indexA + 1
		}
	}
}

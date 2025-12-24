package binarysearch

// solution 0
// runtime: X percentile
// memory : X percentile
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// first I want to make sure nums1 is always the bigger
	// array (or euqal)
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
}

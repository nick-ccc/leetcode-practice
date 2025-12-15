package containsduplicate

import (
	"sort"
)

// ----- Contains Duplicate ------ //

// Solution 0
// Stats:
// - Runtime: 57.99 precentile
// - Memory: 94.58 percentile
func containsDuplicateI0(nums []int) bool {
	// empty struct occupies 0 bytes in mem,
	// hence our map acts as a set
	numSet := make(map[int]struct{})

	for _, value := range nums {
		_, ok := numSet[value]
		if ok {
			return true
		} else {
			// struct{}{} is an empty struct
			// literal (0 bytes)
			numSet[value] = struct{}{}
		}
	}

	return false
}

// Solution 1
// Stats:
// - Runtime: 98.78 precentile
// - Memory: 88.90 percentile
func containsDuplicateI1(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// ----- Contains Duplicate III ------ //

// This is ugly and is showing my lack of experience in GoLang...
// Scrapping this, and rethinking...

// python style argsort
// func argsort(arr []int) []int {
// 	indices := make([]int, len(arr))
// 	for i := range indices {
// 		indices[i] = i
// 	}
// 	sort.Slice(indices, func(i, j int) bool { return arr[indices[i]] < arr[indices[j]] })
// 	return indices
// }

// func containsduplicateIII0(nums []int, indexDiff int, valueDiff int) bool {
// 	indexArr := argsort(nums)
// 	sort.Ints(nums)
// 	var indexDiffCurrent int
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := 0; j < len(nums)-1; j++ {

// 			indexDiffCurrent = indexArr[i] - indexArr[j]
// 			if indexDiffCurrent < 0 {
// 				indexDiffCurrent = -1 * indexDiffCurrent
// 			}
// 			if indexDiffCurrent <= indexDiff {
// 				indexDiffCurrent = indexArr[i] - indexArr[j]
// 				if indexDiffCurrent < 0 {
// 					indexDiffCurrent = -1 * indexDiffCurrent
// 				}
// 			}
// 		}
// 	}

// }

// used hints and dicussions - seems like a revolving binning strategy
// that keeps updating the state is the best solution here
func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// Solution 1
// Stats:
// - Runtime: 61.11 precentile
// - Memory: 80.56 percentile
func containsduplicateIII1(nums []int, indexDiff int, valueDiff int) bool {

	//size := len(nums)

	bucket := make(map[int]int)

	// width of each bucket
	width := valueDiff + 1

	for idx, number := range nums {
		var bucketIdx int

		if number >= 0 {
			bucketIdx = number / width
		} else {
			// offset patch for negative number
			bucketIdx = (number / width) - 1
		}

		if _, exist := bucket[bucketIdx]; exist {
			// two numbers in the same bucket, gap must be smaller than width
			return true

		} else if _, exist := bucket[bucketIdx+1]; exist && (Abs(number-bucket[bucketIdx+1]) < width) {
			// two number in two consecutive buckets, and gap is smaller than width
			return true

		} else if _, exist := bucket[bucketIdx-1]; exist && (Abs(number-bucket[bucketIdx-1]) < width) {
			// two number in two consecutive buckets, and gap is smaller than width
			return true
		}

		// put current number into corresponding bucket
		bucket[bucketIdx] = number

		if idx >= indexDiff {
			delete(bucket, (nums[idx-indexDiff] / width))
		}

	}

	return false
}

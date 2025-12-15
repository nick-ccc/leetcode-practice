package twosum

// Solution 0
// Stats:
// - Runtime: 100.00 precentile
// - Memory: 63.51 percentile
func twoSum0(nums []int, target int) []int {
	bucket := make(map[int]int)

	for idx, number := range nums {
		// Check this first, because if [3,3] is nums, we need to return 0,1
		if compliment, exist := bucket[target-number]; exist && (compliment != idx) {
			return []int{idx, compliment}
		}

		bucket[number] = idx
	}
	return nil
}

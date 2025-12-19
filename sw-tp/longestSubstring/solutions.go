package longestsubstring

// Solution 0
// Stats:
// - Runtime: 84.3% precentile
// - Memory: 26.03% percentile
func lengthOfLongestSubstring(s string) int {
	longest_substring_len := 0

	if len(s) > 0 {
		letters := []rune(s)
		substring_map := make(map[rune]int)
		left_most_slider_idx := 0

		for idx, val := range letters {
			if val_idx, exist := substring_map[val]; exist && val_idx >= left_most_slider_idx {
				// First check if this is an old letter not
				// in the current sliding window
				longest_substring_len = max(
					idx-left_most_slider_idx,
					longest_substring_len,
				)
				left_most_slider_idx = val_idx + 1
			}
			// Update the map and iterate right most slider index
			substring_map[val] = idx
		}

		longest_substring_len = max(
			len(letters)-left_most_slider_idx,
			longest_substring_len,
		)
	}

	return longest_substring_len
}

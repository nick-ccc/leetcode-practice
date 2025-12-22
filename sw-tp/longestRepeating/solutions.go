package longestsubstring

// Solution 0
// Runtime: 72.86 Percentile
// Memory: 8.54 Precntile
func characterReplacement0(s string, k int) int {

	sRunes := []rune(s)
	letterMap := make(map[rune]int)
	left := 0

	highestFrequency := 0

	for right, val := range sRunes {
		letterMap[val]++
		highestFrequency = max(
			highestFrequency,
			letterMap[val],
		)
		// If size of current window, has a high count
		// of some variable such the difference is more
		// than the tolerance, than our window is too
		// big, and we need to readjust
		if (right - left + 1 - highestFrequency) > k {
			letterMap[sRunes[left]]--
			left++
		}
	}

	if highestFrequency+k > len(s) {
		return len(s)
	} else {
		return highestFrequency + k
	}
}

// Taking advantage of GO bytes and limit
// of characters being capital letters
// Runtime: 100 Percentile
// Memory: 95.10 Precntile
func characterReplacement1(s string, k int) int {
	var charArray [26]int
	left := 0
	highestFrequency := 0

	for right := 0; right < len(s); right++ {
		i := int(s[right] - 'A')
		charArray[i]++

		highestFrequency = max(
			highestFrequency,
			charArray[i],
		)

		if (right - left + 1 - highestFrequency) > k {
			charArray[s[left]-'A']--
			left++
		}
	}

	if highestFrequency+k > len(s) {
		return len(s)
	} else {
		return highestFrequency + k
	}
}

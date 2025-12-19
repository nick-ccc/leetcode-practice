package minimumwindowsubstring

func generateMap(s string) map[rune]int {
	letters := []rune(s)
	letter_map := make(map[rune]int)

	for _, letter := range letters {
		letter_map[letter]++
	}
	return letter_map
}

func compare_letter_maps(
	benchmark map[rune]int,
	comparison map[rune]int,
) bool {
	// Assume that the benchmark must be smaller than or
	// equal to size of comparison. I.e. if our substring has
	// has less keys than the word t, we can't have a subset
	// contianing all letters in t
	if len(benchmark) > len(comparison) {
		return false
	}

	for key, value := range benchmark {
		if value_comp, exist := comparison[key]; !exist || value > value_comp {
			return false
		}
	}

	return true
}

// Solution 0
// Stats:
// - Runtime: 11.4% precentile
// - Memory: 11.83% percentile
func MinWindow0(s string, t string) string {
	substring := ""

	if len(s) >= len(t) {
		min_length := len(s) + 1
		letters := []rune(s)
		left := 0

		benchmark := generateMap(t)
		comparison := make(map[rune]int)

		for right, val := range letters {
			comparison[val]++

			for compare_letter_maps(benchmark, comparison) {
				if right+1-left < min_length {
					substring = s[left : right+1]
					min_length = right + 1 - left
				}
				comparison[letters[left]]--
				left++
			}
		}
	}

	return substring
}

// Solution 1
// Stats:
// - Runtime: 100.0% precentile
// - Memory: 97.86% percentile
func MinWindow1(s string, t string) string {
	substring := ""

	if len(s) >= len(t) {
		min_length := len(s) + 1

		letters_s := []rune(s)
		left := 0
		letters_satisfied := 0

		// bench mark hashmap
		letters_benchmark := []rune(t)
		benchmark := make(map[rune]int)
		for _, letter := range letters_benchmark {
			benchmark[letter]++
		}

		len_benchmark := len(benchmark)

		for right, val := range letters_s {
			if _, exist := benchmark[val]; exist {
				benchmark[val]--

				// only satisifeds when crossing 0 threshold otherwise
				// we double dip for same letter in keeping track
				if benchmark[val] == 0 {
					letters_satisfied++
				}
			}

			for letters_satisfied == len_benchmark {
				// fmt.Println(benchmark)
				// fmt.Println(s[left : right+1])
				// Update if the new min length is less than old
				if right+1-left < min_length {
					substring = s[left : right+1]
					min_length = right + 1 - left
				}

				// If letter exists in benchmark increment
				// and check total amount is not greater than 0
				if _, exist := benchmark[letters_s[left]]; exist {
					benchmark[letters_s[left]]++
					if benchmark[letters_s[left]] > 0 {
						letters_satisfied--
					}
				}
				left++
			}
		}
	}

	return substring
}

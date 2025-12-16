package validanagram

type Count struct {
	s int
	t int
}

// Solution 0
// Stats:
// - Runtime: 5.00 precentile
// - Memory: 77.00 percentile
func isAnagram0(s string, t string) bool {
	length_s := len(s)
	if length_s == len(t) {
		hashmap := make(map[byte]Count)
		for i := 0; i < length_s; i++ {
			c := hashmap[s[i]]
			c.s++
			hashmap[s[i]] = c

			c = hashmap[t[i]]
			c.t++
			hashmap[t[i]] = c
		}

		for k := range hashmap {
			if hashmap[k].s != hashmap[k].t {
				return false
			}
		}
		return true

	}
	return false
}

// Solution “
// Stats:
// - Runtime: 100.00 precentile
// - Memory: 73.00 percentile
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	alph_arr := [26]int{}

	for _, ch := range s {
		alph_arr[ch-'a'] += 1
	}
	for _, ch := range t {
		alph_arr[ch-'a'] -= 1
		if alph_arr[ch-'a'] < 0 {
			return false
		}
	}
	return true
}

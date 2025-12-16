package groupanagrams

import (
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams0(strs []string) [][]string {
	sorting_map := make(map[string][]string)

	var sorted_str string
	for _, str := range strs {
		sorted_str = SortString(str)
		arr := sorting_map[sorted_str]
		newArr := append(arr, str)
		sorting_map[sorted_str] = newArr
	}

	values := make([][]string, 0, len(sorting_map))
	for _, v := range sorting_map {
		values = append(values, v)
	}

	return values
}

package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	freq := make([][]int, len(nums)+1)

	for _, n := range nums {
		count[n]++
	}

	for n, _count := range count {
		freq[_count] = append(freq[_count], n)
	}

	res := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		for _, n := range freq[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}

func main() {
	// fmt.Println(topKFrequent([]int{1, 2, 1, 4, 1, 2, 3, 1, 3, 2}, 4))
	fmt.Println(topKFrequent([]int{2, 1, 1}, 2))
}

package longestsubstring

import "container/heap"

// An IntHeap implements heap.Interface and holds ints.
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// The Less method defines the ordering. For a min-heap, we use <.
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push and Pop use pointer receivers because
// they modify the slice's length, not just its contents.
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func findKthLargest(nums []int, k int) int {
	minHeap := IntHeap{}
	heap.Init(&minHeap)

	for _, val := range nums {
		heap.Push(&minHeap, val)

		// pop if len > k
		if minHeap.Len() > k {
			heap.Pop(&minHeap)
		}
	}

	return minHeap[0]
}

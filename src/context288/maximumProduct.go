package context288

import (
	"container/heap"
	"sort"
)

func maximumProduct(nums []int, k int) (ans int) {
	h := &hp{nums}

	heap.Init(h)

	// O(k*log(n))
	for k > 0 {
		min := h.pop()
		h.push(min + 1)
		k--
	}

	ans = 1
	// nums[i] >= 0
	for i := range nums {
		ans = (ans * nums[i]) % 1000000007
	}

	return
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func (h *hp) push(x int) {
	heap.Push(h, x)
}

func (h *hp) pop() int {
	return heap.Pop(h).(int)
}

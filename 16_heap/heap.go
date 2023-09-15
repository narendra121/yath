package heap

import "math"

/*

left(i)= 2i+1
right(i)= 2i+2
parent(i)= floor (i-1/2)

Min Heap:-
	Complete binary tree
	Every node  has value smaller than its descendant
*/

type Heap struct {
	arr  []int
	size int
	cap  int
}

func (h *Heap) Init(cap int) {
	h.arr = make([]int, 0, cap)
	h.cap = cap
	h.size = 0
}

func (h *Heap) GetLeftIdx(idx int) int {
	return (2*idx + 1)
}

func (h *Heap) GetRightIdx(idx int) int {
	return (2*idx + 2)
}

func (h *Heap) GetParentIdx(idx int) int {
	return (idx - 1) / 2
}

func (h *Heap) Insert(x int) {
	if h.size == h.cap {
		return
	}
	h.size++
	h.arr[h.size-1] = x
	i := h.size - 1
	for i != 0 && h.arr[h.GetParentIdx(i)] > h.arr[i] {
		h.arr[i], h.arr[h.GetParentIdx(i)] = h.arr[h.GetParentIdx(i)], h.arr[i]
		i = h.GetParentIdx(i)
	}
}

// heapify specific node
func (h *Heap) MinHeapify(idx int) {
	lt, rt := h.GetLeftIdx(idx), h.GetRightIdx(idx)
	small := idx
	if lt < h.size && h.arr[lt] < h.arr[idx] {
		small = lt
	}
	if rt < h.size && h.arr[rt] < h.arr[small] {
		small = rt
	}
	if small != idx {
		h.arr[idx], h.arr[small] = h.arr[small], h.arr[idx]
		h.MinHeapify(small)
	}
}

func (h *Heap) ExtractMin() int {
	if h.size == 0 {
		return math.MaxInt
	}
	if h.size == 1 {
		h.size--
		return h.arr[0]
	}
	h.arr[0], h.arr[h.size-1] = h.arr[h.size-1], h.arr[0]
	h.size--
	h.MinHeapify(0)
	return h.arr[h.size]
}

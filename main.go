package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type sized struct{ size, end int }

type sizeHeap []sized

func (h sizeHeap) Len() int           { return len(h) }
func (h sizeHeap) Less(i, j int) bool { return h[i].size < h[j].size }
func (h sizeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *sizeHeap) Push(x any)        { *h = append(*h, x.(sized)) }
func (h *sizeHeap) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

// MinIntervalPerQuery returns, for each query point q, the size of the
// smallest interval [l, r] containing it (size r-l+1), or -1. Offline sweep:
// queries ascending, intervals admitted by start, expired by end, smallest
// surviving size on top of a min-heap.
func MinIntervalPerQuery(intervals [][2]int, queries []int) []int {
	ivs := append([][2]int{}, intervals...)
	sort.Slice(ivs, func(i, j int) bool { return ivs[i][0] < ivs[j][0] })
	order := make([]int, len(queries))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(a, b int) bool { return queries[order[a]] < queries[order[b]] })

	out := make([]int, len(queries))
	h := &sizeHeap{}
	next := 0
	for _, qi := range order {
		q := queries[qi]
		for next < len(ivs) && ivs[next][0] <= q {
			h.PushSized(ivs[next])
			next++
		}
		for h.Len() > 0 && (*h)[0].end < q {
			heap.Pop(h) // ended before this (and every later) query
		}
		if h.Len() == 0 {
			out[qi] = -1
		} else {
			out[qi] = (*h)[0].size
		}
	}
	return out
}

// PushSized pushes an interval as (size, end).
func (h *sizeHeap) PushSized(iv [2]int) {
	heap.Push(h, sized{iv[1] - iv[0] + 1, iv[1]})
}

func main() {
	fmt.Println(MinIntervalPerQuery([][2]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5}))
}

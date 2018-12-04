package heapsort

import (
	"github.com/paulbuis/sorts/types"
)

// siftDown implements the heap property on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDown(data types.SliceType, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data[child] < data[child+1] {
			child++
		}
		if !(data[root] < data[child]) {
			return
		}
		data[root], data[child] = data[child], data[root]
		root = child
	}
}

func HeapSort(data types.SliceType, a, b int) {
	hi := len(data)

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		data[0], data[i] = data[i], data[0]
		siftDown(data, 0, i)
	}
}
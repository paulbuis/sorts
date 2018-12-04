package quicksort

import (
	"github.com/paulbuis/sorts/types"
)

// medianOfThree moves the median of the three values data[m0], data[m1], data[m2] into data[m1].
func medianOfThree(slice types.SliceType, m1, m0, m2 int) {
	// sort 3 elements
	if slice.Less(m1, m0) {
		slice.Swap(m1, m0)
	}
	// data[m0] <= data[m1]
	if slice.Less(m2, m1) {
		slice.Swap(m2, m1)
		// data[m0] <= data[m2] && data[m1] < data[m2]
		if slice.Less(m1, m0) {
			slice.Swap(m1, m0)
		}
	}
	// now data[m0] <= data[m1] <= data[m2]
}

type PivotChooserFunc func(slice types.SliceType) int

// puts pivot value in slice[0]
func choosePivot(slice types.SliceType) int {
	n := len(slice) - 1
	m := n/2
	if n > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := n / 8
		medianOfThree(slice, 0, s, 2*s)
		medianOfThree(slice, m, m-s, m+s)
		medianOfThree(slice, n-1, n-1-s, n-1-2*s)
	}
	medianOfThree(slice, 0, m, n-1)
	return m
}
package quicksort

import (
	"github.com/paulbuis/sorts/types"
)

// Quicksort, loosely following Bentley and McIlroy,
// ``Engineering a Sort Function,'' SP&E November 1993.


func doPivot(slice types.SliceType, pivotChooser PivotChooseFunc) (int, int) {
	// Invariants are:
	//	data[lo] = pivot (set up by ChoosePivot)
	//	data[lo < i < a] < pivot
	//	data[a <= i < b] <= pivot
	//	data[b <= i < c] unexamined
	//	data[c <= i < hi-1] > pivot
	//	data[hi-1] >= pivot
	m := pivotChooser(slice)
	pivot := 0
	hi := len(slice)-1
	a, c := 1, hi-1

	for ; a < c && slice.Less(a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !slice.Less(pivot, b); b++ { // slice[b] <= pivot
		}
		for ; b < c && slice.Less(pivot, c-1); c-- { // slice[c-1] > pivot
		}
		if b >= c {
			break
		}
		// slice[b] > pivot; slice[c-1] <= pivot
		slice.Swap(b, c-1)
		b++
		c--
	}
	// If hi-c<3 then there are duplicates (by property of median of nine).
	// Let be a bit more conservative, and set border to 5.
	protect := hi-c < 5
	if !protect && hi-c < hi/4 {
		// Lets test some points for equality to pivot
		dups := 0
		if !slice.Less(pivot, hi-1) { // data[hi-1] = pivot
			slice.Swap(c, hi-1)
			c++
			dups++
		}
		if !slice.Less(b-1, pivot) { // data[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> data[m] <= pivot
		if !slice.Less(m, pivot) { // data[m] = pivot
			slice.Swap(m, b-1)
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed distribution
		protect = dups > 1
	}
	if protect {
		// Protect against a lot of duplicates
		// Add invariant:
		//	data[a <= i < b] unexamined
		//	data[b <= i < c] = pivot
		for {
			for ; a < b && !slice.Less(b-1, pivot); b-- { // data[b] == pivot
			}
			for ; a < b && slice.Less(a, pivot); a++ { // data[a] < pivot
			}
			if a >= b {
				break
			}
			// data[a] == pivot; data[b-1] < pivot
			slice.Swap(a, b-1)
			a++
			b--
		}
	}
	// Swap pivot into middle
	slice.Swap(pivot, b-1)
	return b - 1, c
}
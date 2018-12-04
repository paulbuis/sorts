package quicksort

import (
	"github.com/paulbuis/sorts/types"
	"github.com/paulbuis/sorts/insertionsort"
)

type PivotChooseFunc func(slice types.SliceType) int
type PartitionFunc func(slice types.SliceType, choosePivot PivotChooseFunc) (int, int)

func QuickSort(pivotChooser PivotChooseFunc, partitioner PartitionFunc) types.RecursiveSortFunc {
	var quickSort types.RecursiveSortFunc
    	quickSort = func(slice types.SliceType, maxDepth types.Depth) {
		a := 0
		b := len(slice) - 1
	
		// loop is used as tail recursion elimination technique
		for b > 12 { // Use ShellSort for slices <= 12 elements
			// this code needed to guarantee O(NlogN) performance, omit for now, but fix later
			//if maxDepth == 0 {
			//	heapSort(data, a, b)
			//	return
			//}
			maxDepth--
			mlo, mhi := partitioner(slice, pivotChooser)
		    // Avoiding recursion on the larger subproblem guarantees
		    // a stack depth of at most lg(b-a).
		    if mlo < b-mhi {
			    quickSort(slice[a: mlo], maxDepth)
			    a = mhi // i.e., quickSort(data, mhi, b)
		    } else {
			    quickSort(slice[mhi: b], maxDepth)
			    b = mlo // i.e., quickSort(data, a, mlo)
		    }
	    }
	    if b > 1 {
		    // Do ShellSort pass with gap 6
		    // It could be written in this simplified form cause b-a <= 12
		    for i := a + 6; i < b; i++ {
			    if slice.Less(i, i-6) {
				    slice.Swap(i, i-6)
			    }
		    }
		    insertionsort.Sort(slice)
		}
	}
	return quickSort
}

// Sort sorts data.

// Expected O(n*log(n)) calls to Less and Swap.
// Worst case O(n*log(n)) if maxDepth used to switch
// to a different in-place O(n*log(n)) method when approaching
// the degenerate case. Using maxDepth in this way is 
// called the "IntroSort" variant of QuickSort
// The sort is not guaranteed to be stable.
func Sort(slice types.SliceType) {
	n := len(slice)
	quickSort := QuickSort(choosePivot, doPivot)
	quickSort(slice, maxDepth(n))
}

// maxDepth returns a threshold at which quicksort should switch
// to heapsort. It returns 2*ceil(lg(n+1)).
func maxDepth(n int) types.Depth {
	var depth types.Depth
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}
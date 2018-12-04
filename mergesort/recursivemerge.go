package mergesort

import (
	"sort" // standard sort package, provides SearchInts to do a binary search on a []int
	"github.com/paulbuis/sorts/types"
)

func RecursiveMerge(A types.SliceType, B types.SliceType, C types.SliceType) {
	lenA := len(A)
	lenB := len(B)
	if lenA < lenB {
		A, B = B, A
		lenA, lenB = lenB, lenA
	}

	if lenA <= 10000 { // better threashold?, note that 0 does not work!
		IterativeMerge(A,B,C) // non-recursive
		return
	}

	midAindex := lenA / 2
	searchTarget := A[midAindex]
	midBindex := sort.Search(len(B), func(i int) bool { return !types.Less(B[i],searchTarget) })
	midCindex := midAindex + midBindex

	C[midCindex] = searchTarget
	RecursiveMerge(A[:midAindex], B[:midBindex], C[:midCindex])
	RecursiveMerge(A[midAindex+1:], B[midBindex:], C[midCindex+1:])
}
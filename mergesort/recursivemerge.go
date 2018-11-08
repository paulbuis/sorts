package mergesort

import (
	"sort" // standard sort package, provides SearchInts to do a binary search on a []int
)

func RecursiveMerge(A []int, B []int, C []int) {
	lenA := len(A)
	lenB := len(B)
	if lenA < lenB {
		A, B = B, A
		lenA, lenB = lenB, lenA
	}

	if lenA <= 1000 { // better threashold?, note that 0 does not work!
		IterativeMerge(A,B,C) // non-recursive
		return
	}

	midAindex := lenA / 2
	midBindex := sort.SearchInts(B, A[midAindex])
	midCindex := midAindex + midBindex

	C[midCindex] = A[midAindex]
	RecursiveMerge(A[:midAindex], B[:midBindex], C[:midCindex])
	RecursiveMerge(A[midAindex+1:], B[midBindex:], C[midCindex+1:])
}
package parallelmergesort

import (
	"github.com/paulbuis/sorts/mergesort"
	"sort"
	// also, use one of the following:
	// "sync" for sync.Mutex
	"github.com/paulbuis/sorts/parallel" //for parallel.Done or parallel.Spawn
	
)

// this is just a copy-paste of mergesort.RecursiveMerge
// with a rename of RecursiveMerge to ParallelMerge
// and the insertion of the Cormen:spawn amd Corment:sync comments near end
// See writeups of Cormen parallel merge algorithm in
// Cormen Algorithms textbook, 3rd edition, chapter 27, or
// https://www2.hawaii.edu/~janst/311/Notes/Topic-22.html
// https://en.wikipedia.org/wiki/Merge_algorithm#Parallel_merge

func ParallelMerge(A []int, B []int, C []int) {
	lenA := len(A)
	lenB := len(B)
	if lenA < lenB {
		A, B = B, A
		lenA, lenB = lenB, lenA
	}

	if lenA <= 100000 { // better threashold?, note that 0 does not work!
		mergesort.RecursiveMerge(A,B,C) // or should it be IterativeMerge ???
		return
	}

	midAindex := lenA / 2
	midBindex := sort.SearchInts(B, A[midAindex])
	midCindex := midAindex + midBindex

	C[midCindex] = A[midAindex]
	lock := parallel.Spawn(func() {
		ParallelMerge(A[:midAindex], B[:midBindex], C[:midCindex])
	})
	ParallelMerge(A[midAindex+1:], B[midBindex:], C[midCindex+1:])
	lock.Lock()
}
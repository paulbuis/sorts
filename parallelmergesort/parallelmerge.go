package parallelmergesort

import (
	"github.com/paulbuis/sorts/mergesort"
	"sort"
	// also, use one of the following:
	// "sync" for sync.Mutex
	"github.com/paulbuis/sorts/parallel" //for parallel.Done or parallel.Spawn or parallel.TaskGroup
	"github.com/paulbuis/sorts/types"
)

// this started as just a copy-paste of mergesort.RecursiveMerge
// with a rename of RecursiveMerge to ParallelMerge
// and the insertion of the Cormen:spawn amd Corment:sync comments near end
// See writeups of Cormen parallel merge algorithm in
// Cormen Algorithms textbook, 3rd edition, chapter 27, or
// https://www2.hawaii.edu/~janst/311/Notes/Topic-22.html
// https://en.wikipedia.org/wiki/Merge_algorithm#Parallel_merge

func ParallelMerge(A, B, C types.SliceType) {
	tg := parallel.MakeTaskGroup()
	parallelMerge(A, B, C, tg)
	tg.Wait()
}

func parallelMerge(A, B, C types.SliceType, tg parallel.TaskGroup) {
	lenA := len(A)
	lenB := len(B)
	if lenA < lenB {
		A, B = B, A
		lenA, lenB = lenB, lenA
	}

	if lenA <= 500000 { // better threashold?, note that 0 does not work!
		mergesort.RecursiveMerge(A,B,C) // or should it be IterativeMerge ???
		return
	}

	midAindex := lenA / 2
	searchTarget := A[midAindex]
	midBindex := sort.Search(len(B), func(i int) bool { return B[i] >= searchTarget })
	midCindex := midAindex + midBindex

	C[midCindex] = A[midAindex]
	tg.Add(func() { parallelMerge(A[:midAindex], B[:midBindex], C[:midCindex], tg)})
	parallelMerge(A[midAindex+1:], B[midBindex:], C[midCindex+1:], tg)
	/*
	lock := parallel.Spawn(func() {
		ParallelMerge(A[:midAindex], B[:midBindex], C[:midCindex])
	})
	ParallelMerge(A[midAindex+1:], B[midBindex:], C[midCindex+1:])
	lock.Lock()
	*/
}
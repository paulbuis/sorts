package main

import (
//	"github.com/paulbuis/sorts/insertionsort"
//	"github.com/paulbuis/sorts/mergesort"
//	"github.com/paulbuis/sorts/parallelmergesort"'
	"github.com/paulbuis/sorts/dualpivotquicksort"
	"github.com/paulbuis/sorts/types"
	"sort"
)

/*
func parallelWithIterativeMerge(slice types.SliceType) {
	seqMergeSorter := mergesort.MergeSorter{InPlaceSort: insertionsort.Sort, Merge: mergesort.IterativeMerge}
	pMerge := parallelmergesort.Merge(seqMergeSorter.Merge)
	pms := parallelmergesort.NewParallelMergeSorter(seqMergeSorter, pMerge)
	pms.Sort(slice)
}

func parallelWithRecursiveMerge(slice types.SliceType) {
	seqMergeSorter := mergesort.MergeSorter{InPlaceSort: insertionsort.Sort, Merge: mergesort.RecursiveMerge, MinSize:100}
	pMerge := parallelmergesort.Merge(seqMergeSorter.Merge)
	pms := parallelmergesort.NewParallelMergeSorter(seqMergeSorter, pMerge)
	pms.Sort(slice)
}

func parallelWithParallelMerge(slice types.SliceType) {
	seqMergeSorter := mergesort.MergeSorter{InPlaceSort: insertionsort.Sort, Merge: mergesort.RecursiveMerge, MinSize:100}
	pMerge := parallelmergesort.Merge(parallelmergesort.ParallelMerge)
	pms := parallelmergesort.NewParallelMergeSorter(seqMergeSorter, pMerge)
	pms.Sort(slice)
}

func mergeSortIterativeMerge(slice types.SliceType) {
	ms := mergesort.MergeSorter{InPlaceSort: insertionsort.Sort, Merge: mergesort.IterativeMerge, MinSize:100 }
	ms.Sort(slice)
}

func mergeSortRecursiveMerge(slice types.SliceType) {
	ms := mergesort.MergeSorter{InPlaceSort: insertionsort.Sort, Merge: mergesort.RecursiveMerge, MinSize:100}
	ms.Sort(slice)
}
*/

func quickSort(slice types.SliceType) {
	sort.Sort(slice)
}

func dualPivotQuicksort(slice types.SliceType) {
	dualpivotquicksort.Sort(slice)
}

func stableSort(slice types.SliceType) {
	sort.Stable(slice)
}
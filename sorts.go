package main

import (
	"github.com/paulbuis/sorts/insertionsort"
	"github.com/paulbuis/sorts/mergesort"
	"github.com/paulbuis/sorts/parallelmergesort"
	"sort"
)
func parallelWithIterativeMerge(slice []int) {
	seqMergeSorter := mergesort.MergeSorter{InPlaceSort: insertionsort.Sort, Merge: mergesort.IterativeMerge}
	pMerge := parallelmergesort.Merge(seqMergeSorter.Merge)
	pms := parallelmergesort.NewParallelMergeSorter(seqMergeSorter, pMerge)
	pms.Sort(slice)
}

func parallelWithRecursiveMerge(slice []int) {
	seqMergeSorter := mergesort.MergeSorter{InPlaceSort: insertionsort.Sort, Merge: mergesort.RecursiveMerge}
	pMerge := parallelmergesort.Merge(seqMergeSorter.Merge)
	pms := parallelmergesort.NewParallelMergeSorter(seqMergeSorter, pMerge)
	pms.Sort(slice)
}

func parallelWithParallelMerge(slice []int) {
	seqMergeSorter := mergesort.MergeSorter{InPlaceSort: insertionsort.Sort, Merge: mergesort.RecursiveMerge}
	pMerge := parallelmergesort.Merge(parallelmergesort.ParallelMerge)
	pms := parallelmergesort.NewParallelMergeSorter(seqMergeSorter, pMerge)
	pms.Sort(slice)
}

func mergeSortIterativeMerge(slice []int) {
	ms := mergesort.MergeSorter{InPlaceSort: insertionsort.Sort, Merge: mergesort.IterativeMerge}
	ms.Sort(slice)
}

func mergeSortRecursiveMerge(slice []int) {
	ms := mergesort.MergeSorter{InPlaceSort: insertionsort.Sort, Merge: mergesort.RecursiveMerge}
	ms.Sort(slice)
}

func quickSort(slice []int) {
	sort.Ints(slice)
}

func stableSort(slice []int) {
	sort.Stable(sort.IntSlice(slice))
}
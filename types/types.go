package types

import (
//	"errors"
	"math/rand"
	"sort"
)

type ElementType int32
type SliceType []int32

type ReSlice struct {
	Slice SliceType
	LoIndex int
	HiIndex int
}

func ReSliceSlice(slice SliceType, loIndex, hiIndex int) (ReSlice, error) {
	//if loIndex<0 || hiIndex >= len(slice) {
    //
	//}
	return ReSlice{Slice: slice, LoIndex: loIndex, HiIndex: hiIndex}, nil
}

func ReSliceReSlice(reSlice ReSlice, loIndex, hiIndex int) (ReSlice, error) {
	// if loIndex<reSlice.LoIndex || hiIndex > resSlice.hiIndex {
    //
	//}
	return ReSlice{Slice: reSlice.Slice, LoIndex: loIndex, HiIndex: hiIndex}, nil
}


type InPlaceSortFunc func (slice SliceType)
type Depth uint8
type RecursiveSortFunc func (slice SliceType, depth Depth)
type MergeFunc func (in1 SliceType, in2 SliceType, out SliceType)
type MergeSortFunc func (slice SliceType, tmp SliceType, depth Depth)
type NonMutatingSortFunc func (slice SliceType) SliceType

// Len, Swap, and Less methods conform to sort.Interface
func (slice SliceType) Len() int {
	return len(slice)
}

func (slice SliceType) Swap(i,j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice SliceType) Rotate3Left(i,j,k int) {
	slice[j], slice[k], slice[i] = slice[i], slice[j], slice[k]
}

func (slice SliceType) Rotate3Right(i,j,k int) {
	slice[k], slice[i], slice[j] = slice[i], slice[j], slice[k]
}




// use only to compare across different slices
func Less(a, b int32) bool {
	return a < b
}

// should be inlined if invoked on something of type types.SliceType
// but not inlined if invoked via type sort.Interface or types.Interface!!
func (slice SliceType) Less (i,j int) bool {
	return slice[i] < slice[j]
}

func (slice SliceType) Equal(i, j int) bool {
	return slice[i] == slice[j]
}

type Interface interface {
	sort.Interface
	Equal(i,j int) bool
}

// ReSlice also conforms to sort.Interface, but ...
// if these methods invoked direclty rather than via that interface
// we can reasonably expect them to be inlined
func (reSlice ReSlice) Len() int {
	return reSlice.HiIndex - reSlice.LoIndex + 1
}

func (reSlice ReSlice) Less(i, j int) bool {
	return reSlice.Slice[i] < reSlice.Slice[j]
}

func (reSlice ReSlice) Swap(i, j int) {
	reSlice.Slice[i], reSlice.Slice[j] = reSlice.Slice[j], reSlice.Slice[i]
}

func Random(n int) (slice SliceType) {
	slice = make(SliceType, n, n)
	for i := 0; i<n; i += 1 {
		slice[i] = rand.Int31n(int32(n))
	}
	return slice
}

func RecursiveSortFunc2InPlaceSortFunc(rsf RecursiveSortFunc) InPlaceSortFunc {
	return func (slice SliceType) {
		rsf(slice, 0)
	}
}

func MergeSortFunc2RecursiveSortFunc(msf MergeSortFunc) RecursiveSortFunc {
	return func (slice SliceType, depth Depth) {
		tmp := make(SliceType, len(slice), len(slice))
		msf(slice, tmp, depth)
	}
}

func MergeSortFunc2InPlaceSortFunc(msf MergeSortFunc) InPlaceSortFunc {
	rsf := MergeSortFunc2RecursiveSortFunc(msf)
	return RecursiveSortFunc2InPlaceSortFunc(rsf)
}

func SortFunc2NonMutatingSortFunc(ipsf InPlaceSortFunc) NonMutatingSortFunc {
	return func(slice SliceType) SliceType {
		tmp := make(SliceType, len(slice), len(slice))
		copy(tmp, slice)
		ipsf(tmp)
		return tmp
	}
}

type Sorter interface {
	Sort(SliceType)
}

type Waiter interface {
	Wait()
}

type ParallelSorter interface {
	ParallelSort(SliceType) Waiter
}

func ParallelSorter2InPlaceSortFunc(ps ParallelSorter) InPlaceSortFunc {
	return func(slice SliceType) {
		ps.ParallelSort(slice).Wait()
	}
}

type MergeSorter interface {
	MergeSort(slice SliceType, tmp SliceType)
}

type Size int

type RecursiveSortStrategy interface {
	RecursiveSort() RecursiveSortFunc
	RecursionDepthLimit() Depth
	RecursionSizeLimit() Size
}

type MergeSortStrategy interface {
	RecursiveSortStrategy()
	Merge() MergeFunc
}

type ParallelSortStrategy interface {
	ParallelSort() Sorter
	ParallelSizeLimit() Size
}

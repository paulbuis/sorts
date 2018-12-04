package parallelmergesort

import (
	"github.com/paulbuis/sorts/mergesort"
	"github.com/paulbuis/sorts/parallel"
	"github.com/paulbuis/sorts/types"
	"sync"
)

type parallelMergeSortFunc func(slice types.SliceType, tmp types.SliceType) sync.Locker

type ParallelMergeFunc func (leftDone sync.Locker, rightDone sync.Locker, left types.SliceType, right types.SliceType, tmp types.SliceType) sync.Locker

type ParallelMergeSorter struct {
	seqMergeSorter mergesort.MergeSorter
	pMerge ParallelMergeFunc
}

func NewParallelMergeSorter(seqMergeSorter mergesort.MergeSorter, pMerge ParallelMergeFunc) ParallelMergeSorter {
	return ParallelMergeSorter{seqMergeSorter: seqMergeSorter, pMerge: pMerge}
}

func Merge(merge mergesort.MergeFunc) ParallelMergeFunc {
	return func (leftDone sync.Locker, rightDone sync.Locker, left types.SliceType, right types.SliceType, tmp types.SliceType) sync.Locker {
		mergeDone := parallel.NewDone()
		go func() { 
			leftDone.Lock()
			rightDone.Lock()
			merge(left, right, tmp)
			mergeDone.Unlock()
		}()
		return mergeDone
	}
}

func (pms ParallelMergeSorter) Sort(slice types.SliceType) {
	tmp := make(types.SliceType, len(slice))
    mergeSort := parallelMergeSort(pms.seqMergeSorter, pms.pMerge)
	allDone := mergeSort(slice, tmp)
	allDone.Lock()
}

func parallelMergeSort(seqMergeSorter mergesort.MergeSorter, pMerge ParallelMergeFunc) parallelMergeSortFunc {
	// note: creating a closure to return
	var inner parallelMergeSortFunc
	inner = func (slice types.SliceType, tmp types.SliceType) sync.Locker {
		if len(slice) < 25000 { // find a better threashold!
	    	seqMergeSorter.MergeSort(slice, tmp)
			return &sync.Mutex{} //created in Unlocked state
		}

		ddTmp := mergesort.Dice2(tmp)
		ddSlice := mergesort.Dice2(slice)

		leftLeftDone := inner(ddSlice.LeftLeft, ddTmp.LeftLeft)
		leftRightDone := inner(ddSlice.LeftRight, ddTmp.LeftRight)
		rightLeftDone := inner(ddSlice.RightLeft, ddTmp.RightLeft)
		rightRightDone := inner(ddSlice.RightRight, ddTmp.RightRight)

		leftDone := pMerge(leftLeftDone, leftRightDone, ddSlice.LeftLeft, ddSlice.LeftRight, ddTmp.Left)
		rightDone := pMerge(rightLeftDone, rightRightDone, ddSlice.RightLeft, ddSlice.RightRight, ddTmp.Right)

		return pMerge(leftDone, rightDone, ddTmp.Left, ddTmp.Right, slice)
	}
	return inner
}
package mergesort

import (
	"github.com/paulbuis/sorts/types"
)

type MergeFunc func(leftSrc types.SliceType, rightSrc types.SliceType, dst types.SliceType)


type MergeSorter struct {
	InPlaceSort types.InPlaceSortFunc
	Merge MergeFunc
	MinSize int
	MaxDepth int
}

func (ms MergeSorter) RecursionDepthLimit() types.Depth {
	return types.Depth(ms.MaxDepth)
}

func (ms MergeSorter) RecursionSizeLimit() types.Size {
	return types.Size(ms.MinSize)
}


func (ms MergeSorter)Sort(slice types.SliceType) {
	tmp := make(types.SliceType, len(slice), len(slice))
	mergeSort := MergeSort(ms.InPlaceSort, ms.Merge, ms.RecursionDepthLimit(), ms.RecursionSizeLimit())
	mergeSort(slice, tmp, 0)
}

// used by parallel merge sort
func (ms MergeSorter)MergeSort(slice types.SliceType, tmp types.SliceType) {
	mergeSort := MergeSort(ms.InPlaceSort, ms.Merge, ms.RecursionDepthLimit(), ms.RecursionSizeLimit())
	mergeSort(slice, tmp, 0)
}


func MergeSort(inPlaceSort types.InPlaceSortFunc, merge MergeFunc, maxDepth types.Depth, minSize types.Size) types.MergeSortFunc {
	if minSize < 10 {
		minSize = 10
	} 
	var mergeSort types.MergeSortFunc

	mergeSort = func (slice types.SliceType, tmp types.SliceType, depth types.Depth)  {
		if types.Size(len(slice)) <= minSize /*|| depth >= maxDepth */ {
			inPlaceSort(slice)
			return
		}

		ddTmp := Dice2(tmp)
		ddSlice := Dice2(slice)
		depth += 1
		mergeSort(ddSlice.LeftLeft, ddTmp.LeftLeft, depth)
		mergeSort(ddSlice.LeftRight, ddTmp.LeftRight, depth)
		mergeSort(ddSlice.RightLeft, ddTmp.RightLeft, depth)
		mergeSort(ddSlice.RightRight, ddTmp.RightRight, depth)

		merge(ddSlice.LeftLeft, ddSlice.LeftRight, ddTmp.Left)
		merge(ddSlice.RightLeft, ddSlice.RightRight, ddTmp.Right)

		merge(ddTmp.Left, ddTmp.Right, slice)
	}
	return mergeSort
}
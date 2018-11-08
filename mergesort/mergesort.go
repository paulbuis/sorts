package mergesort

type SortFunc func(slice []int)

type MergeFunc func(leftSrc []int, rightSrc []int, dst []int)

type MergeSortFunc func(slice []int, tmp []int)

type MergeSorter struct {
	InPlaceSort SortFunc
	Merge MergeFunc 
}

func (ms MergeSorter)Sort(slice []int) {
	tmp := make([]int, len(slice), len(slice))
	mergeSort := MergeSort(ms.InPlaceSort, ms.Merge)
	mergeSort(slice, tmp)
}

func (ms MergeSorter)MergeSort(slice []int, tmp []int) {
	mergeSort := MergeSort(ms.InPlaceSort, ms.Merge)
	mergeSort(slice, tmp)
}

func MergeSort(inPlaceSort SortFunc, merge MergeFunc) MergeSortFunc {
	var mergeSort MergeSortFunc 
	mergeSort = func (slice []int, tmp []int)  {
		if len(slice) <= 30 {
			inPlaceSort(slice)
			return
		}

		ddTmp := Dice2(tmp)
		ddSlice := Dice2(slice)

		mergeSort(ddSlice.LeftLeft, ddTmp.LeftLeft)
		mergeSort(ddSlice.LeftRight, ddTmp.LeftRight)
		mergeSort(ddSlice.RightLeft, ddTmp.RightLeft)
		mergeSort(ddSlice.RightRight, ddTmp.RightRight)

		merge(ddSlice.LeftLeft, ddSlice.LeftRight, ddTmp.Left)
		merge(ddSlice.RightLeft, ddSlice.RightRight, ddTmp.Right)

		merge(ddTmp.Left, ddTmp.Right, slice)
	}
	return mergeSort
}
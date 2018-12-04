package mergesort

import (
	"github.com/paulbuis/sorts/types"
)

func IterativeMerge(left types.SliceType, right types.SliceType, slice types.SliceType) {
	indexSlice := 0
	indexLeft := 0
	indexRight := 0
	lenLeft := len(left)
	lenRight := len(right)

	for indexLeft < lenLeft && indexRight < lenRight {
		leftValue := left[indexLeft]
		rightValue := right[indexRight]
		if types.Less(left[indexLeft], right[indexRight]) {
			slice[indexSlice] = leftValue
			indexLeft += 1
		} else {
			slice[indexSlice] = rightValue
			indexRight += 1
		}
		indexSlice += 1
	}

	if indexLeft < lenLeft {
		copy(slice[indexSlice:], left[indexLeft:])
	}
	
	if indexRight < lenRight {
		copy(slice[indexSlice:], right[indexRight:])
	}
}
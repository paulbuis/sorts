package mergesort

func IterativeMerge(left []int, right []int, slice []int) {
	indexSlice := 0
	indexLeft := 0
	indexRight := 0
	lenLeft := len(left)
	lenRight := len(right)

	for indexLeft < lenLeft && indexRight < lenRight {
		leftValue := left[indexLeft]
		rightValue := right[indexRight]
		if leftValue < rightValue {
			slice[indexSlice] = leftValue
			indexLeft += 1
		} else {
			slice[indexSlice] = rightValue
			indexRight += 1
		}
		indexSlice += 1
	}

	for indexLeft < lenLeft  {
		slice[indexSlice] = left[indexLeft]
		indexSlice += 1
		indexLeft += 1
	}
	
	for indexRight < lenRight {
		slice[indexSlice] = right[indexRight];
		indexSlice += 1
		indexRight += 1
	}
}
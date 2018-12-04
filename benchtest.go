package main

import (
	"github.com/paulbuis/sorts/types"
	"sort"
	"time"
)

func benchtest(sortFun types.InPlaceSortFunc, originalList types.SliceType) time.Duration {
	sortList := make(types.SliceType, len(originalList), len(originalList))
	copy(sortList, originalList)
	start := time.Now()
	sortFun(sortList)
	elapsed := time.Since(start)
	if !sort.IsSorted(sortList) {
		out.Printf("*** sort unsuccessful! ***\n")
	}
	return elapsed
	// also check that it is a sorting of originalList !!!
}
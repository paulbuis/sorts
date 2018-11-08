package main

import (
	"sort"
	"time"
)

func benchtest(sortFun func ([]int), originalList []int, name string) {
	sortList := make([]int, len(originalList), len(originalList))
	copy(sortList, originalList)
	start := time.Now()
	sortFun(sortList)
	elapsed := time.Since(start)
	out.Printf("%s\nElapsed time %s\n", name, elapsed)
	if !sort.IntsAreSorted(sortList) {
		out.Printf("*** sort unsuccessful! ***\n")
	}
	out.Println()
	// also check that it is a sorting of originalList !!!
}
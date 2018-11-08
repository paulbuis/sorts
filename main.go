package main

import (
	"golang.org/x/text/message" // need to be sure to have GOPATH environment variable set
	                            // and "go get golang.org/x/text/message"
	"github.com/paulbuis/sorts/insertionsort"
	"math/rand"
	"runtime"
	//"sort"
	//"time"
)

var (
	out *message.Printer = message.NewPrinter(message.MatchLanguage("en-US"))
)

func main() {
	run(10000000, 4)
}

func run(n int, maxThreads int) {
	out.Printf("N=%d\n", n)
	runtime.GOMAXPROCS(maxThreads) // avoid using hyperthreading??

	out.Printf("GOMAXPROCS=%d\n", runtime.GOMAXPROCS(-1))
	out.Printf("GOMAXPROCS=%d\n", runtime.GOMAXPROCS(-1))
	out.Printf("NumCPU=%d\n", runtime.NumCPU())
	out.Printf("====================\n\n")

	var testSlice = rand.Perm(n)
	if n <= 10000 { 
		benchtest(insertionsort.Sort, testSlice, "Insertion Sort")
	}
	
	if n <= 1000000 {
		benchtest(quickSort, testSlice, "Library Quick Sort")
		benchtest(stableSort, testSlice, "Library Stable Sort")
	}

	benchtest(mergeSortIterativeMerge, testSlice, "Merge Sort: non-recursive merge")
	benchtest(mergeSortRecursiveMerge, testSlice, "Merge Sort: recursive merge")
	benchtest(parallelWithIterativeMerge, testSlice, "Parallel Merge Sort: non-recursive, non-parallel merge")
	benchtest(parallelWithRecursiveMerge, testSlice, "Parallel Merge Sort: recursive, non-parallel merge")
	benchtest(parallelWithRecursiveMerge, testSlice, "Parallel Merge Sort: recursive, parallel merge")
}
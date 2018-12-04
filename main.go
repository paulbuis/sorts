package main

import (
	"golang.org/x/text/message" // need to be sure to have GOPATH environment variable set
								// and "go get golang.org/x/text/message"
	"github.com/paulbuis/sorts/types"
	"runtime"
	"time"
)

var (
	out *message.Printer = message.NewPrinter(message.MatchLanguage("en-US"))
)

func main() {
	for i := 10000; i<=10000000; i*=10 {
		run(i, 1)
	}
}

func run(n int, maxThreads int) {
	out.Printf("N=%d\n", n)
	runtime.GOMAXPROCS(maxThreads) // avoid using hyperthreading??

	testSlice := types.Random(n)
	out.Println("Random number generation complete")
	var elapsed time.Duration
	
	elapsed = benchtest(quickSort, testSlice)
	out.Printf("Library Quick Sort\nElapsed time: %s\n\n", elapsed)
	elapsed = benchtest(dualPivotQuicksort, testSlice)
	out.Printf("Library Stable Sort\nElapsed time: %s\n\n", elapsed)
}


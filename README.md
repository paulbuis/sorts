# sorts
golang sorting, including using goroutines for parallelism

A strategy-pattern-based set of sorting algorithms written in the Go Programming Language.

Example strategy #1: Mergesort (recursive), switching to Insertionsort below some size or recursion depth threashold
using an iterative merge. In this example, the strategies include )what threashold to switch algorithms, 
2)which algorithm to switch to, and 3)the merge algorithm (also strategy based) 

Example strategy #2: Mergesort with performing each merge operation in parallel with other merge operations, and 2) using a recursive and parallel merge algoritm.

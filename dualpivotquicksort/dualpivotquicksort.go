package dualpivotquicksort

import (
	"sort"
	"github.com/paulbuis/sorts/types"
)

func Sort(slice types.SliceType) {
	dualpivot(slice, 0, slice.Len(), 3)
}

// remember to use SliceType's Swap() and Less() methods!

func dualpivot(slice sort.Interface, lo int, hi int, div int) {
  // see https://codeblab.com/wp-content/uploads/2009/09/DualPivotQuicksort.pdf

}

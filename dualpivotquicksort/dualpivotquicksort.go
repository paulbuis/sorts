package dualpivotquicksort

import (
	"github.com/paulbuis/sorts/types"
)

func Sort(slice types.SliceType) {
	dualpivot(slice, 0, len(slice)-1, 3)
}

// remember to use SliceType's Swap() and Less() methods!

func dualpivot(slice types.Interface, lo int, hi int, div int) {
  // see https://codeblab.com/wp-content/uploads/2009/09/DualPivotQuicksort.pdf

}

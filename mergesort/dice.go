package mergesort

import (
	"github.com/paulbuis/sorts/types"
)
type DoubleDice struct {
	Left types.SliceType
	Right types.SliceType
	LeftLeft types.SliceType
	LeftRight types.SliceType
	RightLeft types.SliceType
	RightRight types.SliceType
}

func Dice(slice types.SliceType) (left types.SliceType, right types.SliceType) {
	middle := len(slice) / 2
	left = slice[0:middle]
	right = slice[middle:]
	return
}

func Dice2(slice types.SliceType) (dd DoubleDice) {
	dd.Left, dd.Right = Dice(slice)
	dd.LeftLeft, dd.LeftRight = Dice(dd.Left)
	dd.RightLeft, dd.RightRight = Dice(dd.Right)
	return
}
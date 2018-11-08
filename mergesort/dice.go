package mergesort

type DoubleDice struct {
	Left []int
	Right []int
	LeftLeft []int
	LeftRight []int
	RightLeft []int
	RightRight []int
}

func Dice(slice []int) (left []int, right []int) {
	middle := len(slice) / 2
	left = slice[0:middle]
	right = slice[middle:]
	return
}

func Dice2(slice []int) (dd DoubleDice) {
	dd.Left, dd.Right = Dice(slice)
	dd.LeftLeft, dd.LeftRight = Dice(dd.Left)
	dd.RightLeft, dd.RightRight = Dice(dd.Right)
	return
}
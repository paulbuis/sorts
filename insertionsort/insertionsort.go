package insertionsort


func Sort(slice[]int) {
	left := 0
	right := len(slice) - 1

    for i := left; i <= right; i += 1 {
	   value := slice[i]
	   j := i - 1
       for j > -1 && slice[j] > value {
		  slice[j + 1] = slice[j];
		  j -= 1
       }
       slice[j + 1] = value;
    }
}
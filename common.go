package learn

func Swap(array []int, x, y int) {
	array[x], array[y] = array[y], array[x]
}

func SwapCom(array []Comparable, x, y int) {
	array[x], array[y] = array[y], array[x]
}

func SwapPtr(x, y *int) {
	*x, *y = *y, *x
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

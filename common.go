package learn

func Swap(array []int, x, y int) {
	tmp := array[x]
	array[x] = array[y]
	array[y] = tmp
}

func SwapCom(array []Comparable, x, y int) {
	tmp := array[x]
	array[x] = array[y]
	array[y] = tmp
}

func SwapPtr(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
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

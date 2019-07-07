package sort

import (
	"learn"
)

// 选中位数，放于第一位
func ThreeChoice(array []int) {
	if len(array) < 2 {
		return
	}
	li, mi, ri := 0, len(array)/2, len(array)-1
	l, m, r := array[li], array[mi], array[ri]
	if l >= m && m >= r || r >= m && m >= l {
		learn.SwapPtr(&array[0], &array[mi])
	} else if l >= r && r >= m || m >= r && r >= l {
		learn.SwapPtr(&array[0], &array[ri])
	} else if m >= l && l >= r || r >= l && l >= m {
		learn.SwapPtr(&array[0], &array[li])
	}
}

package sort

import (
	"learn"
	"math/rand"
)

func QuickSort(array []int) []int {
	// 双指针搜索
	// 随机一个index
	// 更优秀的算法是三取样切分,取三者中位数
	length := len(array)
	if length == 0 {
		return array
	}
	index := rand.Intn(length)
	value := array[index]
	learn.SwapPtr(&array[0], &array[index])
	lo, hi := 1, length-1 // array[lo] <= value < array[hi]
	for lo <= hi {
		if array[lo] > value && array[hi] < value {
			learn.SwapPtr(&array[lo], &array[hi])
			lo++
			hi-- // lval < value
		} else if array[lo] > value {
			hi--
		} else if array[hi] < value {
			lo++
		} else {
			lo++
			hi--
		}
	}
	// hi < lo
	learn.SwapPtr(&array[0], &array[hi])
	QuickSort(array[:hi])
	QuickSort(array[hi+1:])
	return array
}

// 适合多个重复值数组
func QuickSort3Way(array []int) []int {
	length := len(array)
	if length < 2 {
		return array
	}
	ThreeChoice(array) // 随机化
	// [0, lt) <
	// [lt, i) ==
	// [i,gt) unknown
	// [gt, len) >
	lt, gt, i, v := 0, length, 1, array[0] // 目前完全符合预期
	for i < gt {                           // 按照语义 gt > i
		t := array[i]
		if t < v {
			learn.SwapPtr(&array[i], &array[lt])
			lt++
			i++
		} else if t == v {
			i++
		} else { // v < t
			gt--
			learn.SwapPtr(&array[gt], &array[i])
		}
	}
	QuickSort3Way(array[:lt])
	QuickSort3Way(array[gt:])
	return array
}

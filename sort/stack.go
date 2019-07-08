package sort

import (
	"learn"
	arr "learn/array"
	"learn/queue"
)

// 堆排序
type icomparable int

func (i icomparable) Compare(data interface{}) int {
	other := data.(icomparable)
	if other > i {
		return -1
	} else if other == i {
		return 0
	} else {
		return 1
	}
}

// 这种形式不好的原因，因为堆的构建并不需要这么次插入，实际上，一个数组只需要一半的时间即可构建
func StackSortByQueue(array []int) []int {
	// 最low的办法
	queue := queue.NewPriorityQueue()
	for _, value := range array {
		queue.Push(icomparable(value))
	}
	queue.Sort()
	sortedArray, length := queue.Array, len(array)
	for i := queue.Size; i > 0; i-- {
		j := length - i
		array[j] = int(sortedArray[i].(icomparable))
	}
	return array
}

func StackSort(array []int) []int {
	array = append([]int{-1}, array...)
	length := len(array)
	for N := (length - 1) / 2; N >= 1; N-- {
		Sink(N, length-1, array)
	}
	for N := length - 1; N > 1; N-- {
		learn.Swap(array, N, 1)
		Sink(1, N-1, array)
	}
	return arr.Reverse(array[1:])
}

func Sink(i, N int, array []int) {
	for i <= N/2 {
		var min int
		l, r := i*2, i*2+1
		if r > N || array[l] < array[r] {
			min = l
		} else {
			min = r
		}
		if array[min] < array[i] {
			learn.Swap(array, i, min)
		} else {
			break
		}
		i = min
	}
}

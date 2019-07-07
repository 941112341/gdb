package sort

import "learn/queue"

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

func StackSort(array []int) []int {
	// 最low的办法
	queue := queue.PriorityQueue{}
	for _, value := range array {
		queue.Push(icomparable(value))
	}
	for i := 0; i < queue.Size; i++ {
		array[i] = int(queue.Pop().(icomparable))
	}
	return array
}

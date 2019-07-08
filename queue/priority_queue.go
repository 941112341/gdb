package queue

import "learn"

type PriorityQueue struct {
	Array []learn.Comparable
	Size  int // Size == len(array) - 1, 实际元素数
}

func Build(array []learn.Comparable) *PriorityQueue {
	array = append([]learn.Comparable{nil}, array...)
	length := len(array)
	queue := PriorityQueue{Array: array, Size: length - 1}
	for N := (length - 1) / 2; N >= 1; N-- {
		queue.Sink(N, length-1)
	}
	return &queue
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		Array: []learn.Comparable{nil}, // 必须又一个初始化元素
	}
}

func (queue *PriorityQueue) Push(data learn.Comparable) {
	queue.Array = append(queue.Array, data)
	queue.Size++
	queue.Lift(queue.Size)
}

func (queue *PriorityQueue) Pop() learn.Comparable {
	if queue.Size == 0 {
		return nil
	}
	elem := queue.Array[1]
	learn.SwapCom(queue.Array, 1, queue.Size)
	queue.Sink(1, queue.Size)
	queue.Array = queue.Array[:queue.Size]
	queue.Size--
	return elem
}

func (queue *PriorityQueue) Peek() learn.Comparable {
	if queue.Size == 0 {
		return nil
	}
	return queue.Array[0]
}

// 上浮
func (queue *PriorityQueue) Lift(i int) { // i = len(array) - 1
	for ; i > 1 && learn.Lt(queue.Array[i], queue.Array[i/2]); i /= 2 {
		learn.SwapCom(queue.Array, i, i/2)
	}
}

// 下沉 [i, N] 必须都在 [0, len - 1] => i * 2 必定安全
func (queue *PriorityQueue) Sink(i, N int) {
	array := queue.Array
	for i <= N/2 {
		var min int
		l, r := i*2, i*2+1
		if r > N || learn.Lt(array[l], array[r]) {
			min = l
		} else {
			min = r
		}

		if learn.Lt(array[min], array[i]) {
			learn.SwapCom(array, i, min)
		} else {
			break
		}
		i = min
	}
}

func (queue *PriorityQueue) Sort() {
	array := queue.Array
	if queue.Size < 2 {
		return
	}
	N := queue.Size
	//M := N / 2
	learn.SwapCom(array, 1, N)
	for N--; N > 1; N-- {
		queue.Sink(1, N)
		learn.SwapCom(array, 1, N)
	}
}

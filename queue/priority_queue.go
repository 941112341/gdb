package queue

import "learn"

type PriorityQueue struct {
	Array []learn.Comparable
	Size  int
}

func (queue *PriorityQueue) Push(data learn.Comparable) {
	queue.Array = append(queue.Array, data)
	// 上浮
	queue.Up(queue.Size)
	queue.Size++
}

func (queue *PriorityQueue) Pop() learn.Comparable {
	if queue.Size == 0 {
		return nil
	}
	queue.Size--
	elem := queue.Array[0]
	queue.Array = queue.Array[1:]
	return elem
}

func (queue *PriorityQueue) Peek() learn.Comparable {
	if queue.Size == 0 {
		return nil
	}
	return queue.Array[0]
}

// x >= 0
func (queue *PriorityQueue) Up(i int) {
	for ; i > 0 && queue.Array[i/2].Compare(queue.Array[i]) > 0; i /= 2 {
		learn.SwapCom(queue.Array, i, i/2)
	}
}

func (queue *PriorityQueue) Down() {
	for i := 0; ; {

	}
}

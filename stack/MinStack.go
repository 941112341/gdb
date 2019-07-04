package stack

import "learn"

type MinStack struct {
	elements   *Stack
	minElments *Stack
}

func (stack *MinStack) Pop() interface{} {
	elem := stack.elements.Pop()
	if elem == stack.minElments.Peek() {
		stack.minElments.Pop()
	}
	return elem
}

func (stack *MinStack) Peek() interface{} {
	return stack.elements.Peek()
}

func (stack *MinStack) Min() interface{} {
	return stack.minElments.Peek()
}

func (stack *MinStack) Push(data interface{}) {
	stack.elements.Push(data)
	compareable, ok := data.(learn.Comparable)
	if !ok {
		return
	}
	if compareable.Compare(stack.minElments.Peek()) < 0 {
		stack.minElments.Push(data)
	}
}

package stack

import "testing"

func TestStack(t *testing.T) {
	i := []interface{}{1, 3, 5, 7, 9}
	queue := NewQueue(i)
	queue.Println()
	stack := NewStack(i)
	stack.Println()
}

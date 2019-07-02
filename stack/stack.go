package stack

import "fmt"

type Array struct {
	len int
	element []interface{}
}

func (array *Array) PshFirst(data interface{}) {
	array.len++
	array.element = append([]interface{}{data}, array.element...)
}

func (array *Array) PushLast(data interface{}) {
	array.len++
	array.element = append(array.element, data)
}

func (array *Array) IsEmpty() bool {
	return array.len == 0
}

func (array *Array) PopFirst() interface{} {
	if array.IsEmpty() {
		return nil
	}
	array.len--
	elem := array.element[0]
	array.element = array.element[1:]
	return elem
}


func (array *Array) PopLast() interface{} {
	if array.IsEmpty() {
		return nil
	}
	array.len--
	elem := array.element[array.len-1]
	array.element = array.element[:array.LastIndex()]
	return elem
}

func (array *Array) PeekFirst() interface{} {
	if array.IsEmpty() {
		return nil
	}
	return array.element[0]
}


func (array *Array) PeekLast() interface{} {
	if array.IsEmpty() {
		return nil
	}
	return array.element[array.LastIndex()]
}

func (array *Array) LastIndex() int {
	return array.len - 1
}

func (array *Array) Println() {
	fmt.Println("%#v", array)
}


type Stack LinerProto

type Queue LinerProto

type Liner interface {
	Pop() interface{}
	Peek() interface{}
	Push(data interface{})
	Println()
}

type LinerProto struct {
	Array Array
}

func (proto *LinerProto) Println() {
	elem := proto.Pop()
	for elem != nil {
		fmt.Print(elem, "\t")
	}
	fmt.Println()
}

func (proto *LinerProto) Pop() interface{} {
	panic("implement me")
}

func (proto *LinerProto) Peek() interface{} {
	panic("implement me")
}

func (proto *LinerProto) Push(data interface{}) {
	proto.Array.PushLast(data)
}

func (Stack *Stack) Pop() interface{} {
	return 	Stack.Array.PopLast()
}

func (stack *Stack) Peek() interface{} {
	return stack.Array.PeekLast()
}

func (Queue *Queue) Pop() interface{} {
	return 	Queue.Array.PopLast()
}

func (Queue *Queue) Peek() interface{} {
	return Queue.Array.PeekLast()
}

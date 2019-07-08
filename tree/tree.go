package tree

import "learn"

type Node interface {
	Add(value learn.Comparable) *Node
	//Del
	Select(value learn.Comparable) int
	Size() int
	Max() *Node
	Min() *Node
	Floor() *Node
	Ceiling() *Node
}

var Leaf = NewTree(nil, nil, nil)

type Tree struct {
	value       learn.Comparable
	left, right *Tree
	count       int
}

func NewTree(value learn.Comparable, left, right *Tree) *Tree {
	return &Tree{
		value: value, left: left, right: right,
	}
}

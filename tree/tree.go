package tree

import "learn"



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

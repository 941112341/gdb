package tree

import "fmt"

type Node2 struct {
	obj  interface{}
	i    int
	l, r *Node2
}

type Node3 struct {
	Node2
	m      *Node2
	j      int
	attach interface{}
}

type Tree struct {
	root *Node2
}

func (node *Node2) Print() {
	if node == nil {
		return
	}
	node.l.Print()
	fmt.Print("\t", node.obj)
	node.r.Print()
}

func (node *Node3) Print() {
	if node == nil {
		return
	}
	node.l.Print()
	fmt.Print("\t", node.attach)
	node.m.Print()
	fmt.Println("\t", node.obj)
	node.r.Print()
}

func (node *Node2) Get(x int) (*Node2, interface{}) {
	if node == nil {
		return nil, nil
	}
	if node.i == x {
		return nil, node.obj
	}
	if node.i > x {
		parent, obj := node.r.Get(x)
		if parent == nil && obj == nil {
			return nil, nil
		}
		if parent == nil {
			return node, obj
		}
		return parent, obj
	}
	if node.i < x {
		parent, obj := node.l.Get(x)
		if parent == nil && obj == nil {
			return nil, nil
		}
		if parent == nil {
			return node, obj
		}
		return parent, obj
	}
	return nil, nil
}

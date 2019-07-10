package table

import (
	"learn"
	"learn/stack"
)

type Table interface {
	Put(key learn.Comparable, value interface{}) (Table, interface{}) // old value
	PutNX(key learn.Comparable, value interface{}) (Table, bool)
	Get(key learn.Comparable) interface{}
	Del(key learn.Comparable) Table // 返回newTable
	Exist(key learn.Comparable) bool
	Size() int
	Rank(key learn.Comparable) int
	Max() Table
	Min() Table
	DelMax() Table
	DelMin() Table
	Floor(key learn.Comparable) Table
	Ceiling(key learn.Comparable) Table
	Keys() []Table
	Range(lk, rk learn.Comparable) []Table
	Height() int
	Merge(table Table) Table
}

func equal(node *TreeNode, key learn.Comparable) bool {
	return learn.Eq(node.key, key)
}

func less(node *TreeNode, key learn.Comparable) bool {
	return learn.Lt(node.key, key)
}

func greater(node *TreeNode, key learn.Comparable) bool {
	return learn.Gt(node.key, key)
}

type TreeNode struct {
	key                 learn.Comparable
	value               interface{}
	left, right, parent *TreeNode
}

func (node *TreeNode) Height() int {
	if node == nil {
		return 0
	}
	return learn.Max(node.right.Height(), node.left.Height()) + 1
}

func (node *TreeNode) Put(key learn.Comparable, value interface{}) (Table, interface{}) {
	if node == nil {
		return &TreeNode{
			key: key, value: value,
		}, nil
	}
	if equal(node, key) {
		elem := node.value
		node.value = value
		return node, elem
	} else if less(node, key) {
		table, elem := node.right.Put(key, value)
		tableNode := table.(*TreeNode)
		tableNode.parent, node.right = node, tableNode
		return table, elem
	} else {
		table, elem := node.left.Put(key, value)
		tableNode := table.(*TreeNode)
		tableNode.parent, node.left = node, tableNode
		return table, elem
	}
}

func (node *TreeNode) PutNX(key learn.Comparable, value interface{}) (Table, bool) {
	if node == nil {
		return &TreeNode{
			key: key, value: value,
		}, true
	}
	if equal(node, key) { // 不插入
		return node, false
	} else if less(node, key) {
		table, elem := node.right.PutNX(key, value)
		tableNode := table.(*TreeNode)
		tableNode.parent, node.right = node, tableNode
		return table, elem
	} else {
		table, elem := node.left.PutNX(key, value)
		tableNode := table.(*TreeNode)
		tableNode.parent, node.left = node, tableNode
		return table, elem
	}
}

func (node *TreeNode) Get(key learn.Comparable) interface{} {
	if node == nil {
		return nil
	}
	if equal(node, key) {
		return node.value
	} else if less(node, key) {
		return node.right.Get(key)
	} else {
		return node.left.Get(key)
	}
}

func (node *TreeNode) Merge(table Table) Table {
	other := table.(*TreeNode)
	if node == nil {
		return other
	} else if other == nil {
		return node
	}
	if node.Height() >= table.Height() { //
		if learn.Lte(node.key, other.key) {
			lNode, rNode := other.left, node.right
			node.right, other.parent = other, node
			other.left, _ = rNode.Merge(lNode).(*TreeNode)
			return node
		} else {
			lNode, rNode := node.left, other.right
			node.left, other.parent = other, node
			other.right, _ = rNode.Merge(lNode).(*TreeNode)
			return node
		}
	} else {
		if learn.Lte(node.key, other.key) {
			lNode, rNode := other.left, node.right
			other.left, node.parent = node, other
			node.right, _ = lNode.Merge(rNode).(*TreeNode)
			return other
		} else {
			lNode, rNode := node.left, other.right
			other.right, node.parent = node, other
			node.left, _ = lNode.Merge(rNode).(*TreeNode)
			return other
		}
	}
}

func (node *TreeNode) Del(key learn.Comparable) Table {
	if node == nil {
		return nil
	}
	if equal(node, key) {
		dNode := node.left.Merge(node.right).(*TreeNode)
		parent := node.parent
		if parent != nil {
			if parent.left == node {
				parent.left = dNode
			} else {
				parent.right = dNode
			}
		}
		dNode.parent = parent
		return dNode
	} else if less(node, key) {
		nNode, ok := node.right.Del(key).(*TreeNode)
		if ok {
			node.right, nNode.parent = nNode, node
		}
		return node
	} else {
		nNode, ok := node.left.Del(key).(*TreeNode)
		if ok {
			node.left, nNode.parent = nNode, node
		}
		return node
	}
}

func (node *TreeNode) Exist(key learn.Comparable) bool {
	return node.Get(key) != nil
}

func (node *TreeNode) Size() int {
	if node == nil {
		return 0
	}
	return node.right.Size() + 1 + node.left.Size()
}

func (node *TreeNode) Rank(key learn.Comparable) int {
	return node.left.Size()
}

func (node *TreeNode) Max() Table {
	if node == nil {
		return node.parent
	}
	return node.right.Max()
}

func (node *TreeNode) Min() Table {
	if node == nil {
		return node.parent
	}
	return node.left.parent
}

func (node *TreeNode) DelMax() Table {
	mNode := node.Max().(*TreeNode)
	mNode.Del(mNode.key)
	return node
}

func (node *TreeNode) DelMin() Table {
	mNode := node.Min().(*TreeNode)
	mNode.Del(mNode.key)
	return node
}

func (node *TreeNode) Floor(key learn.Comparable) Table {
	if node == nil {
		return nil
	}
	if equal(node, key) {
		return node
	} else if less(node, key) {
		rNode, ok := node.right.Floor(key).(*TreeNode) // <
		if !ok {                                       // 返回的是nil
			return node
		}
		return rNode
	} else {
		return node.left.Floor(key)
	}
}

func (node *TreeNode) Ceiling(key learn.Comparable) Table {
	if node == nil {
		return nil
	}
	if equal(node, key) {
		return node
	} else if less(node, key) {
		return node.left.Ceiling(key)
	} else {
		lNode, ok := node.left.Ceiling(key).(*TreeNode) // <
		if !ok {                                        // 返回的是nil
			return node
		}
		return lNode
	}
}

func (node *TreeNode) Keys() []Table { //顺序, 使用数组完成遍历
	s := stack.Stack{}
	for node != nil && !s.IsEmpty() {

	}
	return nil
}

func (node *TreeNode) Range(lk, rk learn.Comparable) []Table {
	panic("implement me")
}

type TreeTable struct {
	max, min, root *TreeNode
	size           int
}

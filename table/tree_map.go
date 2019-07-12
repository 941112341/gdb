package table

import (
	"learn"
	"learn/generator"
	"learn/reflection"
	"learn/stack"
	"math/rand"
	"reflect"
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

func (node *TreeNode) Put(key learn.Comparable, value interface{}) (*TreeNode, interface{}) {
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
		treeNode, elem := node.right.Put(key, value)
		treeNode.parent, node.right = node, treeNode
		return treeNode, elem
	} else {
		treeNode, elem := node.left.Put(key, value)
		treeNode.parent, node.left = node, treeNode
		return treeNode, elem
	}
}

func (node *TreeNode) PutNX(key learn.Comparable, value interface{}) (*TreeNode, bool) {
	if node == nil {
		return &TreeNode{
			key: key, value: value,
		}, true
	}
	if equal(node, key) { // 不插入
		return node, false
	} else if less(node, key) {
		treeNode, elem := node.right.PutNX(key, value)
		treeNode.parent, node.right = node, treeNode
		return treeNode, elem
	} else {
		treeNode, elem := node.left.PutNX(key, value)
		treeNode.parent, node.left = node, treeNode
		return treeNode, elem
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

func (node *TreeNode) Del(key learn.Comparable) *TreeNode {
	if node == nil {
		return nil
	}
	if less(node, key) {
		node.right = node.right.Del(key)
	} else if greater(node, key) {
		node.left = node.left.Del(key)
	} else {
		if generator.RandBool() { // 右为根

		}
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
		if node == nil {
			node = s.Pop().(*TreeNode).right
		} else {
			s.Push(node)
			node = node.left
		}
	}
	return reflection.Cast(s.Elements(), reflect.TypeOf([]Table{})).([]Table)
}

func (node *TreeNode) Range(lk, rk learn.Comparable) []Table {
	if node == nil {
		return []Table{}
	}
	if less(node, lk) {
		return node.right.Range(lk, rk)
	} else if greater(node, rk) {
		return node.left.Range(lk, rk)
	} else {
		array := []Table{node}
		array = append(array, node.right.Range(lk, rk)...)
		array = append(node.left.Range(lk, rk), array...)
		return array
	}
}

type TreeTable struct {
	max, min, root *TreeNode
	size           int
}

func (table *TreeTable) Put(key learn.Comparable, value interface{}) (Table, interface{}) {
	panic("implement me")
}

func (table *TreeTable) PutNX(key learn.Comparable, value interface{}) (Table, bool) {
	panic("implement me")
}

func (table *TreeTable) Get(key learn.Comparable) interface{} {
	panic("implement me")
}

func (table *TreeTable) Del(key learn.Comparable) Table {
	panic("implement me")
}

func (table *TreeTable) Exist(key learn.Comparable) bool {
	panic("implement me")
}

func (table *TreeTable) Size() int {
	panic("implement me")
}

func (table *TreeTable) Rank(key learn.Comparable) int {
	panic("implement me")
}

func (table *TreeTable) Max() Table {
	panic("implement me")
}

func (table *TreeTable) Min() Table {
	panic("implement me")
}

func (table *TreeTable) DelMax() Table {
	panic("implement me")
}

func (table *TreeTable) DelMin() Table {
	panic("implement me")
}

func (table *TreeTable) Floor(key learn.Comparable) Table {
	panic("implement me")
}

func (table *TreeTable) Ceiling(key learn.Comparable) Table {
	panic("implement me")
}

func (table *TreeTable) Keys() []Table {
	panic("implement me")
}

func (table *TreeTable) Range(lk, rk learn.Comparable) []Table {
	panic("implement me")
}

func (table *TreeTable) Height() int {
	panic("implement me")
}

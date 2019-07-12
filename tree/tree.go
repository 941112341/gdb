package tree

import "learn"

// 2 - 3 树
type Tree23 struct {
	lval, rval       learn.Comparable // 24
	left, mid, right *Tree23          // 24
	parent           *Tree23          // 8
}

package tree

type Node struct {
	Red         bool
	Data        int
	Left, Right *Node
}

func rotateRight(h *Node) *Node { // h.Right != nil
	r := h.Right
	h.Right = r.Left
	r.Left = h
	r.Red = h.Red
	h.Red = true
	return r
}

func rotateLeft(h *Node) *Node { // h.Left != ni
	l := h.Left
	h.Left = l.Right
	l.Right = h
	l.Red = h.Red
	h.Red = true
	return l
}

func (node *Node) IsRed() bool {
	if node == nil {
		return false
	}
	return node.Red
}

func (node *Node) Insert(i int) *Node {
	if node == nil {
		return &Node{Red: true, Data: i}
	}
	if node.Data >= i {
		node.Left = node.Left.Insert(i)
		if node.Left.IsRed() && node.Left.Left.IsRed() {
			node.Left = rotateLeft(node.Left)
		} else if node.Left.IsRed() && node.Left.Right.IsRed() {
			node.Left = rotateRight(node.Left)
			node.Left = rotateLeft(node.Left)
		}
	} else {
		node.Right = node.Right.Insert(i)
		if node.Right.IsRed() && node.Right.Right.IsRed() {
			node.Right = rotateRight(node)
		} else if node.Right.IsRed() && node.Right.Left.IsRed() {
			node.Right = rotateLeft(node.Right)
			node.Right = rotateRight(node.Right)
		}
	}
	return node
}

func (node *Node) height() int {
	if node == nil {
		return 0
	}
	lh, rh := node.Left.height(), node.Right.height()
	if lh > rh {
		return lh + 1
	} else {
		return rh + 1
	}
}

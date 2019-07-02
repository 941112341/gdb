package list

import "fmt"

type List struct {
	head *Node
}

type Node struct {
	next *Node
	data interface{}
}


// REVERSE

func (list *List) Reverse() {
	head := list.head
	next := head.next
	head.next = nil

	// head.next right , next is wrong
	// should next -> head
	for next != nil {
		nNext := next.next
		next.next = head
		head = next
		next = nNext
	}

	list.head = head
}

func (list *List)Println()  {
	node := list.head
	for node != nil {
		fmt.Print("\t",node.data, "\t")
		node = node.next
	}
	fmt.Println()
}

func (node *Node) Reverse() (head *Node, tail *Node) {

	if node.next == nil {
		head, tail = node, node
		return
	}
	// x1 + f(x...) => x1 + (xn...2)
	head, tail = node.next.Reverse()
	tail.next = node
	node.next = nil
	return head, node
}

func (list *List) ReverseRec()  {
	if list.head == nil {
		return
	}
	list.head , _ = list.head.Reverse()
}

func (list *List) InsertReverse()  {
	head := list.head
	if head == nil {
		return
	}
	next := head.next

	// list -> head -> next -> .....
	for next != nil {
		// begin insert
		nNext := next.next
		next.next = list.head
		list.head = next
		next = nNext
	}
	head.next = nil
}
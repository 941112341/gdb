package tree

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	var node *Node
	ints := []int{2, 5, 7, 8, 9, 11, 2, 80, 55}
	for _, value := range ints {
		node = node.Insert(value)
	}
	//fmt.Println(node)
	fmt.Println(node.height())
}

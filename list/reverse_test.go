package list

import (
	"testing"
	"time"
)

var nodes = []Node {
	{
		data:1,
	},
	{
		data:2,
	},
	{
		data:3,
	},
	{
		data:4,
	},
	{
		data:5,
	},
	{
		data:6,
	},
}

func init() {
	for key := range nodes {
		if key != len(nodes) - 1 {
			nodes[key].next = &nodes[key + 1]
		}
	}
}

func TestReverse(t *testing.T) {
	list := List{head: &nodes[0],}
	list.Println()
	list.Reverse()
	list.Println()
}

func TestList_ReverseRec(t *testing.T) {
	go func() {
		list := List{head: &nodes[0],}
		list.ReverseRec()
		list.Println()
	}()
	time.Sleep(time.Second * 2)
}

func TestList_InsertReverse(t *testing.T) {
	go func() {
		list := List{head: &nodes[0],}
		list.InsertReverse()
		list.Println()
	}()

	time.Sleep(time.Second)
}
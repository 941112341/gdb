package queue

import (
	"fmt"
	"learn"
	"testing"
)

type icomparable int

func (i icomparable) Compare(data interface{}) int {
	other := data.(icomparable)
	if other > i {
		return -1
	} else if other == i {
		return 0
	} else {
		return 1
	}
}

func TestBuild(t *testing.T) {
	queue := Build([]learn.Comparable{icomparable(4), icomparable(3), icomparable(7), icomparable(8), icomparable(1), icomparable(3)})
	fmt.Println(queue)
}

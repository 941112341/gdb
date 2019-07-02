package list

import (
	"fmt"
	"testing"
)

func TestList_GetLastK(t *testing.T) {
	list := List{head: &nodes[0],}
	fmt.Print(list.GetLastK(3))
}

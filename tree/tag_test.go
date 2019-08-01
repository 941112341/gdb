package tree

import (
	"fmt"
	"testing"
)

func TestTag(t *testing.T) {
	fmt.Print("g")
	if 1 == 2 {
		goto tag
	}
tag:
	fmt.Println("?")
}

package learn

import (
	"fmt"
	"testing"
)

func TestSwapPtr(t *testing.T) {
	a := []int{1, 2, 3}
	SwapPtr(&a[0], &a[2])
	fmt.Print(a)
}

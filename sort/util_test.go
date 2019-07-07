package sort

import (
	"fmt"
	"testing"
)

func TestThreeChoice(t *testing.T) {
	arr := []int{3, 4, 5, 6, 7, 8}
	ThreeChoice(arr)
	fmt.Println(arr)
}

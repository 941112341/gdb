package learn

import (
	"fmt"
	"testing"
)

func TestBinarySearchRec(t *testing.T) {
	array := []int{1, 3, 5, 7, 9, 12, 16, 20, 25, 30}
	i := BinarySearchRec(array, 12, 0, len(array))
	fmt.Println(i)
	i = BinarySearchRec(array, 25, 0, len(array))
	fmt.Println(i)
	i = BinarySearchRec(array, 1, 0, len(array))
	fmt.Println(i)
	i = BinarySearchRec(array, 6, 0, len(array))
	fmt.Println(i)
}
func TestBinarySearch(t *testing.T) {
	array := []int{1, 3, 5, 7, 9, 12, 16, 20, 25, 30}
	i := BinarySearch(array, 12, 0, len(array)-1)
	fmt.Println(i)
	i = BinarySearch(array, 25, 0, len(array)-1)
	fmt.Println(i)
	i = BinarySearch(array, 1, 0, len(array)-1)
	fmt.Println(i)
	i = BinarySearch(array, 6, 0, len(array)-1)
	fmt.Println(i)
	i = BinaryFind(array, 15, 0, len(array))
	fmt.Println(i)
}

func TestSplit(t *testing.T) {
	array := []int{3, 5, 7}
	fmt.Println(array[:2])
	fmt.Println(array[:0])
}

func TestBinaryFindDef(t *testing.T) {
	array := []int{1, 2, 5, 7, 21, 44, 45, 71}
	i := BinaryFindDef(array, 2)
	fmt.Println(i)
	i = BinaryFindDef(array, 4)
	fmt.Println(i)
	i = BinaryFindDef(array, 14)
	fmt.Println(i)
	i = BinaryFindDef(array, 34)
	fmt.Println(i)
	i = BinaryFindDef(array, 64)
	fmt.Println(i)
	i = BinaryFindDef(array, 88)
	fmt.Println(i)
}

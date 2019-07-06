package array

import (
	"fmt"
	"testing"
)

func TestMigrate(t *testing.T) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Migrate(array, 1, 4, 2)
	fmt.Println(array)
	Migrate(array, 2, 8, -3)
	fmt.Println(array)
}

func TestMerge2(t *testing.T) {
	array := []int{1, 2, 5, 6, 7, 9}
	array2 := []int{0, 1, 3, 5, 8, 11}
	fmt.Println(MergeSorted(array, array2))
}

func TestArrayCopy(t *testing.T) {
	array := []int{1, 2, 5, 6, 7, 9}
	array2 := []int{0, 1, 3, 5, 8, 11}
	ArrayCopy(array2, array, 2, 0, 4)
	fmt.Println(array)
}

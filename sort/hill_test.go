package sort

import (
	"fmt"
	"learn/generator"
	"testing"
)

func TestHillSort(t *testing.T) {
	var array = generator.RandomArray(20, 100)
	HillSort(array)
	generator.AssertSorted(array)
	fmt.Print("success")
}

func TestMerge(t *testing.T) {
	var array = generator.RandomArray(20, 100)
	//InnerMerge(array)
	//fmt.Println(array)
	//array = Merge(array)
	//generator.AssertSorted(array)
	array = MergeFromBottom(array)
	fmt.Println(array)
	generator.AssertSorted(array)
	fmt.Print("success")
}

func TestSplit(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6}
	subarray := array[2:5]
	subarray[0] = -1
	fmt.Println(array)
}

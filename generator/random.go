package generator

import "math/rand"

func RandomArray(n, max int) []int {
	array := make([]int, n)
	for key := range array {
		array[key] = rand.Intn(max)
	}
	return array
}

func IsSorted(array []int) bool {
	len := len(array)
	if len == 0 || len == 1 {
		return true
	}
	for i := 1; i < len; i++ {
		if array[i-1] > array[i] {
			return false
		}
	}
	return true
}

func AssertSorted(array []int) {
	if !IsSorted(array) {
		panic("not sorted")
	}
}

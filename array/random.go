package array

import "math/rand"

func RandomArray(n, max int) []int {
	array := make([]int, n)
	for key := range array {
		array[key] = rand.Intn(max)
	}
	return array
}

func IsSorted(array []int) bool {
	if len(array) == 0 {
		return true
	}
	x := array[0]
	for _, value := range array {
		if value < x {
			return false
		}
	}
	return true
}



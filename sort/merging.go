package sort

import (
	"learn"
	arr "learn/array"
)

// 额外归并
func Merge(array []int) []int {
	if len(array) < 2 {
		return array
	}
	lo, hi := 0, len(array)
	mi := (lo + hi) / 2
	lArr, rArr := Merge(array[:mi]), Merge(array[mi:])
	return arr.MergeSorted(lArr, rArr)
}

// 更高效的原地归并
// fx = t(f[:m], f[m:])
// 证明f[:m], f[m:] t() 收敛，以及index valid
func InnerMerge(array []int) []int {
	alen := len(array)
	if alen < 2 {
		return array
	}
	mid := alen / 2
	// 特证明 len = 0 and 1 and 2
	// len = 0， 1无法收敛，设为起始条件
	larr, rarr := InnerMerge(array[:mid]), InnerMerge(array[mid:])
	// 插入排序
	for key, value := range rarr {
		i := learn.BinaryFind(array, value, 0, len(larr)+key-1)
		arr.Migrate(array, i, len(larr)-1+key, 1)
		array[i] = value
	}
	return array
}

func MergeFromBottom(array []int) []int {
	length := len(array)
	for x := 1; x < length; x = x * 2 {
		// 确实收敛
		for i := 0; i < length-x; i += 2 * x {
			// 左半区域一定index valid
			// 右半区域一定可取
			sortedArray := arr.MergeSorted(array[i:i+x], array[i+x:learn.Min(length, i+2*x)]) // 如何使用条件上推?
			arr.ArrayCopy(sortedArray, array, 0, i, len(sortedArray))
		}
	}
	return array
}

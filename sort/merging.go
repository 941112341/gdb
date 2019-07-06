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
	// for 代表合并个数 [2, x > length]
	for i := 2; i < length; i *= 2 {
		// i index out of range?
		// j valid
		for j := 0; j < length; j += i { // 是否收敛?

			hi := j + i
			if hi > length {
				hi = length
			}
			mi := (hi + j) / 2
			//收敛？尤其是右边缺失部分的证明逻辑比较复杂
			//
			tmp := arr.MergeSorted(array[j:mi], array[mi:hi])
			arr.ArrayCopy(tmp, array, 0, j, len(tmp))
		}
	}
	return array
}

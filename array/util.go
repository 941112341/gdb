package array

import (
	"fmt"
)

// [from, to]
// 		[from + distance, to + distance]
// [0, len - 1]
func Migrate(array []int, from, to, distance int) {
	alen := len(array)
	if from < 0 || to >= alen {
		panic(fmt.Errorf("[%v, %v] len=%v", from, to, alen))
	}
	from, to = from+distance, to+distance
	if distance > 0 {
		if to >= alen {
			to = alen - 1
		}

		// 前提 from >= 0, to < alen
		// from - distance安全
		for ; from <= to; to-- {
			array[to] = array[to-distance]
		}
	} else if distance < 0 {
		if from < 0 {
			from = 0
		}

		// 前提 from >= 0, to < alen
		// from - distance安全
		for ; from <= to; from++ {
			array[from] = array[from-distance]
		}
	}

}

func MergeSorted(lArr, rArr []int) []int {
	llen, rlen := len(lArr), len(rArr)
	lo, hi := 0, llen+rlen
	tmpArr := make([]int, hi)
	for i, j := 0, 0; lo < hi; lo++ {
		if i >= llen {
			for ; j < rlen; j++ {
				tmpArr[lo] = rArr[j]
				lo++
			}
		} else if j >= rlen {
			for ; i < llen; i++ {
				tmpArr[lo] = lArr[i]
				lo++
			}
		} else {
			if lArr[i] < rArr[j] {
				tmpArr[lo] = lArr[i]
				i++
			} else {
				tmpArr[lo] = rArr[j]
				j++
			}
		}
	}
	return tmpArr
}

// from 是source起点，to是destination起点
func ArrayCopy(source, destination []int, from, to, length int) {
	copy(destination[to:to+length], source[from:from+length])
}

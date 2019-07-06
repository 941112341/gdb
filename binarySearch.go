package learn

func BinarySearchDef(array []int, value int) int {
	return BinarySearch(array, value, 0, len(array))
}

func BinaryFindDef(array []int, value int) int {
	return BinaryFind(array, value, 0, len(array)-1)
}

// 找到第一个value大于参数的index
// 思路：input [start, end]
// 对于[start, end)模式来说缺点在于算法本身会有游动性，可能偏左，或右取决于奇偶性, n = end - start output: [start, end + 1]
// f(array) = f(array[:mid]) | mVal > val
// 			  f(array[mid+1:) |  mval <= val
// 证明区间 必定收敛，必定不index out of range
func BinaryFind(array []int, value, start, end int) int {
	for start <= end {
		mid := (start + end) / 2 // 这样永远偏右是比较节省代码的
		mval := array[mid]
		if value < mval {
			if end == mid {
				// 死循环
				return end
			}
			end = mid
		} else {
			start = mid + 1 // nstart (ostart,end+1]
		}
	}
	return start
}

func BinarySearch(array []int, value, start, end int) int {
	for start <= end {
		mid := (start + end) / 2
		mValue := array[mid]
		if value == mValue {
			return mid
		} else if value > mValue {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func BinarySearchRec(array []int, value, start, end int) int {
	if start == end {
		return -1
	} else {
		mid := (start + end) / 2
		mValue := array[mid]
		if value < mValue {
			end = mid
		} else if value > mValue {
			start = mid + 1
		} else {
			return mid
		}
		return BinarySearch(array, value, start, end)
	}
}

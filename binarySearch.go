package learn

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

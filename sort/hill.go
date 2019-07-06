package sort

import "learn"

// 代码简洁， 对于h的选择值有很多种，但是最差情况不会差于插入排序
// 不使用额外空间，时间复杂度 nlogn ~ n^2 -----> 有待于再学习如何计算
func HillSort(array []int) {
	len, h := len(array), 1
	// 生成序列
	for h < len/3 {
		h = h*3 + 1
	}
	for h >= 1 { // log3N
		for i := h; i < len; i++ { // 起点
			// 开始插入排序(实际上我觉得更像冒泡) 我开始想的想法是按照分组，先找子数组，遍历子数组
			// 子数组冒泡排序，前继大于当前就降低
			for j := i; j >= h; j -= h {
				if array[j] < array[j-h] {
					learn.Swap(array, j, j-h)
				}
			}
		}
		h /= 3
	}
}

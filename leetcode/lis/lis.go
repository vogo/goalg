package lis

import "math"

// Lis longest-increasing-subsequence
// 给定长度为n的序列a，从a中抽取出一个子序列，这个子序列需要单调递增。问最长的上升子序列（LIS）的长度。
//　　e.g. 1,5,3,4,6,9,7,8的LIS为1,3,4,6,7,8，长度为6。
func Lis(a []int) int {
	size := len(a)
	f := make([]int, size+1)

	for i := 0; i < size; i++ {
		f[i] = 1
	}

	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			if a[j] < a[i] && f[i] < f[j]+1 {
				f[i] = f[j] + 1
			}
		}
	}
	ans := 0
	for i := 0; i < size; i++ {
		if f[i] > ans {
			ans = f[i]
		}
	}
	return ans
}

// FastLis O(nlogn)
func FastLis(a []int) int {
	size := len(a)

	// 子序列尾部元素的值(最小值)
	// d[2] 表示最长子序列位长度为2的序列的最后一个值中最小值
	// d 是单调递增序列
	d := make([]int, size+1)

	for i := 1; i < size; i++ {
		d[i] = math.MaxInt64
	}

	// 注意: d[0]未使用, 方便后续索引变量计算
	d[1] = a[0]
	length := 1

	for i := 1; i < size; i++ {
		l := 1
		r := length
		mid := 0

		if a[i] > d[length] {
			// 能拼接到后面,则子序列加1
			length++
			d[length] = a[i]
		} else {
			// 二分查找 d[j], 其中 d[j-1] < a[i] < d[j+1]
			// 用 a[i] 替代 d[j]
			for l < r {
				mid = (l + r) / 2
				if d[mid] > a[i] {
					r = mid
				} else {
					l = mid + 1
				}
			}
			if a[i] < d[l] {
				// 替换 最长子序列位长度为 l 的序列的最后一个值中最小值
				// 以便能够匹配更长的子串
				d[l] = a[i]
			}
		}
	}

	return length
}

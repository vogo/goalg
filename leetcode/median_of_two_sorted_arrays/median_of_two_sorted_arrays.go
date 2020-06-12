package median_of_two_sorted_arrays

import "math"

// 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
//
// 请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
// 你可以假设 nums1 和 nums2 不会同时为空。
//
// 示例 1:
//
// nums1 = [1, 3]
// nums2 = [2]
//
// 则中位数是 2.0
// 示例 2:
//
// nums1 = [1, 2]
// nums2 = [3, 4]
//
// 则中位数是 (2 + 3)/2 = 2.5
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	if len1+len2 == 0 {
		return 0
	}

	// 直接定位中位数位置
	maxLeft := (len1 + len2) / 2
	minRight := maxLeft

	if (len1+len2)%2 == 0 {
		maxLeft -= 1
	}

	// 定位左中位数
	left1, left2 := 0, 0
	leftMedian := 0
	for i := 0; i <= maxLeft; i++ {
		if left1 < len1 && left2 < len2 {
			if nums1[left1] < nums2[left2] {
				leftMedian = nums1[left1]
				left1++
			} else {
				leftMedian = nums2[left2]
				left2++
			}
		} else if left1 < len1 {
			leftMedian = nums1[left1]
			left1++
		} else {
			leftMedian = nums2[left2]
			left2++
		}
	}

	// 左右中位数是同一个
	if maxLeft == minRight {
		return float64(leftMedian)
	}

	// 右中位数是挨着左中位数最小者
	rightMedian := math.MaxInt64
	if left1 < len1 {
		rightMedian = nums1[left1]
	}
	if left2 < len2 && nums2[left2] < rightMedian {
		rightMedian = nums2[left2]
	}

	return float64(leftMedian+rightMedian) / 2
}

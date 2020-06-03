// Copyright 2020 wongoo@apache.org. All rights reserved.

// 容器盛水问题
//
// 给定一个整形数组arr，已知其中所有的值都是非负的，将这个数组看作一个容器，请返回容器能装多少水。
//
// 具体请参考样例解释
// 输入描述
// 第一行一个整数N，表示数组长度。
//
// 接下来一行N个数表示数组内的数。
// 输出描述
// 输出一个整数表示能装多少水。
// 示例1
// 输入
// 6
// 3 1 2 5 2 4
// 输出
// 5
//
//       |
//       | * |
// | * * | * |
// | * | | | |
// | | | | | |
//
// 示例2
// 输入
// 5
// 4 5 1 3 2
// 输出
// 2

package trapping_rain_water

// CapacityByShortBoard calc capacity of number bucket by finding the short board in only one loop
//
// 分析: 仔细分析，根据短板理论, 能盛多少水由最短的一块决定。只要找到短板，和短板比较就知道一个位置的盛水容量。
// 这个桶可以看成是许多小桶拼在一起的大桶，算出每一个小桶的容量加总即可。
// 每个小桶，只有左右两块板
// 找出小桶的短板，短板是小桶左侧或右侧最高的板
func CapacityByShortBoard(arr []int) int {
	// 取第一个为左侧板
	leftMax := arr[0]

	// 取最后一个为右侧板
	rightMax := arr[len(arr)-1]

	capacity := 0

	left := 1
	right := len(arr) - 2

	// 左右指针向中间聚合
	for left <= right {
		// 左侧为短板
		if leftMax < rightMax {
			if arr[left] > leftMax {
				// 用更高的板替换左侧短板
				leftMax = arr[left]
			} else {
				// 计算小桶容量
				capacity += leftMax - arr[left]
			}

			// 向中间聚合，计算下一个小桶
			left++
		} else {
			// 右侧为短板
			if arr[right] > rightMax {
				// 用更高的板替换右侧短板
				rightMax = arr[right]
			} else {
				// 计算小桶容量
				capacity += rightMax - arr[right]
			}

			// 向中间聚合，计算下一个小桶
			right--
		}
	}

	return capacity
}

// CapacityByCutOneByOne calc capacity of number bucket by cut floor one by one
//
// 分析: 对于数组 3 1 2 5 2 4，形如下图，星号为可以盛水的地方，一眼看上去特别像俄罗斯方块，于是想到一行一行消除的方式。
// 1. 首先找到最小和最大的数；最小数及其之下的数都是方块，没有空间；最小和最大数之间是有空间的；
// 2. 每一行，左侧第一个方块和右侧第一个方块之间，统计空方块的数量；
// 3. 左侧第一个方块和右侧第一个方块的位置是向中间聚合的；
//       |
//       | * |
// | * * | * |
// | * | | | |
// | | | | | |
//
// Note: 此方法性能比 CapacityByShortBoard 要差，性能跟桶深度有关系，保留此算法只是为了进行对比。
func CapacityByCutOneByOne(arr []int) int {
	min := arr[0]
	max := min

	// 找出最大最小 O(n)
	for _, n := range arr {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	capacity := 0

	// 左右方块位置
	left, right := 0, len(arr)-1

	// 遍历最大最小之间的行, 复杂度看具体数字大小
	for x := min + 1; x <= max; x++ {
		// 找到最左侧方块位置 O(n)
		for i := left; i <= right; i++ {
			left = i
			if arr[i] >= x {
				break
			}
		}

		// 已经和和右侧方块碰撞, 停止
		if left == right {
			break
		}

		// 找到最右侧方块位置 O(n)
		for i := right; i >= left; i-- {
			right = i
			if arr[i] >= x {
				break
			}
		}

		// 已经和和左侧方块碰撞, 停止
		if right == left {
			break
		}

		// 计算第x层的空格  O(n)
		for i := left; i < right; i++ {
			if arr[i] < x {
				capacity++
			}
		}
	}

	return capacity
}

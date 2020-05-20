/*
容器盛水问题

给定一个整形数组arr，已知其中所有的值都是非负的，将这个数组看作一个容器，请返回容器能装多少水。

具体请参考样例解释
输入描述
第一行一个整数N，表示数组长度。

接下来一行N个数表示数组内的数。
输出描述
输出一个整数表示能装多少水。
示例1
输入
6
3 1 2 5 2 4
输出
5
说明

示例2
输入
5
4 5 1 3 2
输出
2
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	arr := make([]int, len(args))
	for i, n := range args {
		arr[i], _ = strconv.Atoi(n)
	}

	leftMax := arr[0]
	rightMax := arr[len(arr)-1]
	count := 0

	left := 1
	right := (len(arr) - 2)
	for left <= right {
		if arr[left] < arr[right] {
			if arr[left] > leftMax {
				leftMax = arr[left]
			} else {
				count += (leftMax - arr[left])
			}
			left++
		} else {
			if arr[right] > rightMax {
				rightMax = arr[right]
			} else {
				count += (rightMax - arr[right])
			}

			right--
		}
	}

	fmt.Println("total:", count)
}

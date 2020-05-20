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
	min := arr[0]
	max := min
	for _, n := range arr {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("min =", min, ",max =", max)

	from := 0
	to := len(arr) - 1
	count := 0

	for x := min + 1; x <= max; x++ {
		start, end := -1, -1
		for i := from; i < to; i++ {
			if arr[i] >= x {
				start = i
				break
			}
		}

		if start == -1 {
			break
		}

		for i := to; i > start; i-- {
			if arr[i] >= x {
				end = i
				break
			}
		}

		if end == -1 {
			break
		}

		fmt.Printf("count empty from %d to %d on floor %d\n", start, end, x)

		for i := start; i < end; i++ {
			if arr[i] < x {
				count++
			}
		}

		from, to = start, end
	}

	fmt.Println("total:", count)
}

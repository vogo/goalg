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

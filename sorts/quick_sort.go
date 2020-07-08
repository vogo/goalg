package sorts

import (
	"sort"
)

// QuickSortInts
func QuickSortInts(a []int) []int {
	QuickSort(sort.IntSlice(a))
	return a
}

// QuickSort
func QuickSort(a sort.Interface) {
	QuickRangeSort(a, 0, a.Len()-1)
}

// QuickRangeSort quick sort for range index from l to r
func QuickRangeSort(a sort.Interface, l, r int) {
	if l >= r {
		return
	}

	i, j, x := l-1, r+1, int(uint(l+r)>>1)

	for {
		for i++; a.Less(i, x); i++ {
		}

		for j--; a.Less(x, j); j-- {
		}

		if i >= j {
			break
		}

		a.Swap(i, j)

		if i == x {
			x = j
		} else if j == x {
			x = i
		}
	}

	QuickRangeSort(a, l, j)
	QuickRangeSort(a, j+1, r)
}

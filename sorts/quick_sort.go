package sorts

import "sort"

// QuickSort
func QuickSort(a sort.Interface) {
	QuickRangeSort(a, 0, a.Len()-1)
}

// QuickRangeSort quick sort for range index from l to r
func QuickRangeSort(a sort.Interface, l, r int) {
	if l >= r {
		return
	}

	i, j := l+1, r

	for {

		for i < r && a.Less(i, l) {
			i++
		}

		for j > l && a.Less(l, j) {
			j--
		}

		if i >= j {
			break
		}

		a.Swap(i, j)
		i++
		j--
	}

	a.Swap(l, j)

	QuickRangeSort(a, l, j-1)
	QuickRangeSort(a, j+1, r)
}

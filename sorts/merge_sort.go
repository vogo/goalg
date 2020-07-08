package sorts

import (
	"github.com/vogo/goalg/compare"
)

func MergeSort(a compare.Array) {
	temp := a.Clone()
	MergeSortRange(a, temp, 0, a.Len()-1)
}

func MergeSortRange(a, temp compare.Array, l int, r int) {
	if l >= r {
		return
	}

	mid := int(uint(l+r) >> 1)

	MergeSortRange(a, temp, l, mid)
	MergeSortRange(a, temp, mid+1, r)

	i, j, x := l, mid+1, l
	for i <= mid && j <= r {
		// less or equal
		for ; i <= mid && (a.Less(i, j) || !a.Less(j, i)); i++ {
			temp.Set(x, a.Get(i))
			x++
		}

		for ; j <= r && a.Less(j, i); j++ {
			temp.Set(x, a.Get(j))
			x++
		}
	}

	if i <= mid {
		temp.CopyFrom(x, a.Sub(i, mid+1))
		x += mid - i + 1
	}

	if j <= r {
		temp.CopyFrom(x, a.Sub(j, r+1))
	}

	a.CopyFrom(l, temp.Sub(l, r+1))
}

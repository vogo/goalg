package sorts

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/vogo/goalg/compare"
)

var (
	benchmarkIntSize  = 10000
	benchmarkIntSlice = rand.Perm(benchmarkIntSize)
)

func BenchmarkGolangSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make([]int, benchmarkIntSize)
		copy(a, benchmarkIntSlice)
		sort.Ints(a)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make([]int, benchmarkIntSize)
		copy(a, benchmarkIntSlice)
		QuickSort(sort.IntSlice(a))
	}
}

func BenchmarkMergeSort(b *testing.B) {
	temp := compare.IntArray(benchmarkIntSlice).Clone()
	for i := 0; i < b.N; i++ {
		a := make([]int, benchmarkIntSize)
		copy(a, benchmarkIntSlice)
		MergeSortRange(compare.IntArray(a), temp, 0, benchmarkIntSize-1)
	}
}

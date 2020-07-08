package sorts

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, QuickSortInts([]int{4, 3, 9, 1, 7, 0, 6, 2, 5, 8}))
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, QuickSortInts([]int{7, 6, 3, 0, 1, 5, 9, 8, 4, 2}))

	assert.Equal(t, []int{1, 2, 2, 3, 3, 4, 5, 6}, QuickSortInts([]int{3, 2, 1, 6, 5, 4, 3, 2}))

	for i := 0; i < 10; i++ {
		size := 10000
		a1 := rand.Perm(size)
		a2 := make([]int, size)
		copy(a2, a1)
		sort.Ints(a1)
		QuickSort(sort.IntSlice(a2))
		assert.Equal(t, a1, a2)
	}
}

var (
	benchmarkIntSize  = 10000
	benchmarkIntSlice = rand.Perm(benchmarkIntSize)
)

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make([]int, benchmarkIntSize)
		copy(a, benchmarkIntSlice)
		QuickSort(sort.IntSlice(benchmarkIntSlice))
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make([]int, benchmarkIntSize)
		copy(a, benchmarkIntSlice)
		sort.Ints(a)
	}
}

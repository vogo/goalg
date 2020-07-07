package sorts

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	a := sort.IntSlice([]int{3, 2, 1, 6, 5, 4})
	QuickSort(a)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, []int(a))
}

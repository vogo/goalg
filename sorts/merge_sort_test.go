package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/goalg/compare"
)

func TestMergeSort(t *testing.T) {
	a := []int{3, 6, 7, 0, 2, 1, 4, 5}
	MergeSort(compare.IntArray(a))
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7}, a)
}

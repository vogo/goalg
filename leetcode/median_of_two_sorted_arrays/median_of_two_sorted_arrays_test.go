package median_of_two_sorted_arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	assert.Equal(t, 8.0, FindMedianSortedArrays([]int{7, 8}, []int{9}))
	assert.Equal(t, -1.0, FindMedianSortedArrays([]int{3}, []int{-2, -1}))
	assert.Equal(t, 2.0, FindMedianSortedArrays([]int{}, []int{2}))
	assert.Equal(t, 3.0, FindMedianSortedArrays([]int{3}, []int{}))
	assert.Equal(t, 2.5, FindMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, 4.0, FindMedianSortedArrays([]int{1}, []int{4, 5}))
	assert.Equal(t, 3.5, FindMedianSortedArrays([]int{1, 2, 5}, []int{3, 4, 7}))
	assert.Equal(t, 3.5, FindMedianSortedArrays([]int{1, 2, 3}, []int{4, 5, 6}))
	assert.Equal(t, 4.0, FindMedianSortedArrays([]int{1, 2, 3}, []int{4, 5, 6, 7}))
	assert.Equal(t, 4.5, FindMedianSortedArrays([]int{1, 2, 3}, []int{4, 5, 6, 7, 8}))
}

// author: wongoo@apache.org

package searches

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/goalg/compare"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, 0, BinarySearch(compare.NewComparers([]int{1, 2, 3, 4}), compare.Int(1)))
	assert.Equal(t, 1, BinarySearch(compare.NewComparers([]int{1, 2, 3, 4}), compare.Int(2)))
	assert.Equal(t, 3, BinarySearch(compare.NewComparers([]int{1, 2, 3, 4}), compare.Int(4)))
	assert.Equal(t, -1, BinarySearch(compare.NewComparers([]int{1, 2, 3, 4}), compare.Int(5)))
	assert.Equal(t, 2, BinarySearch(compare.NewComparers([]int{1, 2, 3, 4, 5}), compare.Int(3)))
	assert.Equal(t, 3, BinarySearch(compare.NewComparers([]int{1, 2, 3, 4, 5}), compare.Int(4)))

	assert.Equal(t, 0, BinarySearch(compare.NewComparers([]int{1}), compare.Int(1)))
	assert.Equal(t, -1, BinarySearch(compare.NewComparers([]int{}), compare.Int(1)))
}

func TestLeftBound(t *testing.T) {
	assert.Equal(t, 0, BinarySearchLeftBound(compare.NewLessers([]int{1, 2, 2, 3, 4}), compare.Int(1)))
	assert.Equal(t, 1, BinarySearchLeftBound(compare.NewLessers([]int{1, 2, 2, 3, 4}), compare.Int(2)))
	assert.Equal(t, 1, BinarySearchLeftBound(compare.NewLessers([]int{1, 2, 2, 2, 3, 4}), compare.Int(2)))
	assert.Equal(t, 3, BinarySearchLeftBound(compare.NewLessers([]int{1, 2, 3, 4, 4}), compare.Int(4)))
	assert.Equal(t, -1, BinarySearchLeftBound(compare.NewLessers([]int{1, 2, 3, 4}), compare.Int(5)))
	assert.Equal(t, 2, BinarySearchLeftBound(compare.NewLessers([]int{1, 2, 3, 3, 4, 5}), compare.Int(3)))
	assert.Equal(t, 3, BinarySearchLeftBound(compare.NewLessers([]int{1, 2, 3, 4, 5}), compare.Int(4)))

	assert.Equal(t, 0, BinarySearchLeftBound(compare.NewLessers([]int{1}), compare.Int(1)))
	assert.Equal(t, 0, BinarySearchLeftBound(compare.NewLessers([]int{1, 1}), compare.Int(1)))
	assert.Equal(t, -1, BinarySearchLeftBound(compare.NewLessers([]int{}), compare.Int(1)))
}

func TestRightBound(t *testing.T) {
	assert.Equal(t, 0, BinarySearchRightBound(compare.NewLessers([]int{1, 2, 2, 3, 4}), compare.Int(1)))
	assert.Equal(t, 2, BinarySearchRightBound(compare.NewLessers([]int{1, 2, 2, 3, 4}), compare.Int(2)))
	assert.Equal(t, 3, BinarySearchRightBound(compare.NewLessers([]int{1, 2, 2, 2, 3, 4}), compare.Int(2)))
	assert.Equal(t, 4, BinarySearchRightBound(compare.NewLessers([]int{1, 2, 3, 4, 4}), compare.Int(4)))
	assert.Equal(t, -1, BinarySearchRightBound(compare.NewLessers([]int{1, 2, 3, 4}), compare.Int(5)))
	assert.Equal(t, 3, BinarySearchRightBound(compare.NewLessers([]int{1, 2, 3, 3, 4, 5}), compare.Int(3)))
	assert.Equal(t, 3, BinarySearchRightBound(compare.NewLessers([]int{1, 2, 3, 4, 5}), compare.Int(4)))

	assert.Equal(t, 0, BinarySearchRightBound(compare.NewLessers([]int{1}), compare.Int(1)))
	assert.Equal(t, 1, BinarySearchRightBound(compare.NewLessers([]int{1, 1}), compare.Int(1)))
	assert.Equal(t, -1, BinarySearchRightBound(compare.NewLessers([]int{}), compare.Int(1)))
}

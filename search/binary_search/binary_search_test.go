// author: wongoo@apache.org

package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/goalg/compare"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, 0, Search(compare.NewComparers([]int{1, 2, 3, 4}), compare.Int(1)))
	assert.Equal(t, 1, Search(compare.NewComparers([]int{1, 2, 3, 4}), compare.Int(2)))
	assert.Equal(t, 3, Search(compare.NewComparers([]int{1, 2, 3, 4}), compare.Int(4)))
	assert.Equal(t, -1, Search(compare.NewComparers([]int{1, 2, 3, 4}), compare.Int(5)))
	assert.Equal(t, 2, Search(compare.NewComparers([]int{1, 2, 3, 4, 5}), compare.Int(3)))
	assert.Equal(t, 3, Search(compare.NewComparers([]int{1, 2, 3, 4, 5}), compare.Int(4)))

	assert.Equal(t, 0, Search(compare.NewComparers([]int{1}), compare.Int(1)))
	assert.Equal(t, -1, Search(compare.NewComparers([]int{}), compare.Int(1)))
}

func TestLeftBound(t *testing.T) {
	assert.Equal(t, 0, LeftBound(compare.NewLessers([]int{1, 2, 2, 3, 4}), compare.Int(1)))
	assert.Equal(t, 1, LeftBound(compare.NewLessers([]int{1, 2, 2, 3, 4}), compare.Int(2)))
	assert.Equal(t, 1, LeftBound(compare.NewLessers([]int{1, 2, 2, 2, 3, 4}), compare.Int(2)))
	assert.Equal(t, 3, LeftBound(compare.NewLessers([]int{1, 2, 3, 4, 4}), compare.Int(4)))
	assert.Equal(t, -1, LeftBound(compare.NewLessers([]int{1, 2, 3, 4}), compare.Int(5)))
	assert.Equal(t, 2, LeftBound(compare.NewLessers([]int{1, 2, 3, 3, 4, 5}), compare.Int(3)))
	assert.Equal(t, 3, LeftBound(compare.NewLessers([]int{1, 2, 3, 4, 5}), compare.Int(4)))

	assert.Equal(t, 0, LeftBound(compare.NewLessers([]int{1}), compare.Int(1)))
	assert.Equal(t, 0, LeftBound(compare.NewLessers([]int{1, 1}), compare.Int(1)))
	assert.Equal(t, -1, LeftBound(compare.NewLessers([]int{}), compare.Int(1)))
}

func TestRightBound(t *testing.T) {
	assert.Equal(t, 0, RightBound(compare.NewLessers([]int{1, 2, 2, 3, 4}), compare.Int(1)))
	assert.Equal(t, 2, RightBound(compare.NewLessers([]int{1, 2, 2, 3, 4}), compare.Int(2)))
	assert.Equal(t, 3, RightBound(compare.NewLessers([]int{1, 2, 2, 2, 3, 4}), compare.Int(2)))
	assert.Equal(t, 4, RightBound(compare.NewLessers([]int{1, 2, 3, 4, 4}), compare.Int(4)))
	assert.Equal(t, -1, RightBound(compare.NewLessers([]int{1, 2, 3, 4}), compare.Int(5)))
	assert.Equal(t, 3, RightBound(compare.NewLessers([]int{1, 2, 3, 3, 4, 5}), compare.Int(3)))
	assert.Equal(t, 3, RightBound(compare.NewLessers([]int{1, 2, 3, 4, 5}), compare.Int(4)))

	assert.Equal(t, 0, RightBound(compare.NewLessers([]int{1}), compare.Int(1)))
	assert.Equal(t, 1, RightBound(compare.NewLessers([]int{1, 1}), compare.Int(1)))
	assert.Equal(t, -1, RightBound(compare.NewLessers([]int{}), compare.Int(1)))
}

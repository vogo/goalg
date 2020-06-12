package lis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLis(t *testing.T) {
	testLis(t, 5, []int{6, 7, 8, 9, 1, 2, 3, 4, 5})
	testLis(t, 6, []int{1, 5, 3, 4, 6, 9, 7, 8})

}

func testLis(t *testing.T, l int, ints []int) {
	assert.Equal(t, l, Lis(ints))
	assert.Equal(t, l, FastLis(ints))
}

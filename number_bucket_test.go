package goalg_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wongoo/goalg"
)

func TestCountCapacityByCutOneByOne(t *testing.T) {
	arr1 := []int{3, 1, 2, 5, 2, 4}
	capacity1 := goalg.CountCapacityByCutOneByOne(arr1)
	capacity2 := goalg.CountCapacityByShortBoard(arr1)
	assert.Equal(t, 5, capacity1)
	assert.Equal(t, 5, capacity2)

	arr2 := []int{4, 5, 1, 3, 2}
	capacity1 = goalg.CountCapacityByCutOneByOne(arr2)
	capacity2 = goalg.CountCapacityByShortBoard(arr2)
	assert.Equal(t, 2, capacity1)
	assert.Equal(t, 2, capacity2)

	capacity1 = goalg.CountCapacityByCutOneByOne(goalg.NumberBucketIntArr10)
	capacity2 = goalg.CountCapacityByShortBoard(goalg.NumberBucketIntArr10)
	assert.Equal(t, capacity1, capacity2)

	capacity1 = goalg.CountCapacityByCutOneByOne(goalg.NumberBucketIntArr50)
	capacity2 = goalg.CountCapacityByShortBoard(goalg.NumberBucketIntArr50)
	assert.Equal(t, capacity1, capacity2)

	capacity1 = goalg.CountCapacityByCutOneByOne(goalg.NumberBucketIntArr100)
	capacity2 = goalg.CountCapacityByShortBoard(goalg.NumberBucketIntArr100)
	assert.Equal(t, capacity1, capacity2)
}

func BenchmarkCountCapacityByCutOneByOne10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goalg.CountCapacityByCutOneByOne(goalg.NumberBucketIntArr10)
	}
}

func BenchmarkCountCapacityByCutOneByOne50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goalg.CountCapacityByCutOneByOne(goalg.NumberBucketIntArr50)
	}
}

func BenchmarkCountCapacityByCutOneByOne100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goalg.CountCapacityByCutOneByOne(goalg.NumberBucketIntArr100)
	}
}

func BenchmarkCountCapacityByShortBoard10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goalg.CountCapacityByShortBoard(goalg.NumberBucketIntArr10)
	}
}

func BenchmarkCountCapacityByShortBoard50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goalg.CountCapacityByShortBoard(goalg.NumberBucketIntArr50)
	}
}

func BenchmarkCountCapacityByShortBoard100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goalg.CountCapacityByShortBoard(goalg.NumberBucketIntArr100)
	}
}

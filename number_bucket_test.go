package goalg_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wongoo/goalg"
)

func TestNumberBucketCapacity(t *testing.T) {
	arr1 := []int{3, 1, 2, 5, 2, 4}
	capacity1 := goalg.NumberBucketCapacityByCutOneByOne(arr1)
	capacity2 := goalg.NumberBucketCapacityByShortBoard(arr1)
	assert.Equal(t, 5, capacity1)
	assert.Equal(t, 5, capacity2)

	arr2 := []int{4, 5, 1, 3, 2}
	capacity1 = goalg.NumberBucketCapacityByCutOneByOne(arr2)
	capacity2 = goalg.NumberBucketCapacityByShortBoard(arr2)
	assert.Equal(t, 2, capacity1)
	assert.Equal(t, 2, capacity2)

	capacity1 = goalg.NumberBucketCapacityByCutOneByOne(goalg.NumberBucketIntArr10)
	capacity2 = goalg.NumberBucketCapacityByShortBoard(goalg.NumberBucketIntArr10)
	assert.Equal(t, capacity1, capacity2)

	capacity1 = goalg.NumberBucketCapacityByCutOneByOne(goalg.NumberBucketIntArr50)
	capacity2 = goalg.NumberBucketCapacityByShortBoard(goalg.NumberBucketIntArr50)
	assert.Equal(t, capacity1, capacity2)

	capacity1 = goalg.NumberBucketCapacityByCutOneByOne(goalg.NumberBucketIntArr100)
	capacity2 = goalg.NumberBucketCapacityByShortBoard(goalg.NumberBucketIntArr100)
	assert.Equal(t, capacity1, capacity2)
}

func BenchmarkNumberBucketCapacityByCutOneByOne10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goalg.NumberBucketCapacityByCutOneByOne(goalg.NumberBucketIntArr10)
	}
}

func BenchmarkNumberBucketCapacityByCutOneByOne50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goalg.NumberBucketCapacityByCutOneByOne(goalg.NumberBucketIntArr50)
	}
}

func BenchmarkNumberBucketCapacityByCutOneByOne100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goalg.NumberBucketCapacityByCutOneByOne(goalg.NumberBucketIntArr100)
	}
}

func BenchmarkNumberBucketCapacityByShortBoard10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goalg.NumberBucketCapacityByShortBoard(goalg.NumberBucketIntArr10)
	}
}

func BenchmarkNumberBucketCapacityByShortBoard50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goalg.NumberBucketCapacityByShortBoard(goalg.NumberBucketIntArr50)
	}
}

func BenchmarkNumberBucketCapacityByShortBoard100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goalg.NumberBucketCapacityByShortBoard(goalg.NumberBucketIntArr100)
	}
}

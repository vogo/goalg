package numbucket_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wongoo/goalg/numbucket"
)

func randIntArr(size, max int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(max)
	}
	return arr
}

var (
	IntArr10  = randIntArr(10, 10)
	IntArr50  = randIntArr(50, 10)
	IntArr100 = randIntArr(100, 10)
)

func TestCapacity(t *testing.T) {
	arr1 := []int{3, 1, 2, 5, 2, 4}
	capacity1 := numbucket.CapacityByCutOneByOne(arr1)
	capacity2 := numbucket.CapacityByShortBoard(arr1)
	assert.Equal(t, 5, capacity1)
	assert.Equal(t, 5, capacity2)

	arr2 := []int{4, 5, 1, 3, 2}
	capacity1 = numbucket.CapacityByCutOneByOne(arr2)
	capacity2 = numbucket.CapacityByShortBoard(arr2)
	assert.Equal(t, 2, capacity1)
	assert.Equal(t, 2, capacity2)

	capacity1 = numbucket.CapacityByCutOneByOne(IntArr10)
	capacity2 = numbucket.CapacityByShortBoard(IntArr10)
	assert.Equal(t, capacity1, capacity2)

	capacity1 = numbucket.CapacityByCutOneByOne(IntArr50)
	capacity2 = numbucket.CapacityByShortBoard(IntArr50)
	assert.Equal(t, capacity1, capacity2)

	capacity1 = numbucket.CapacityByCutOneByOne(IntArr100)
	capacity2 = numbucket.CapacityByShortBoard(IntArr100)
	assert.Equal(t, capacity1, capacity2)
}

func BenchmarkCapacityByCutOneByOne10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbucket.CapacityByCutOneByOne(IntArr10)
	}
}

func BenchmarkCapacityByCutOneByOne50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbucket.CapacityByCutOneByOne(IntArr50)
	}
}

func BenchmarkCapacityByCutOneByOne100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbucket.CapacityByCutOneByOne(IntArr100)
	}
}

func BenchmarkCapacityByShortBoard10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbucket.CapacityByShortBoard(IntArr10)
	}
}

func BenchmarkCapacityByShortBoard50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbucket.CapacityByShortBoard(IntArr50)
	}
}

func BenchmarkCapacityByShortBoard100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbucket.CapacityByShortBoard(IntArr100)
	}
}

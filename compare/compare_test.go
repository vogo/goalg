package compare

import (
	"testing"
)

type CompareItem struct {
	key int
}

func (i *CompareItem) Less(o *CompareItem) bool {
	return i.key < o.key
}

func (i *CompareItem) Compare(o *CompareItem) int {
	if i.key < o.key {
		return -1
	}

	if i.key > o.key {
		return 1
	}

	return 0
}

var (
	itemA = &CompareItem{
		key: 12345,
	}
	itemB = itemA
)

func BenchmarkLess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch {
		case itemA.Less(itemB):
			break
		case itemB.Less(itemA):
			break
		default:
		}
	}
}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch itemA.Compare(itemB) {
		case -1:
			break
		case 1:
			break
		default:
		}
	}
}

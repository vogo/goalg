package compare

import "sort"

// Lesser compare one is less than another
type Lesser interface {
	// Less return true if current less than another,
	// l is equal to o if !l.Less(o) && !o.Less(l)
	Less(o Lesser) bool
}

// Comparer compare one with another
type Comparer interface {
	// Compare compare with another value
	// suggest:
	// - return 1 if current is greater than another,
	// - return -1 if less than another,
	// - return 0 if equal.
	Compare(o Comparer) int
}

type Int int

func (i Int) Less(o Lesser) bool {
	return i < o.(Int)
}

func (i Int) Compare(o Comparer) int {
	if i < o.(Int) {
		return -1
	}

	if i > o.(Int) {
		return 1
	}

	return 0
}

func NewComparers(a []int) []Comparer {
	n := len(a)
	if n == 0 {
		return nil
	}
	arr := make([]Comparer, n)
	for i := 0; i < n; i++ {
		arr[i] = Int(a[i])
	}
	return arr
}

func NewLessers(a []int) []Lesser {
	n := len(a)
	if n == 0 {
		return nil
	}
	arr := make([]Lesser, n)
	for i := 0; i < n; i++ {
		arr[i] = Int(a[i])
	}
	return arr
}

type Array interface {
	sort.Interface
	Clone() Array
	Set(i int, v interface{})
	Get(i int) interface{}
	Sub(i, j int) interface{}
	CopyFrom(start int, src interface{})
}

type IntArray []int

func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p IntArray) Clone() Array {
	b := make([]int, len(p))
	copy(b[:], p[:])
	return IntArray(b)
}

func (p IntArray) Set(i int, v interface{}) {
	p[i] = v.(int)
}

func (p IntArray) Get(i int) interface{} {
	return p[i]
}

func (p IntArray) Sub(i, j int) interface{} {
	return p[i:j]
}

func (p IntArray) CopyFrom(start int, src interface{}) {
	copy(p[start:], src.(IntArray))
}

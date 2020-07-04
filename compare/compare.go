package compare

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

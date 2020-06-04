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

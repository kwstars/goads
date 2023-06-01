package common

// Comparator is a function that compares two elements.
// It returns a negative number if a < b, zero if a == b, and a positive number if a > b.
type Comparator[a, b any] func(a, b) int8

var (
	IntComparator = func(a, b int) int8 {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}
	StringComparator = func(a, b string) int8 {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}
)

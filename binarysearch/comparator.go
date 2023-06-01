package binarysearch

// Comparator is a function that compares two elements.
// It returns a negative number if a < b, zero if a == b, and a positive number if a > b.
// a parameter is any type, b parameter is target value type.
type Comparator[a any, target comparable] func(a, target) int8

// Range is a struct that represents a range of integers. for test.
type Range struct {
	Start int
	End   int
}

var (
	rangeComparator = func(a Range, b int) int8 {
		if a.Start <= b && b <= a.End {
			return 0
		} else if a.Start > b {
			return 1
		} else {
			return -1
		}
	}
	IntComparator = func(a, b int) int8 {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}
)

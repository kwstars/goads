package lists

type List[T any] interface {
	// Add inserts an element at a specific position in the list.
	// It returns an error if the operation is unsuccessful.
	Add(index int, element T) error

	// Append adds an element to the end of the list.
	Append(element T)

	// AddAll inserts multiple elements at a specific position in the list.
	// It returns an error if the operation is unsuccessful.
	AddAll(index int, elements []T) error

	// Clear removes all elements from the list.
	Clear()

	// Equals checks if the list is equal to another.
	Equals(other T) bool

	// Get retrieves an element at a specific position in the list.
	Get(index int) (T, error)

	// HashCode computes a hash code for the list.
	HashCode() int

	// IndexOf finds the first occurrence of an element in the list, returning its index.
	// If the element is not present, it returns -1.
	IndexOf(element T) int

	// Iterator provides an iterator over the elements in the list in sequence.
	Iterator() Iterator

	// LastIndexOf finds the last occurrence of an element in the list, returning its index.
	// If the element is not present, it returns -1.
	LastIndexOf(element T) int

	// ListIterator provides a list iterator over the elements in the list in sequence.
	ListIterator() ListIterator

	// ListIteratorFrom provides a list iterator over the elements in the list in sequence, starting at a specific position.
	ListIteratorFrom(index int) ListIterator

	// Remove deletes an element at a specific position in the list.
	// It returns an error if the operation is unsuccessful.
	Remove(index int) error

	// RemoveRange deletes a range of elements from the list, between two indices.
	// It returns an error if the operation is unsuccessful.
	RemoveRange(fromIndex int, toIndex int) error

	// Set replaces an element at a specific position in the list with another element.
	// It returns an error if the operation is unsuccessful.
	Set(index int, element T) error

	// SubList returns a subsection of the list, between two indices.
	SubList(fromIndex int, toIndex int) []T
}

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
}

type ListIterator interface {
	Iterator
	HasPrevious() bool
	Previous() (interface{}, error)
}

package arraylist

import (
	"errors"
	"fmt"
	"github.com/kwstars/goads/lists"
	"sort"
)

var (
	ErrIndexOutOfRange                = errors.New("index out of range")
	ErrFormIndexMustBeLessThanToIndex = errors.New("fromIndex must be less than or equal to toIndex")
	ErrEmptyList                      = errors.New("list is empty")
)

var _ lists.List[int] = (*List[int])(nil)

// List is a struct that holds the elements of the list.
type List[T any] struct {
	elements []T
	cmp      func(a, b T) int8
}

// Option is a function that configures a list.
type Option[T any] func() []T

// WithInitialCapacity sets the initial capacity of the list.
func WithInitialCapacity[T any](capacity int) Option[T] {
	return func() []T {
		return make([]T, 0, capacity)
	}
}

// New creates a new list.
func New[T any](cmp func(a, b T) int8, options ...Option[T]) *List[T] {
	var elements []T
	for _, option := range options {
		elements = option()
	}

	// If elements is still nil after applying options, initialize it with default capacity
	if elements == nil {
		elements = make([]T, 0)
	}

	return &List[T]{
		elements: elements,
		cmp:      cmp,
	}
}

// Empty returns true if the list is empty.
func (l *List[T]) Empty() bool {
	return len(l.elements) == 0
}

// Full returns true if the list is full.
func (l *List[T]) Full() bool {
	return false
}

// Size returns the number of elements in the list.
func (l *List[T]) Size() int {
	return len(l.elements)
}

// Append adds an element to the end of the list.
func (l *List[T]) Append(element T) {
	l.elements = append(l.elements, element)
}

// Prepend adds an element to the front of the list.
func (l *List[T]) Prepend(element T) {
	l.elements = append([]T{element}, l.elements...)
}

// Insert inserts an element at the specified index.
func (l *List[T]) Insert(index int, element T) error {
	if index < 0 || index > len(l.elements) {
		return fmt.Errorf("%w: %d", ErrIndexOutOfRange, index)
	}

	l.elements = append(l.elements[:index], append([]T{element}, l.elements[index:]...)...)

	return nil
}

// InsertAll inserts multiple elements at a specific position in the list.
func (l *List[T]) InsertAll(index int, elements []T) error {
	if index < 0 || index > len(l.elements) {
		return fmt.Errorf("%w: %d", ErrIndexOutOfRange, index)
	}

	l.elements = append(l.elements[:index], append(elements, l.elements[index:]...)...)

	return nil
}

// Get retrieves an element at a specific position in the list.
func (l *List[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(l.elements) {
		var zero = new(T)
		return *zero, fmt.Errorf("%w: %d", ErrIndexOutOfRange, index)
	}

	return l.elements[index], nil
}

// IndexOf finds the first occurrence of an element in the list, returning its index.
func (l *List[T]) IndexOf(element T) int {
	for k, v := range l.elements {
		if l.cmp(v, element) == 0 {
			return k
		}
	}
	return -1
}

// LastIndexOf finds the last occurrence of an element in the list, returning its index.
func (l *List[T]) LastIndexOf(element T) int {
	for i := len(l.elements) - 1; i >= 0; i-- {
		if l.cmp(l.elements[i], element) == 0 {
			return i
		}
	}
	return -1
}

// Remove removes an element at a specific position in the list.
func (l *List[T]) Remove(index int) error {
	if index < 0 || index >= len(l.elements) {
		return fmt.Errorf("%w: %d", ErrIndexOutOfRange, index)
	}

	l.elements = append(l.elements[:index], l.elements[index+1:]...)

	// Alternative: a = a[:i+copy(a[i:], a[i+1:])]
	// This copies all elements after index i and updates a

	return nil
}

// RemoveUnorderedAtIndex removes an element at a specific position in the list.
func (l *List[T]) RemoveUnorderedAtIndex(index int) error {
	if index < 0 || index >= len(l.elements) {
		return fmt.Errorf("%w: %d", ErrIndexOutOfRange, index)
	}

	l.elements[index] = l.elements[len(l.elements)-1]
	l.elements = l.elements[:len(l.elements)-1]

	return nil
}

// PopAtIndex removes the element at the specified position in the list.
func (l *List[T]) PopAtIndex(index int) (T, error) {
	if index < 0 || index >= len(l.elements) {
		zero := new(T)
		return *zero, fmt.Errorf("%w: %d", ErrIndexOutOfRange, index)
	}
	element := l.elements[index]
	l.elements = append(l.elements[:index], l.elements[index+1:]...)
	return element, nil
}

// Pop removes the last element from the list.
func (l *List[T]) Pop() (T, error) {
	if len(l.elements) == 0 {
		zero := new(T)
		return *zero, fmt.Errorf("%w", ErrEmptyList)
	}
	element := l.elements[len(l.elements)-1]
	l.elements = l.elements[:len(l.elements)-1]
	return element, nil
}

// PopFront removes the first element from the list.
func (l *List[T]) PopFront() (T, error) {
	if len(l.elements) == 0 {
		zero := new(T)
		return *zero, fmt.Errorf("%w", ErrEmptyList)
	}
	element := l.elements[0]
	l.elements = l.elements[1:]
	return element, nil
}

// Set replaces the element at the specified position in the list with a new element.
func (l *List[T]) Set(index int, value T) error {
	if index < 0 || index >= len(l.elements) {
		return fmt.Errorf("%w: %d", ErrIndexOutOfRange, index)
	}

	l.elements[index] = value

	return nil
}

// Clear removes all elements from the list.
func (l *List[T]) Clear() {
	l.elements = l.elements[:0]

	//l.elements[:0] clears the slice by truncating it to 0 elements while
	//preserving the underlying array capacity.

	//l.elements[:0:0] clears the slice and releases the underlying array memory.
	//It truncates the slice to 0 elements and 0 capacity.
	//l.elements = l.elements[:0:0]
}

// Contains returns true if the list contains the specified element.
func (l *List[T]) Contains(values ...T) bool {
	for _, searchValue := range values {
		found := false
		for _, element := range l.elements {
			if l.cmp(element, searchValue) == 0 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// RemoveRange removes from this list all the elements whose index is between fromIndex, inclusive, and toIndex, exclusive.
func (l *List[T]) RemoveRange(fromIndex int, toIndex int) error {
	if fromIndex < 0 || fromIndex >= len(l.elements) || toIndex < 0 || toIndex >= len(l.elements) {
		return fmt.Errorf("%w, fromIndex: %d, toIndex: %d, size: %d", ErrIndexOutOfRange, fromIndex, toIndex, len(l.elements))
	}

	if fromIndex > toIndex {
		return fmt.Errorf("%w, fromIndex: %d, toIndex: %d", ErrFormIndexMustBeLessThanToIndex, fromIndex, toIndex)
	}

	l.elements = append(l.elements[:fromIndex], l.elements[toIndex:]...)

	return nil
}

// SubList returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive.
func (l *List[T]) SubList(fromIndex int, toIndex int) ([]T, error) {
	if fromIndex < 0 || fromIndex >= len(l.elements) || toIndex < 0 || toIndex >= len(l.elements) {
		return nil, fmt.Errorf("%w, fromIndex: %d, toIndex: %d, size: %d", ErrIndexOutOfRange, fromIndex, toIndex, len(l.elements))
	}

	if fromIndex > toIndex {
		return nil, fmt.Errorf("%w", ErrFormIndexMustBeLessThanToIndex)
	}

	return l.elements[fromIndex:toIndex], nil
}

// Reverse reverses the order of the elements in the list.
func (l *List[T]) Reverse() {
	length := len(l.elements)
	if length < 2 {
		return
	}
	for left, right := 0, length-1; left < right; left, right = left+1, right-1 {
		l.elements[left], l.elements[right] = l.elements[right], l.elements[left]
	}
	// Alternative
	//for i := length/2 - 1; i >= 0; i-- {
	//	opp := length - 1 - i
	//	l.elements[i], l.elements[opp] = l.elements[opp], l.elements[i]
	//}
}

// RemoveIf removes all elements that match the specified predicate.
func (l *List[T]) RemoveIf(predicate func(T) bool) bool {
	n := 0
	for _, x := range l.elements {
		if !predicate(x) {
			l.elements[n] = x
			n++
		}
	}
	if n < len(l.elements) {
		l.elements = l.elements[:n]
		return true
	}
	return false
}

// Sort sorts the elements of the list.
func (l *List[T]) Sort() {
	sort.Slice(l.elements, func(i, j int) bool {
		return l.cmp(l.elements[i], l.elements[j]) < 0
	})
}

// Copy deep copies the list.
func (l *List[T]) Copy() []T {
	dst := make([]T, len(l.elements))
	copy(dst, l.elements)
	return dst
	// Alternative 1: return append([]T(nil), src...)
	// This creates an empty slice and appends all elements of src, achieving a deep copy.
	//Alternative 2: return append(src[:0:0], src...)
	// This creates an empty slice based on src and appends all elements of src, achieving a deep copy.
}

// BatchSplit splits the list into batches of the specified size.
func (l *List[T]) BatchSplit(batchSize int) [][]T {
	if batchSize <= 0 {
		return [][]T{}
	}

	values := l.elements

	// Calculate the total number of batches. For example, if values length is 10 and batchSize is 3,
	// batchCount is 4 because 10 + 3 - 1 = 12 and 12 / 3 = 4.
	batchCount := (len(values) + batchSize - 1) / batchSize
	batches := make([][]T, 0, batchCount)

	for batchSize < len(values) {
		values, batches = values[batchSize:], append(batches, values[0:batchSize:batchSize])
	}
	batches = append(batches, values)

	return batches
}

// SlidingWindows returns a slice of slices of size where each slice.
func (l *List[T]) SlidingWindows(size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}
	// returns the input slice as the first element
	if len(l.elements) <= size {
		return [][]T{l.elements}
	}

	// allocate slice at the precise size we need
	// For example, if l.elements = [1, 2, 3, 4, 5, 6, 7] and size = 3,
	// then r = make([][]T, 0, 7-3+1) = make([][]T, 0, 5), with capacity of 5.
	r := make([][]T, 0, len(l.elements)-size+1)

	// Get start and end indices of sliding window and append window slice to result slice.
	for i, j := 0, size; j <= len(l.elements); i, j = i+1, j+1 {
		r = append(r, l.elements[i:j])
	}

	return r
}

// BringElementToFront moves the specified element to the front of the ArrayList.
// If the element is already at the front in the list, no action is taken.
// If the element is not in the list, it is added to the front of the list.
// This operation modifies the ArrayList in place whenever possible.
func (l *List[T]) BringElementToFront(element T) {
	if len(l.elements) != 0 && l.cmp(l.elements[0], element) == 0 {
		return
	}
	previousElement := element //Initialize previousElement to the target element
	for i, currentElement := range l.elements {
		switch {
		case i == 0:
			l.elements[0] = element
			previousElement = currentElement
		case l.cmp(currentElement, element) == 0:
			l.elements[i] = previousElement
			return
		default: //Otherwise
			l.elements[i] = previousElement
			previousElement = currentElement
		}
	}
	l.elements = append(l.elements, previousElement)
}

// RemoveDuplicates removes duplicates from the ArrayList.
// Before removing duplicates, it sorts the ArrayList.
func (l *List[T]) RemoveDuplicates() {
	if len(l.elements) <= 1 {
		return
	}

	// Start by sorting the List
	l.Sort()

	slow := 0
	for fast := 1; fast < len(l.elements); fast++ {
		if l.cmp(l.elements[slow], l.elements[fast]) == 0 {
			continue
		}
		slow++
		l.elements[slow] = l.elements[fast]

	}
	// Shrink the size of the List after removing duplicates
	l.elements = l.elements[:slow+1]
}

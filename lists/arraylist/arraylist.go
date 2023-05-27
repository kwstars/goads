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
	ErrPopFromEmptyList               = errors.New("cannot pop from empty list")
)

var _ lists.List[int] = (*ArrayList[int])(nil)

type ArrayList[T any] struct {
	elements []T
	cmp      func(a, b T) int8
}

type Option[T any] func() []T

func WithInitialCapacity[T any](capacity int) Option[T] {
	return func() []T {
		return make([]T, 0, capacity)
	}
}

func NewArrayList[T any](cmp func(a, b T) int8, options ...Option[T]) *ArrayList[T] {
	var elements []T
	for _, option := range options {
		elements = option()
	}

	// If elements is still nil after applying options, initialize it with default capacity
	if elements == nil {
		elements = make([]T, 0)
	}

	return &ArrayList[T]{
		elements: elements,
		cmp:      cmp,
	}
}

// Empty returns true if the list is empty.
func (t *ArrayList[T]) Empty() bool {
	return len(t.elements) == 0
}

// Full returns true if the list is full.
func (t *ArrayList[T]) Full() bool {
	return false
}

// Size returns the number of elements in the list.
func (t *ArrayList[T]) Size() int {
	return len(t.elements)
}

// Add inserts an element at the specified index.
func (t *ArrayList[T]) Add(index int, element T) error {
	if index < 0 || index > len(t.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	t.elements = append(t.elements[:index], append([]T{element}, t.elements[index:]...)...)

	return nil
}

// Append adds an element to the end of the list.
func (t *ArrayList[T]) Append(element T) {
	t.elements = append(t.elements, element)
}

// AppendFront adds an element to the front of the list.
func (t *ArrayList[T]) AppendFront(element T) {
	t.elements = append([]T{element}, t.elements...)
}

// AddAll inserts multiple elements at a specific position in the list.
func (t *ArrayList[T]) AddAll(index int, elements []T) error {
	if index < 0 || index > len(t.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	t.elements = append(t.elements[:index], append(elements, t.elements[index:]...)...)

	return nil
}

// Get retrieves an element at a specific position in the list.
func (t *ArrayList[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(t.elements) {
		var zero = new(T)
		return *zero, fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	return t.elements[index], nil
}

// IndexOf finds the first occurrence of an element in the list, returning its index.
func (t *ArrayList[T]) IndexOf(element T) int {
	for k, v := range t.elements {
		if t.cmp(v, element) == 0 {
			return k
		}
	}
	return -1
}

// LastIndexOf finds the last occurrence of an element in the list, returning its index.
func (t *ArrayList[T]) LastIndexOf(element T) int {
	for i := len(t.elements) - 1; i >= 0; i-- {
		if t.cmp(t.elements[i], element) == 0 {
			return i
		}
	}
	return -1
}

// Remove removes an element at a specific position in the list.
func (t *ArrayList[T]) Remove(index int) error {
	if index < 0 || index >= len(t.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	t.elements = append(t.elements[:index], t.elements[index+1:]...)

	// Alternative: a = a[:i+copy(a[i:], a[i+1:])]
	// This copies all elements after index i and updates a

	return nil
}

// RemoveUnorderedAtIndex removes an element at a specific position in the list.
func (t *ArrayList[T]) RemoveUnorderedAtIndex(index int) error {
	if index < 0 || index >= len(t.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	t.elements[index] = t.elements[len(t.elements)-1]
	t.elements = t.elements[:len(t.elements)-1]

	return nil
}

// PopAtIndex removes the element at the specified position in the list.
func (t *ArrayList[T]) PopAtIndex(index int) (T, error) {
	if index < 0 || index >= len(t.elements) {
		zero := new(T)
		return *zero, fmt.Errorf("%w", ErrIndexOutOfRange)
	}
	element := t.elements[index]
	t.elements = append(t.elements[:index], t.elements[index+1:]...)
	return element, nil
}

// Pop removes the last element from the list.
func (t *ArrayList[T]) Pop() (T, error) {
	if len(t.elements) == 0 {
		zero := new(T)
		return *zero, fmt.Errorf("%w", ErrPopFromEmptyList)
	}
	element := t.elements[len(t.elements)-1]
	t.elements = t.elements[:len(t.elements)-1]
	return element, nil
}

// PopFront removes the first element from the list.
func (t *ArrayList[T]) PopFront() (T, error) {
	if len(t.elements) == 0 {
		zero := new(T)
		return *zero, fmt.Errorf("%w", ErrPopFromEmptyList)
	}
	element := t.elements[0]
	t.elements = t.elements[1:]
	return element, nil
}

// Set replaces the element at the specified position in the list with a new element.
func (t *ArrayList[T]) Set(index int, value T) error {
	if index < 0 || index >= len(t.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	t.elements[index] = value

	return nil
}

// Clear removes all elements from the list.
func (t *ArrayList[T]) Clear() {
	t.elements = t.elements[:0]

	//t.elements[:0] clears the slice by truncating it to 0 elements while
	//preserving the underlying array capacity.

	//t.elements[:0:0] clears the slice and releases the underlying array memory.
	//It truncates the slice to 0 elements and 0 capacity.
	//t.elements = t.elements[:0:0]
}

// Contains returns true if the list contains the specified element.
func (t *ArrayList[T]) Contains(values ...T) bool {
	for _, searchValue := range values {
		found := false
		for _, element := range t.elements {
			if t.cmp(element, searchValue) == 0 {
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
func (t *ArrayList[T]) RemoveRange(fromIndex int, toIndex int) error {
	if fromIndex < 0 || fromIndex >= len(t.elements) || toIndex < 0 || toIndex >= len(t.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	if fromIndex > toIndex {
		return fmt.Errorf("%w", ErrFormIndexMustBeLessThanToIndex)
	}

	t.elements = append(t.elements[:fromIndex], t.elements[toIndex:]...)

	return nil
}

// SubList returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive.
func (t *ArrayList[T]) SubList(fromIndex int, toIndex int) ([]T, error) {
	if fromIndex < 0 || fromIndex >= len(t.elements) || toIndex < 0 || toIndex >= len(t.elements) {
		return nil, fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	if fromIndex > toIndex {
		return nil, fmt.Errorf("%w", ErrFormIndexMustBeLessThanToIndex)
	}

	return t.elements[fromIndex:toIndex], nil
}

// Reverse reverses the order of the elements in the list.
func (t *ArrayList[T]) Reverse() {
	length := len(t.elements)
	if length < 2 {
		return
	}
	for left, right := 0, length-1; left < right; left, right = left+1, right-1 {
		t.elements[left], t.elements[right] = t.elements[right], t.elements[left]
	}
	// Alternative
	//for i := length/2 - 1; i >= 0; i-- {
	//	opp := length - 1 - i
	//	t.elements[i], t.elements[opp] = t.elements[opp], t.elements[i]
	//}
}

// RemoveIf removes all elements that match the specified predicate.
func (t *ArrayList[T]) RemoveIf(predicate func(T) bool) bool {
	n := 0
	for _, x := range t.elements {
		if !predicate(x) {
			t.elements[n] = x
			n++
		}
	}
	if n < len(t.elements) {
		t.elements = t.elements[:n]
		return true
	}
	return false
}

// Sort sorts the elements of the list.
func (t *ArrayList[T]) Sort() {
	sort.Slice(t.elements, func(i, j int) bool {
		return t.cmp(t.elements[i], t.elements[j]) < 0
	})
}

// Copy deep copies the list.
func (t *ArrayList[T]) Copy() []T {
	dst := make([]T, len(t.elements))
	copy(dst, t.elements)
	return dst
	// Alternative 1: return append([]T(nil), src...)
	// This creates an empty slice and appends all elements of src, achieving a deep copy.
	//Alternative 2: return append(src[:0:0], src...)
	// This creates an empty slice based on src and appends all elements of src, achieving a deep copy.
}

// BatchSplit splits the list into batches of the specified size.
func (t *ArrayList[T]) BatchSplit(batchSize int) [][]T {
	if batchSize <= 0 {
		return [][]T{}
	}

	values := t.elements

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
func (t *ArrayList[T]) SlidingWindows(size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}
	// returns the input slice as the first element
	if len(t.elements) <= size {
		return [][]T{t.elements}
	}

	// allocate slice at the precise size we need
	// For example, if t.elements = [1, 2, 3, 4, 5, 6, 7] and size = 3,
	// then r = make([][]T, 0, 7-3+1) = make([][]T, 0, 5), with capacity of 5.
	r := make([][]T, 0, len(t.elements)-size+1)

	// Get start and end indices of sliding window and append window slice to result slice.
	for i, j := 0, size; j <= len(t.elements); i, j = i+1, j+1 {
		r = append(r, t.elements[i:j])
	}

	return r
}

// BringElementToFront moves the specified element to the front of the ArrayList.
// If the element is already at the front in the list, no action is taken.
// If the element is not in the list, it is added to the front of the list.
// This operation modifies the ArrayList in place whenever possible.
func (t *ArrayList[T]) BringElementToFront(element T) {
	if len(t.elements) != 0 && t.cmp(t.elements[0], element) == 0 {
		return
	}
	previousElement := element //Initialize previousElement to the target element
	for i, currentElement := range t.elements {
		switch {
		case i == 0:
			t.elements[0] = element
			previousElement = currentElement
		case t.cmp(currentElement, element) == 0:
			t.elements[i] = previousElement
			return
		default: //Otherwise
			t.elements[i] = previousElement
			previousElement = currentElement
		}
	}
	t.elements = append(t.elements, previousElement)
}

// RemoveDuplicates removes duplicates from the ArrayList.
// Before removing duplicates, it sorts the ArrayList.
func (t *ArrayList[T]) RemoveDuplicates() {
	if len(t.elements) <= 1 {
		return
	}

	// Start by sorting the ArrayList
	t.Sort()

	slow := 0
	for fast := 1; fast < len(t.elements); fast++ {
		if t.cmp(t.elements[slow], t.elements[fast]) == 0 {
			continue
		}
		slow++
		t.elements[slow] = t.elements[fast]

	}
	// Shrink the size of the ArrayList after removing duplicates
	t.elements = t.elements[:slow+1]
}

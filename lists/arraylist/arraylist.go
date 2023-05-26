package arraylist

import (
	"errors"
	"fmt"
	"github.com/kwstars/goads/containers"
	"github.com/kwstars/goads/lists"
	"sort"
)

var (
	ErrIndexOutOfRange                = errors.New("index out of range")
	ErrFormIndexMustBeLessThanToIndex = errors.New("fromIndex must be less than or equal to toIndex")
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

func (t *ArrayList[T]) Values() []interface{} {
	//TODO implement me
	panic("implement me")
}

func (t *ArrayList[T]) String() string {
	//TODO implement me
	panic("implement me")
}

func (t *ArrayList[T]) Iter() containers.Iterator[T] {
	//TODO implement me
	panic("implement me")
}

// Add inserts an element at the specified index.
func (t *ArrayList[T]) Add(index int, element T) error {
	if index < 0 || index >= len(t.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	t.elements = append(t.elements[:index], append([]T{element}, t.elements[index:]...)...)

	return nil
}

// Append adds an element to the end of the list.
func (t *ArrayList[T]) Append(element T) {
	t.elements = append(t.elements, element)
}

// AddAll inserts multiple elements at a specific position in the list.
func (t *ArrayList[T]) AddAll(index int, elements []T) error {
	if index < 0 || index >= len(t.elements) {
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
	t.elements = []T{}
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
}

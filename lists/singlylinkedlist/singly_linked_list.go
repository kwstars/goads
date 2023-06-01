package singlylinkedlist

import (
	"errors"
	"fmt"

	"github.com/kwstars/goads/lists"
)

var (
	ErrIndexOutOfRange                = errors.New("index out of range")
	ErrFormIndexMustBeLessThanToIndex = errors.New("fromIndex must be less than or equal to toIndex")
)

var _ lists.List[int] = (*List[int])(nil)

// Element is a single element in the list.
type element[T any] struct {
	value T
	next  *element[T]
}

// List is a singly linked list.
type List[T any] struct {
	head *element[T]       // head is the first element in the list.
	tail *element[T]       // tail is the last element in the list.
	size int               // size is the number of elements in the list.
	cmp  func(a, b T) int8 // cmp should return a negative number if a < b, zero if a == b, and a positive number if a > b.
}

// New creates a new singly linked list.
func New[T any](cmp func(a, b T) int8) *List[T] {
	zero := new(T)
	sentinel := &element[T]{value: *zero, next: nil}
	return &List[T]{head: sentinel, tail: sentinel, size: 0, cmp: cmp}
}

func (l *List[T]) Empty() bool {
	return l.size == 0
}

func (l *List[T]) Size() int {
	return l.size
}

// Append adds an element to the end of the list.
func (l *List[T]) Append(value T) {
	newElem := &element[T]{value: value, next: nil}
	l.tail.next = newElem // Set next pointer of current tail element to point to the new element
	l.tail = newElem      // Update tail pointer to point to the new tail element
	l.size++
}

// Prepend adds an element to the start of the list.
func (l *List[T]) Prepend(value T) {
	newElem := &element[T]{value: value, next: l.head.next} // Set next pointer of new element to point to the current first element
	l.head.next = newElem                                   // Set next pointer of head to point to the new first element
	if l.tail == l.head {                                   // If list is empty, update tail pointer
		l.tail = newElem
	}
	l.size++
}

// Get retrieves an element at a specific position in the list.
func (l *List[T]) Get(index int) (T, error) {
	if index < 0 || index >= l.size {
		zero := new(T)
		return *zero, fmt.Errorf("%w: %d, size: %d", ErrIndexOutOfRange, index, l.size)
	}
	cur := l.head
	for i := 0; i <= index; i++ {
		cur = cur.next
	}
	return cur.value, nil
}

// Insert inserts an element at a specific position in the list.
func (l *List[T]) Insert(index int, value T) error {
	// Check index bounds
	if index < 0 || index > l.size {
		return fmt.Errorf("%w: %d, size: %d", ErrIndexOutOfRange, index, l.size)
	}

	// Find the element before the insertion point
	prev := l.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	// Create a new element and insert it
	newElem := &element[T]{value: value, next: prev.next}
	prev.next = newElem

	// Update tail if new element was inserted at the end of the list
	if prev == l.tail {
		l.tail = newElem
	}

	// Update list size
	l.size++

	return nil
}

// InsertAll inserts all elements at a specific position in the list.
func (l *List[T]) InsertAll(index int, values []T) error {
	// Check index bounds
	if index < 0 || index > l.size {
		return fmt.Errorf("%w: %d, size: %d", ErrIndexOutOfRange, index, l.size)
	}

	prev := l.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	// Insert all elements
	for _, value := range values {
		newElem := &element[T]{value: value, next: prev.next}
		prev.next = newElem
		prev = newElem

		// Update tail if new element was inserted at the end of the list
		if prev == l.tail {
			l.tail = newElem
		}

		// Update list size
		l.size++
	}

	return nil
}

// Clear removes all elements from the list.
func (l *List[T]) Clear() {
	// Reset the head and tail to sentinel
	l.head.next = nil
	l.tail = l.head
	l.size = 0
}

// Remove removes an element at a specific position in the list.
func (l *List[T]) Remove(index int) error {
	// Check index bounds
	if index < 0 || index >= l.size {
		return fmt.Errorf("%w: %d, size: %d", ErrIndexOutOfRange, index, l.size)
	}

	// Find the element before the one to be removed
	prev := l.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	// Remove the element
	prev.next = prev.next.next

	// If we removed the last element, update the tail
	if prev.next == nil {
		l.tail = prev
	}

	// Update list size
	l.size--

	return nil
}

// RemoveRange removes a range of elements from the list.
func (l *List[T]) RemoveRange(fromIndex int, toIndex int) error {
	// Check index bounds
	if fromIndex < 0 || toIndex > l.size {
		return fmt.Errorf("%w, fromIndex: %d, toIndex: %d, size: %d", ErrIndexOutOfRange, fromIndex, toIndex, l.size)
	}

	if fromIndex > toIndex {
		return fmt.Errorf("%w, fromIndex: %d, toIndex: %d", ErrFormIndexMustBeLessThanToIndex, fromIndex, toIndex)
	}

	// If fromIndex and toIndex are the same, there's no range to remove.
	if fromIndex == toIndex {
		return nil
	}

	// Find the element before the start of the range to be removed
	prev := l.head
	for i := 0; i < fromIndex; i++ {
		prev = prev.next
	}

	// Find the element at the end of the range to be removed
	end := prev
	for i := fromIndex; i < toIndex; i++ {
		end = end.next
	}

	// Remove the range
	prev.next = end.next

	// If we removed the last element, update the tail
	if end == l.tail {
		l.tail = prev
	}

	// Update list size
	l.size -= toIndex - fromIndex

	return nil
}

// Set sets the value of an element at a specific position in the list.
func (l *List[T]) Set(index int, value T) error {
	// Check index bounds
	if index < 0 || index >= l.size {
		return fmt.Errorf("%w: %d, size: %d", ErrIndexOutOfRange, index, l.size)
	}

	// Find the element at the specified index
	cur := l.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	// Set the value of the element
	cur.value = value

	return nil
}

// IndexOf returns the index of the first occurrence of the specified value in the list.
func (l *List[T]) IndexOf(value T) int {
	// Traverse the list and find the index of the first occurrence of the value
	index := 0
	for cur := l.head.next; cur != nil; cur = cur.next {
		if l.cmp(cur.value, value) == 0 {
			return index
		}
		index++
	}

	// Return -1 if the value is not found
	return -1
}

// LastIndexOf returns the index of the last occurrence of the specified value in the list.
func (l *List[T]) LastIndexOf(value T) int {
	// Traverse the list and find the index of the last occurrence of the value
	index := -1
	for i, cur := 0, l.head.next; cur != nil; i, cur = i+1, cur.next {
		if l.cmp(cur.value, value) == 0 {
			index = i
		}
	}

	return index
}

// SubList returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive.
func (l *List[T]) SubList(fromIndex int, toIndex int) ([]T, error) {
	// Check index bounds
	if fromIndex < 0 || toIndex > l.size {
		return nil, fmt.Errorf("%w, fromIndex: %d, toIndex: %d, size: %d", ErrIndexOutOfRange, fromIndex, toIndex, l.size)
	}

	if fromIndex > toIndex {
		return nil, fmt.Errorf("%w, fromIndex: %d, toIndex: %d", ErrFormIndexMustBeLessThanToIndex, fromIndex, toIndex)
	}

	if fromIndex == toIndex {
		return []T{}, nil
	}

	// Traverse the list until the start of the sublist
	cur := l.head.next
	for i := 0; i < fromIndex; i++ {
		cur = cur.next
	}

	// Traverse the sublist and collect the values
	var values []T
	for i := fromIndex; i < toIndex; i, cur = i+1, cur.next {
		values = append(values, cur.value)
	}

	return values, nil
}

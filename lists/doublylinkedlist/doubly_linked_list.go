package doublylinkedlist

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
	prev  *element[T]
}

// List is a doubly linked list.
type List[T any] struct {
	head *element[T]       // head is a sentinel, its next pointer points to the first element in the list.
	tail *element[T]       // tail is a sentinel, its prev pointer points to the last element in the list.
	size int               // size is the number of elements in the list.
	cmp  func(a, b T) int8 // cmp should return a negative number if a < b, zero if a == b, and a positive number if a > b.
}

// New creates a new doubly linked list.
func New[T any](cmp func(a, b T) int8) *List[T] {
	sentinelHead := &element[T]{}
	sentinelTail := &element[T]{}
	sentinelHead.next = sentinelTail
	sentinelTail.prev = sentinelHead
	return &List[T]{head: sentinelHead, tail: sentinelTail, size: 0, cmp: cmp}
}

func (l *List[T]) Empty() bool {
	return l.size == 0
}

func (l *List[T]) Full() bool {
	// For a linked list, Full doesn't really make sense because a linked list is
	// not typically constrained by a maximum size. This function could always
	// return false, or it could check for a specific condition if needed.
	return false
}

func (l *List[T]) Size() int {
	return l.size
}

// Append adds an element to the end of the list.
func (l *List[T]) Append(value T) {
	newElem := &element[T]{value: value, next: l.tail, prev: l.tail.prev}
	l.tail.prev.next = newElem
	l.tail.prev = newElem
	l.size++
}

// Prepend adds an element to the start of the list.
func (l *List[T]) Prepend(value T) {
	newElem := &element[T]{value: value, next: l.head.next, prev: l.head}
	l.head.next.prev = newElem
	l.head.next = newElem
	l.size++
}

// Get retrieves an element at a specific position in the list.
func (l *List[T]) Get(index int) (T, error) {
	if index < 0 || index >= l.size {
		zero := new(T)
		return *zero, fmt.Errorf("%w: %d", ErrIndexOutOfRange, index)
	}
	cur := l.head
	for i := 0; i <= index; i++ {
		cur = cur.next
	}
	return cur.value, nil
}

// Insert inserts an element at a specific position in the list.
func (l *List[T]) Insert(index int, value T) error {
	if index < 0 || index > l.size {
		return fmt.Errorf("%w: %d", ErrIndexOutOfRange, index)
	}

	prev := l.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	newElem := &element[T]{value: value, next: prev.next, prev: prev}
	prev.next.prev = newElem
	prev.next = newElem

	l.size++

	return nil
}

// InsertAll inserts all elements at a specific position in the list.
func (l *List[T]) InsertAll(index int, values []T) error {
	if index < 0 || index > l.size {
		return fmt.Errorf("%w: %d", ErrIndexOutOfRange, index)
	}

	prev := l.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	for _, value := range values {
		newElem := &element[T]{value: value, next: prev.next, prev: prev}
		prev.next.prev = newElem
		prev.next = newElem
		prev = newElem

		l.size++
	}

	return nil
}

// Clear removes all elements from the list.
func (l *List[T]) Clear() {
	// Reset the head and tail to sentinel
	l.head.next = l.tail
	l.tail.prev = l.head
	l.size = 0
}

// Remove removes an element at a specific position in the list.
func (l *List[T]) Remove(index int) error {
	if index < 0 || index >= l.size {
		return fmt.Errorf("%w: %d", ErrIndexOutOfRange, index)
	}

	cur := l.head
	for i := 0; i <= index; i++ {
		cur = cur.next
	}

	cur.prev.next = cur.next
	cur.next.prev = cur.prev

	l.size--

	return nil
}

// RemoveRange removes a range of elements from the list.
func (l *List[T]) RemoveRange(fromIndex int, toIndex int) error {
	if fromIndex < 0 || toIndex >= l.size || fromIndex > toIndex {
		return fmt.Errorf("%w, fromIndex: %d, toIndex: %d", ErrIndexOutOfRange, fromIndex, toIndex)
	}

	start := l.head
	for i := 0; i <= fromIndex; i++ {
		start = start.next
	}

	end := start
	for i := fromIndex; i <= toIndex; i++ {
		end = end.next
	}

	start.prev.next = end.next
	end.next.prev = start.prev

	l.size -= toIndex - fromIndex + 1

	return nil
}

// Set sets the value of an element at a specific position in the list.
func (l *List[T]) Set(index int, value T) error {
	if index < 0 || index >= l.size {
		return fmt.Errorf("%w: %d", ErrIndexOutOfRange, index)
	}

	cur := l.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	cur.value = value

	return nil
}

// IndexOf returns the index of the first occurrence of the specified value in the list.
func (l *List[T]) IndexOf(value T) int {
	index := 0
	for cur := l.head.next; cur != nil; cur = cur.next {
		if l.cmp(cur.value, value) == 0 {
			return index
		}
		index++
	}

	return -1
}

// LastIndexOf returns the index of the last occurrence of the specified value in the list.
func (l *List[T]) LastIndexOf(value T) int {
	index := l.size - 1
	for cur := l.tail.prev; cur != l.head; cur = cur.prev {
		if l.cmp(cur.value, value) == 0 {
			return index
		}
		index--
	}

	return -1
}

// SubList returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive.
func (l *List[T]) SubList(fromIndex int, toIndex int) ([]T, error) {
	if fromIndex < 0 || toIndex > l.size || fromIndex > toIndex {
		return nil, fmt.Errorf("%w, fromIndex: %d, toIndex: %d", ErrIndexOutOfRange, fromIndex, toIndex)
	}

	cur := l.head.next
	for i := 0; i < fromIndex; i++ {
		cur = cur.next
	}

	var values []T
	for i := fromIndex; i < toIndex; i, cur = i+1, cur.next {
		values = append(values, cur.value)
	}

	return values, nil
}

package deque

import "github.com/kwstars/goads/containers"

// Deque is an interface for a deque data structure.
type Deque[T any] interface {
	containers.Container[T]
	// PushFront inserts an element at the front of the deque.
	PushFront(element T) bool
	// PushBack inserts an element at the rear of the deque.
	PushBack(element T) bool
	// PopFront removes and returns the element at the front of the deque.
	PopFront() (T, bool)
	// PopBack removes and returns the element at the rear of the deque.
	PopBack() (T, bool)
	// Front returns the element at the front of the deque.
	Front() (T, bool)
	// Back returns the element at the rear of the deque.
	Back() (T, bool)
	// Contains determines whether a given item is already in the deque,
	Contains(elem T) bool
}

package arrayheap

import (
	"errors"
	"fmt"
	"github.com/kwstars/goads/queues/priorityqueue"
)

var (
	ErrHeapEmpty = errors.New("heap is empty")
)

// ArrayHeap implements the priorityqueue.PriorityQueue interface.
var _ priorityqueue.PriorityQueue[int] = (*ArrayHeap[int])(nil)

// ArrayHeap is a generic data structure that supports min-heap and max-heap.
type ArrayHeap[T any] struct {
	data []T               // Stores elements of the heap
	comp func(a, b T) bool // Comparison function to determine min/max heap
}

// NewHeap creates a new ArrayHeap.
func NewHeap[T any](comp func(a, b T) bool) *ArrayHeap[T] {
	return &ArrayHeap[T]{
		data: []T{}, // Initialize empty slice
		comp: comp,  // Set comparison function
	}
}

// Push adds an element to the ArrayHeap.
func (h *ArrayHeap[T]) Push(x T) {
	h.data = append(h.data, x) // Append new element to slice
	h.upHeap(len(h.data) - 1)  // Ensure heap property is maintained
}

// Pop removes the minimum or maximum element from the ArrayHeap.
func (h *ArrayHeap[T]) Pop() (T, error) {
	// Return error if empty
	if len(h.data) == 0 {
		var zero T
		return zero, fmt.Errorf("%w", ErrHeapEmpty)
	}
	x := h.data[0]                    // Get min/max element
	h.data[0] = h.data[len(h.data)-1] // Replace with last element
	h.data = h.data[:len(h.data)-1]   // Remove last element
	h.downHeap(0)                     // Ensure heap property is maintained
	return x, nil                     // Return min/max element
}

// Peek returns the minimum or maximum element from the ArrayHeap without removing it.
func (h *ArrayHeap[T]) Peek() (T, error) {
	if len(h.data) == 0 { // Return error if empty
		var zero T
		return zero, fmt.Errorf("%w", ErrHeapEmpty)
	}
	return h.data[0], nil // Return min/max element
}

// Size returns the number of elements in the ArrayHeap.
func (h *ArrayHeap[T]) Size() int {
	return len(h.data) // Return length of internal slice
}

// IsEmpty returns true if the ArrayHeap is empty.
func (h *ArrayHeap[T]) IsEmpty() bool {
	return len(h.data) == 0 // Check if slice is empty
}

// upHeap ensures that the heap property is maintained for a newly added element.
func (h *ArrayHeap[T]) upHeap(i int) {
	// While not at root and parent is greater/less than current element
	for i > 0 && h.comp(h.data[i], h.data[i/2]) {
		// Swap parent and current element
		h.data[i], h.data[i/2] = h.data[i/2], h.data[i]
		// Move to parent
		i /= 2
	}
}

// downHeap ensures that the heap property is maintained after an element is removed.
func (h *ArrayHeap[T]) downHeap(i int) {
	for {
		// Find children
		left := 2*i + 1
		right := left + 1
		// Break if no children
		if left >= len(h.data) {
			break
		}
		// Find minimum/maximum child
		j := left
		if right < len(h.data) && h.comp(h.data[right], h.data[left]) {
			j = right
		}
		// Swap with child if child is greater/less than parent
		if h.comp(h.data[j], h.data[i]) {
			h.data[i], h.data[j] = h.data[j], h.data[i]
			i = j
		} else {
			break
		}
	}
}

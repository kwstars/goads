// Copyright (c) 2023, Kira。All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package binaryheap implements a binary heap backed by an array list.
// The Comparator function determines whether this heap is a min heap or a max heap.
// A binaryheap is a data structure that is a complete binary tree.
// Note that this structure is not thread-safe.
// References: http://en.wikipedia.org/wiki/Binary_heap
// Visual: https://visualgo.net/en/heap
package binaryheap

import (
	"errors"
	"fmt"

	"github.com/kwstars/goads/pkg/common"
)

var (
	ErrHeapEmpty = errors.New("binary heap is empty")
)

// Option is a function that can be passed to New to customize the BinaryHeap.
type Option[T any] func(*BinaryHeap[T])

// WithInitialCapacity sets the initial capacity of the BinaryHeap.
func WithInitialCapacity[T any](capacity int) Option[T] {
	return func(h *BinaryHeap[T]) {
		h.data = make([]T, 0, capacity) // Initialize slice with capacity
	}
}

// BinaryHeap is a generic data structure that supports min-heap and max-heap.
type BinaryHeap[T any] struct {
	data []T // Stores elements of the heap

	// `comp` is a comparison function for the BinaryHeap that determines the order of elements.
	// It follows different rules in the upHeap and downHeap operations:
	//
	//   - In the upHeap operation:
	//     -- `comp` takes the newly inserted child node as the first argument and its parent node as the second.
	//     -- In a max heap, if the child node > parent node, a swap operation is needed.
	//     -- In a min heap, if the child node < parent node, a swap operation is needed.
	//
	//   - In the downHeap operation:
	//     -- `comp` takes the left or right child node of the heap root as the first argument and the current node as the second.
	//     -- In a max heap, if either left or right child node > current node, a swap operation is needed.
	//     -- In a min heap, if either left or right child node < current node, a swap operation is needed.
	comp common.Comparator[T, T]
}

// New creates a new BinaryHeap.
func New[T any](comp common.Comparator[T, T], options ...Option[T]) *BinaryHeap[T] {
	bh := &BinaryHeap[T]{
		comp: comp,
	}

	for _, option := range options {
		option(bh)
	}

	if bh.data == nil {
		bh.data = make([]T, 0)
	}

	return bh
}

/*
		 0
		/ \
	  1    2
	 / \  / \
	3  4 5   6
*/

// parent returns the parent node index for the node at index i.
func (h *BinaryHeap[T]) parent(i int) int {
	// (i-1)/2
	return (i - 1) >> 1
}

// leftChild Returns the left child node index for the node at index i.
func (h *BinaryHeap[T]) leftChild(i int) int {
	// 2*i + 1
	return i<<1 + 1
}

// rightChild Returns the right child node index for the node at index i.
func (h *BinaryHeap[T]) rightChild(i int) int {
	// 2*i + 2
	return i<<1 + 2
}

// hasLeftChild Returns true if the node at index i has a left child node.
func (h *BinaryHeap[T]) hasLeftChild(i int) bool {
	return h.leftChild(i) < len(h.data)
}

// hasRightChild Returns true if the node at index i has a right child node.
func (h *BinaryHeap[T]) hasRightChild(i int) bool {
	return h.rightChild(i) < len(h.data)
}

// Push adds an element to the BinaryHeap.
func (h *BinaryHeap[T]) Push(x T) {
	h.data = append(h.data, x) // Append new element to slice (at end)
	h.upHeap(len(h.data) - 1)  // Ensure heap property is maintained
}

// Pop removes the minimum or maximum element from the BinaryHeap.
func (h *BinaryHeap[T]) Pop() (T, error) {
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

// Peek returns the minimum or maximum element from the BinaryHeap without removing it.
func (h *BinaryHeap[T]) Peek() (T, error) {
	if len(h.data) == 0 { // Return error if empty
		var zero T
		return zero, fmt.Errorf("%w", ErrHeapEmpty)
	}
	return h.data[0], nil // Return min/max element
}

// Size returns the number of elements in the BinaryHeap.
func (h *BinaryHeap[T]) Size() int {
	return len(h.data) // Return length of internal slice
}

// IsEmpty returns true if the BinaryHeap is empty.
func (h *BinaryHeap[T]) IsEmpty() bool {
	return len(h.data) == 0 // Check if slice is empty
}

// upHeap ensures that the heap property is maintained for a newly added element.
func (h *BinaryHeap[T]) upHeap(i int) {
	// While not at root and parent is greater/less than current element
	for i > 0 && (h.comp(h.data[i], h.data[h.parent(i)]) > 0) {
		// Swap current and parent element
		h.data[i], h.data[h.parent(i)] = h.data[h.parent(i)], h.data[i]
		// Move to parent
		i = h.parent(i)
	}
}

// downHeap ensures that the heap property is maintained after an element is removed.
func (h *BinaryHeap[T]) downHeap(i int) {
	for {
		left := 2*i + 1   // Compute the index of the left child
		right := left + 1 // Compute the index of the right child

		smallest := i // Assume the smallest value is at the parent node

		// Check if the left child exists and if it is smaller than the parent
		if left < len(h.data) && (h.comp(h.data[left], h.data[smallest]) > 0) {
			smallest = left // If so, update smallest to point to the left child
		}

		// Do the same for the right child
		if right < len(h.data) && (h.comp(h.data[right], h.data[smallest])) > 0 {
			smallest = right // If so, update smallest to point to the right child
		}

		// If the smallest value is not at the parent node (i.e., either the left child or the right child is smaller), swap the smallest value with the parent
		if smallest != i {
			h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
			i = smallest // Update i to point to the index of the child node that was swapped
		} else {
			break // If the smallest value is at the parent node, then the heap property is restored, so exit the loop
		}
	}
}

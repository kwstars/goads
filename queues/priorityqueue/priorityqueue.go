// Copyright (c) 2023, Kira. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package priorityqueue implements a priority queue backed by binary heap.
//
// An unbounded priority queue based on a priority queue.
// The elements of the priority queue are ordered by a comparator provided at queue construction time.
//
// The heap of this queue is the least/smallest element with respect to the specified ordering.
// If multiple elements are tied for least value, the heap is one of those elements arbitrarily.
//
// Structure is not thread safe.
//
// References: https://en.wikipedia.org/wiki/Priority_queue
package priorityqueue

import (
	"github.com/kwstars/goads/pkg/common"
	"github.com/kwstars/goads/queues"
	"github.com/kwstars/goads/trees/binaryheap"
)

var _ queues.Queue[int] = (*Queue[int])(nil)

// Queue is a priority queue backed by binary heap.
type Queue[T any] struct {
	heap *binaryheap.BinaryHeap[T] // Binary heap
	comp common.Comparator[T, T]   // Comparator function
}

// New creates a new priority queue.
func New[T any](comp common.Comparator[T, T], options ...binaryheap.Option[T]) *Queue[T] {
	return &Queue[T]{
		heap: binaryheap.New[T](comp, options...),
		comp: comp,
	}
}

// Empty returns true if queue does not contain any elements.
func (pq *Queue[T]) Empty() bool {
	return pq.heap.Empty()
}

// Size returns number of elements within the queue.
func (pq *Queue[T]) Size() int {
	return pq.heap.Size()
}

// Clear clears all values in the queue.
func (pq *Queue[T]) Clear() {
	pq.heap.Clear()
}

// Enqueue adds an element to the priority queue.
func (pq *Queue[T]) Enqueue(item T) {
	pq.heap.Push(item)
}

// Dequeue removes and returns the front element of the priority queue.
func (pq *Queue[T]) Dequeue() (T, error) {
	return pq.heap.Pop()
}

// Peek returns the front element of the priority queue without removing it.
func (pq *Queue[T]) Peek() (T, error) {
	return pq.heap.Peek()
}

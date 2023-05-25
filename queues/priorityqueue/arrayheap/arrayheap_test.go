package arrayheap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHeap(t *testing.T) {
	// Create a new min heap
	heap := NewHeap[int](func(a, b int) bool { return a < b })
	assert.NotNil(t, heap)
	assert.Equal(t, 0, heap.Size())
}
func TestPush(t *testing.T) {
	heap := NewHeap[int](func(a, b int) bool { return a < b })
	heap.Push(2)
	assert.Equal(t, 1, heap.Size()) // Check size after push
	heap.Push(1)
	assert.Equal(t, 2, heap.Size())
	heap.Push(3)
	assert.Equal(t, 3, heap.Size())
}

func TestPop(t *testing.T) {
	heap := NewHeap[int](func(a, b int) bool { return a < b })
	heap.Push(2)
	heap.Push(1)
	heap.Push(3)
	val, err := heap.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 1, val)
	val, err = heap.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)
	val, err = heap.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 3, val)
	// Pop empty heap
	_, err = heap.Pop()
	assert.ErrorIs(t, err, ErrHeapEmpty)
}

func TestPeek(t *testing.T) {
	heap := NewHeap[int](func(a, b int) bool { return a < b })
	heap.Push(2)
	heap.Push(1)
	heap.Push(3)
	val, err := heap.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 1, val)
	// Peek does not remove element
	assert.Equal(t, 3, heap.Size())
}

func TestIsEmpty(t *testing.T) {
	heap := NewHeap[int](func(a, b int) bool { return a < b })
	assert.True(t, heap.IsEmpty()) // Check empty heap

	heap.Push(2)
	assert.False(t, heap.IsEmpty()) // Check non-empty heap

	_, _ = heap.Pop()
	assert.True(t, heap.IsEmpty())
}

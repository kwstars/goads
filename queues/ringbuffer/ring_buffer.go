package main

import (
	"fmt"
	"sync"
)

type Element[T any] struct {
	Value T
}

type RingBuffer[T any] struct {
	mu       sync.Mutex
	elements []Element[T]
	readPos  int
	writePos int
	full     bool
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		elements: make([]Element[T], size),
	}
}

func (rb *RingBuffer[T]) Push(val T) bool {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if rb.full {
		return false
	}

	rb.elements[rb.writePos] = Element[T]{Value: val}
	rb.writePos = (rb.writePos + 1) % len(rb.elements)

	if rb.writePos == rb.readPos {
		rb.full = true
	}

	return true
}

func (rb *RingBuffer[T]) Pop() (T, bool) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if !rb.full && rb.readPos == rb.writePos {
		zeroValue := new(T)
		return *zeroValue, false
	}

	val := rb.elements[rb.readPos].Value
	rb.readPos = (rb.readPos + 1) % len(rb.elements)
	rb.full = false

	return val, true
}

func main() {
	r := NewRingBuffer[int](3)

	for i := 0; i < 5; i++ {
		ok := r.Push(i)
		fmt.Printf("Push(%d): %v\n", i, ok)
	}

	for {
		val, ok := r.Pop()
		if ok {
			fmt.Printf("Pop: %d\n", val)
		} else {
			break
		}
	}
}

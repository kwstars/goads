package containers

import "fmt"

// Iterator is an interface for an iterator over the elements in a container.
type Iterator[T any] interface {
	// HasNext determines if there are more elements in the container to iterate over.
	HasNext() bool
	// Next returns the next element in the container.
	Next() T
}

type Slice[T any] []T

func (s Slice[T]) Iter() Iterator[T] {
	return &sliceIterator[T]{s, 0}
}

type sliceIterator[T any] struct {
	slice []T
	index int
}

func (it *sliceIterator[T]) HasNext() bool {
	return it.index < len(it.slice)
}

func (it *sliceIterator[T]) Next() T {
	v := it.slice[it.index]
	it.index++
	return v
}

func printAll[T any](c Container[T]) {
	for it := c.Iter(); it.HasNext(); {
		fmt.Println(it.Next())
	}
}

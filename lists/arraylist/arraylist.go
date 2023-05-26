package ArrayList

import (
	"errors"
	"fmt"
)

var (
	ErrIndexOutOfRange                = errors.New("index out of range")
	ErrFormIndexMustBeLessThanToIndex = errors.New("fromIndex must be less than or equal to toIndex")
)

type ArrayList[T any] struct {
	elements []T
}

func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{
		elements: make([]T, 0),
	}
}

func (list *ArrayList[T]) Add(index int, element T) error {
	if index < 0 || index >= len(list.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	list.elements = append(list.elements[:index], list.elements[index:]...)
	list.elements[index] = element

	return nil
}

func (list *ArrayList[T]) Append(element T) {
	list.elements = append(list.elements, element)
}

func (list *ArrayList[T]) AddAll(index int, elements []T) error {
	if index < 0 || index >= len(list.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	list.elements = append(list.elements[:index], elements...)
	list.elements = append(list.elements, list.elements[index:]...)

	return nil
}

func (list *ArrayList[T]) Clear() {
	list.elements = []T{}
}

func (list *ArrayList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(list.elements) {
		var zero = new(T)
		return *zero, false
	}

	return list.elements[index], true
}

func (list *ArrayList[T]) Remove(index int) {
	if index < 0 || index >= len(list.elements) {
		return
	}

	list.elements = append(list.elements[:index], list.elements[index+1:]...)
}

func (list *ArrayList[T]) Set(index int, value T) {
	if index < 0 || index >= len(list.elements) {
		return
	}

	list.elements[index] = value
}

func (list *ArrayList[T]) Contains(values ...T) bool {
	for _, value := range values {
		if !list.Contains(value) {
			return false
		}
	}

	return true
}

func (list *ArrayList[T]) RemoveRange(fromIndex int, toIndex int) error {
	if fromIndex < 0 || fromIndex >= len(list.elements) || toIndex < 0 || toIndex >= len(list.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	if fromIndex > toIndex {
		return fmt.Errorf("%w", ErrFormIndexMustBeLessThanToIndex)
	}

	list.elements = append(list.elements[:fromIndex], list.elements[toIndex:]...)

	return nil
}

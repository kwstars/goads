package ArrayList

import (
	"errors"
	"fmt"
	"github.com/kwstars/goads/containers"
	"github.com/kwstars/goads/lists"
)

var (
	ErrIndexOutOfRange                = errors.New("index out of range")
	ErrFormIndexMustBeLessThanToIndex = errors.New("fromIndex must be less than or equal to toIndex")
)

var _ lists.List[int] = (*ArrayList[int])(nil)

type ArrayList[T any] struct {
	elements []T
	cmp      func(a, b T) int8
}

func (t *ArrayList[T]) Empty() bool {
	return len(t.elements) == 0
}

func (t *ArrayList[T]) Full() bool {
	return false
}

func (t *ArrayList[T]) Size() int {
	return len(t.elements)
}

func (t *ArrayList[T]) Values() []interface{} {
	//TODO implement me
	panic("implement me")
}

func (t *ArrayList[T]) String() string {
	//TODO implement me
	panic("implement me")
}

func (t *ArrayList[T]) Iter() containers.Iterator[T] {
	//TODO implement me
	panic("implement me")
}

func (t *ArrayList[T]) Add(index int, element T) error {
	if index < 0 || index >= len(t.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	t.elements = append(t.elements[:index], t.elements[index:]...)
	t.elements[index] = element

	return nil
}

func (t *ArrayList[T]) Append(element T) {
	t.elements = append(t.elements, element)
}

func (t *ArrayList[T]) AddAll(index int, elements []T) error {
	if index < 0 || index >= len(t.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	t.elements = append(t.elements[:index], elements...)
	t.elements = append(t.elements, t.elements[index:]...)

	return nil
}

func (t *ArrayList[T]) Equals(other T) bool {
	//TODO implement me
	panic("implement me")
}

func (t *ArrayList[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(t.elements) {
		var zero = new(T)
		return *zero, fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	return t.elements[index], nil
}

func (t *ArrayList[T]) IndexOf(element T) int {
	for k, v := range t.elements {
		if t.cmp(v, element) == 0 {
			return k
		}
	}
	return -1
}

func (t *ArrayList[T]) LastIndexOf(element T) int {
	for i := len(t.elements) - 1; i >= 0; i-- {
		if t.cmp(t.elements[i], element) == 0 {
			return i
		}
	}
	return -1
}

func (t *ArrayList[T]) Remove(index int) error {
	if index < 0 || index >= len(t.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	t.elements = append(t.elements[:index], t.elements[index+1:]...)

	return nil
}

func (t *ArrayList[T]) Set(index int, value T) error {
	if index < 0 || index >= len(t.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	t.elements[index] = value

	return nil
}

func (t *ArrayList[T]) SubList(fromIndex int, toIndex int) []T {
	//TODO implement me
	panic("implement me")
}

func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{
		elements: make([]T, 0),
	}
}

func (t *ArrayList[T]) Clear() {
	t.elements = []T{}
}

func (t *ArrayList[T]) Contains(values ...T) bool {
	for _, value := range values {
		if !t.Contains(value) {
			return false
		}
	}

	return true
}

func (t *ArrayList[T]) RemoveRange(fromIndex int, toIndex int) error {
	if fromIndex < 0 || fromIndex >= len(t.elements) || toIndex < 0 || toIndex >= len(t.elements) {
		return fmt.Errorf("%w", ErrIndexOutOfRange)
	}

	if fromIndex > toIndex {
		return fmt.Errorf("%w", ErrFormIndexMustBeLessThanToIndex)
	}

	t.elements = append(t.elements[:fromIndex], t.elements[toIndex:]...)

	return nil
}

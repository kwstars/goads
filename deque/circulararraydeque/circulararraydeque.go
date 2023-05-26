package ringbuffer

// CircularArrayDeque is a circular array implementation of a deque.
type CircularArrayDeque[T any] struct {
	buf        []T // buf is the underlying slice
	head, tail int // head and tail are the indices of the first and last elements, respectively
	size       int // size is the number of elements in the deque
	capacity   int // capacity is the size of the underlying slice
}

func NewCircularArrayDeque[T any](capacity int) *CircularArrayDeque[T] {
	return &CircularArrayDeque[T]{
		buf:      make([]T, capacity),
		head:     0,
		tail:     0,
		size:     0,
		capacity: capacity,
	}
}

func (d *CircularArrayDeque[T]) PushFront(element T) bool {
	if d.size == d.capacity {
		return false // or resize and continue
	}

	d.head = (d.head - 1 + d.capacity) % d.capacity // update d.head, decrementing and wrapping around
	d.buf[d.head] = element
	d.size++

	return true
}

func (d *CircularArrayDeque[T]) PushBack(element T) bool {
	if d.size == d.capacity {
		return false // or resize and continue
	}

	d.buf[d.tail] = element            // insert at d.tail
	d.tail = (d.tail + 1) % d.capacity // increment d.tail, wrapping around
	d.size++

	return true
}

func (d *CircularArrayDeque[T]) PopFront() (T, bool) {
	if d.size <= 0 {
		var zero T
		return zero, false
	}

	value := d.buf[d.head]             // get value at d.head
	d.head = (d.head + 1) % d.capacity // increment d.head, wrapping around
	d.size--

	return value, true
}

func (d *CircularArrayDeque[T]) PopBack() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}

	d.tail = (d.tail - 1 + d.capacity) % d.capacity // decrement d.tail, wrapping around
	value := d.buf[d.tail]
	d.size--

	return value, true
}

func (d *CircularArrayDeque[T]) Front() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}

	return d.buf[d.head], true
}

func (d *CircularArrayDeque[T]) Back() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}

	tail := (d.tail - 1 + d.capacity) % d.capacity

	return d.buf[tail], true
}

//func (d *CircularArrayDeque[T]) Contains(elem T) bool {
//	for i := 0; i < d.size; i++ {
//		index := (d.head + i) % d.capacity
//		if d.buf[index] == elem {
//			return true
//		}
//	}
//
//	return false
//}

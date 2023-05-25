package priorityqueue

// PriorityQueue is an interface for a priority queue data structure.
// Elements in the priority queue are ordered by priority, allowing quick
// access to the highest (or lowest) priority element.
type PriorityQueue[T any] interface {
	// Push an element into the priority queue.
	Push(item T)
	// Pop the element with the highest (or lowest) priority from the priority queue.
	Pop() (T, error)
	// Peek Views the element with the highest (or lowest) priority in the priority queue but
	// does not delete the element.
	Peek() (T, error)
	// Size Returns the number of elements in the priority queue.
	Size() int
	// IsEmpty Determines if the priority queue is empty.
	IsEmpty() bool
}

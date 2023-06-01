package queues

import "github.com/kwstars/goads/containers"

type Queue[T any] interface {
	containers.Container[T]

	// Enqueue adds an element to the back of the queue.
	Enqueue(item T)
	// Dequeue removes and returns the front element of the queue.
	Dequeue() (T, error)
	// Peek returns the front element of the queue without removing it.
	Peek() (T, error)
}

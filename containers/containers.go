package containers

// Container is an interface for a container data structure.
type Container[T any] interface {
	// Empty determines if a container is empty.
	Empty() bool
	// Size returns the number of elements in a container.
	Size() int
	// Clear removes all elements from a container.
	Clear()
	// String returns a string representation of container.
	//String() string
}

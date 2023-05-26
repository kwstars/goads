package containers

// Container is an interface for a container data structure.
type Container[T any] interface {
	// Empty determines if a container is empty.
	Empty() bool
	// Full determines if a container is full.
	Full() bool
	// Size returns the number of elements in a container.
	Size() int
	// Clear removes all elements from a container.
	Clear()
	// Values returns all elements in a container.
	Values() []interface{}
	// String returns a string representation of container.
	String() string
	// Iter returns an iterator over the elements in a container.
	Iter() Iterator[T]
}

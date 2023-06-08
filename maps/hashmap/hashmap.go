package hashmap

import (
	"github.com/kwstars/goads/maps"
)

// Verify that Map implements the maps.Map interface.
var _ maps.Map[int, int] = (*Map[int, int])(nil)

// Option is a function than can be passed to New to customize the Map.
type Option[K comparable, V any] func(*Map[K, V])

// WithInitialCapacity sets the initial capacity of the BinaryHeap.
func WithInitialCapacity[K comparable, V any](capacity int) Option[K, V] {
	return func(m *Map[K, V]) {
		m.m = make(map[K]V, capacity)
	}
}

// Map is a hash map implementation of the Map interface.
type Map[K comparable, V any] struct {
	m map[K]V
}

// New returns a new hash map.
func New[K comparable, V any](opts ...Option[K, V]) *Map[K, V] {
	m := &Map[K, V]{}

	for _, option := range opts {
		option(m)
	}

	if m.m == nil {
		m.m = make(map[K]V)
	}

	return m
}

// Put inserts a key-value pair into the hash map.
func (m *Map[K, V]) Put(key K, value V) {
	m.m[key] = value
}

// Get returns the value associated with the given key.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	v, b := m.m[key]
	return v, b
}

// Remove removes the key-value pair associated with the given key.
func (m *Map[K, V]) Remove(key K) {
	delete(m.m, key)
}

// Empty returns true if the hash map is empty, false otherwise.
func (m *Map[K, V]) Empty() bool {
	return len(m.m) == 0
}

// Size returns the number of elements in the hash map.
func (m *Map[K, V]) Size() int {
	return len(m.m)
}

// Clear removes all elements from the hash map.
func (m *Map[K, V]) Clear() {
	for k := range m.m {
		delete(m.m, k)
	}
}

package maps

import "github.com/kwstars/goads/containers"

type Map[K comparable, V any] interface {
	containers.Container[K]
	// Put inserts a key-value pair into the map.
	Put(key K, value V)
	// Get returns the value associated with the given key.
	Get(key K) (value V, found bool)
	// Remove removes the key-value pair associated with the given key.
	Remove(key K)
}

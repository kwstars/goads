package lru

import (
	"testing"
)

func TestLRU(t *testing.T) {
	// Create a new LRU cache with capacity 2
	lru := New(2)

	// Add two items to the cache
	lru.Put("key1", "value1")
	lru.Put("key2", "value2")

	// Retrieve the first item from the cache
	value, ok := lru.Get("key1")
	if !ok || value != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}

	// Add a third item to the cache, which should evict the oldest item
	lru.Put("key3", "value3")

	// Try to retrieve the evicted item from the cache
	_, ok = lru.Get("key2")
	if ok {
		t.Errorf("Expected key2 to be evicted from the cache")
	}
}

package lru

import "container/list"

// item is an entry in the cache
type item struct {
	key   string
	value interface{}
}

// LRU is an LRU cache implementation
type LRU struct {
	capacity int                      // maximum number of items in the cache
	linklist *list.List               // doubly linked list
	cache    map[string]*list.Element // hashmap for quick lookups
}

// New creates a new LRU cache with the given capacity.
// If capacity is zero or less, the cache will have a capacity of 1.
func New(capacity int) *LRU {
	if capacity <= 0 {
		capacity = 1
	}
	return &LRU{
		capacity: capacity,
		linklist: list.New(),
		cache:    make(map[string]*list.Element),
	}
}

// Get retrieves an item from the cache. Returns the value for the key and
// true if the key was found.
func (l *LRU) Get(key string) (interface{}, bool) {
	if element, ok := l.cache[key]; ok {
		l.linklist.MoveToFront(element)
		return element.Value.(*item).value, true
	}
	return nil, false
}

// Put adds an item to the cache. If an item with the given key already
// exists, its value is updated.
func (l *LRU) Put(key string, value interface{}) {
	if element, ok := l.cache[key]; ok {
		l.linklist.MoveToFront(element)
		element.Value.(*item).value = value
		return
	}

	element := l.linklist.PushFront(&item{key, value})
	l.cache[key] = element

	if l.capacity != 0 && l.linklist.Len() > l.capacity {
		l.purgeOldest()
	}
}

// purgeOldest removes the oldest item from the cache.
func (l *LRU) purgeOldest() {
	if element := l.linklist.Back(); element != nil {
		l.linklist.Remove(element)
		key := element.Value.(*item).key
		delete(l.cache, key)
	}
}

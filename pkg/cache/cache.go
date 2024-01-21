package cache

import "sync"

type Cache[Key comparable, Val any] struct {
	mu    sync.RWMutex
	items map[Key]Val
	locks map[Key]*sync.Mutex
}

func NewCache[Key comparable, Val any]() *Cache[Key, Val] {
	return &Cache[Key, Val]{
		items: make(map[Key]Val),
		locks: make(map[Key]*sync.Mutex),
	}
}

func (c *Cache[Key, Val]) Set(key Key, value Val) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = value
}

func (c *Cache[Key, Val]) Get(key Key) (Val, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, isFound := c.items[key]
	return value, isFound
}

func (c *Cache[Key, Val]) LockForKey(key Key) *sync.Mutex {
	c.mu.Lock()
	defer c.mu.Unlock()

	lock, exists := c.locks[key]

	if !exists {
		lock = &sync.Mutex{}
		c.locks[key] = lock
	}

	return lock
}

func (c *Cache[Key, Val]) UnlockForKey(key Key) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if lock, exists := c.locks[key]; exists {
		lock.Unlock()
	}
}

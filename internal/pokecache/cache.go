package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mut   sync.Mutex
}

// create a new Cache, start the reapLoop and return the *Cache
func NewCache(interval time.Duration) *Cache {
	c := Cache{cache: make(map[string]cacheEntry)}
	c.reapLoop(interval)
	return &c
}

// add a cachentry
func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock()
	c.cache[key] = cacheEntry{time.Now(), val}
	c.mut.Unlock()
}

// get a cacheEntry, if it exists
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	entry, exists := c.cache[key]
	c.mut.Unlock()
	if exists {
		return entry.val, true
	} else {
		return nil, false
	}

}

// reap entries older than interval
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.mut.Lock()
			for key, value := range c.cache {
				if time.Now().After(value.createdAt.Add(interval)) {
					delete(c.cache, key)
				}
			}
			c.mut.Unlock()
		}
	}()
}

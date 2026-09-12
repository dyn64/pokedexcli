package pokecache

import (
	"sync"
	"time"
)

// cacheEntry just contains the time of creatoin and data as []byte
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// the Cache struct is just a map of entries and a mutex to be able to lock/unlock the cache
type Cache struct {
	cache map[string]cacheEntry
	mut   *sync.Mutex
}

// create a new Cache, start the reapLoop and return the *Cache
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mut:   &sync.Mutex{},
	}
	// starts the reaploop with the set interval as a (go) subroutine
	go c.reapLoop(interval)
	return c
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
	defer c.mut.Unlock()
	entry, exists := c.cache[key]
	return entry.val, exists

}

// reap entries older than interval
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

// does the actual reaping
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mut.Lock()
	defer c.mut.Unlock()
	for key, value := range c.cache {
		if now.After(value.createdAt.Add(last)) {
			delete(c.cache, key)
		}
	}

}

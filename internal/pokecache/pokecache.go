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
	mux   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: map[string]cacheEntry{},
		mux:   &sync.Mutex{},
	}

	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mux.Lock()
		now := time.Now()
		for key, entry := range c.cache {
			if now.Sub(entry.createdAt) > interval {
				delete(c.cache, key)
			}
		}
		c.mux.Unlock()
	}
}

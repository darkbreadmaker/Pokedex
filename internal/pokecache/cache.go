package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	cache map[string]cacheEntry
	mux *sync.Mutex
	}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cache: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}
	go newCache.reaploop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()	
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock() 
	defer c.mux.Unlock()
	value, ok := c.cache[key]
	return value.val, ok
}

func (c *Cache) reaploop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mux.Lock()
		for key, value := range c.cache {
			if time.Since(value.createdAt) > interval {
				delete(c.cache, key)
			}
		}
		c.mux.Unlock()
	}
}

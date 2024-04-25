package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       *sync.Mutex
	CacheMap map[string]CacheEntry
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(interval time.Duration) *Cache {

	c := &Cache{
		CacheMap: make(map[string]CacheEntry),
		mu:       &sync.Mutex{},
	}

	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.CacheMap[key] = CacheEntry{
		CreatedAt: time.Now().UTC(),
		Val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.CacheMap[key]
	return val.Val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.CacheMap {
		if v.CreatedAt.Before(now.Add(-interval)) {
			delete(c.CacheMap, k)
		}
	}
}

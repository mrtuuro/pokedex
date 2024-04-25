package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       *sync.Mutex
	CacheMap map[string]CacheEntry
	Pokedex  map[string]Pokemon
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Pokemon struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	BaseExp int    `json:"base_experience"`
	Height  int    `json:"height"`
	Weight  int    `json:"weight"`
	Stats   []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			StatName string `json:"name"`
		} `json:"stat"`
    } `json:"stats"`
	Types []struct {
		Type struct {
            TypeName string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func NewCache(interval time.Duration) *Cache {

	c := &Cache{
		CacheMap: make(map[string]CacheEntry),
		Pokedex:  make(map[string]Pokemon),
		mu:       &sync.Mutex{},
	}

	return c
}

func (c *Cache) AddPokemon(key string, value Pokemon) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Pokedex[key] = value
}

func (c *Cache) GetPokemon(key string) (Pokemon, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.Pokedex[key]
	return val, ok
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

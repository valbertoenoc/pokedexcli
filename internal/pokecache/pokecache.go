package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	ticker   time.Ticker
	interval time.Duration
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		entries:  map[string]cacheEntry{},
		ticker:   *time.NewTicker(interval),
		interval: interval,
	}

	go c.reapLoop()

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, ok
	}

	return entry.val, ok
}

func (c *Cache) reapLoop() {
	for range c.ticker.C {
		c.sweepCache()
	}
}

func (c *Cache) sweepCache() {
	for key, entry := range c.entries {
		if time.Since(entry.createdAt) > c.interval {
			c.mu.Lock()
			delete(c.entries, key)
			c.mu.Unlock()
		}
	}
}

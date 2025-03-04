package pokecache

import (
	"errors"
	"sync"
	"time"
)

type cacheEntry struct {
	Val       []byte
	createdAt time.Time
}

type Cache struct {
	Entries map[string]cacheEntry
	mu      *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Entries: map[string]cacheEntry{},
		mu:      &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	entries := c.Entries

	if _, ok := entries[key]; !ok {
		entries[key] = cacheEntry{
			Val:       val,
			createdAt: time.Now(),
		}

		return nil
	}

	return errors.New("key already exists")
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entries := c.Entries

	if e, ok := entries[key]; ok {
		return e.Val, true
	}

	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		entries := c.Entries
		for key, e := range entries {
			if time.Since(e.createdAt) > interval {
				delete(entries, key)
			}
		}
		c.mu.Unlock()

	}
}

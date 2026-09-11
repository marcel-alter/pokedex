package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry
	mu      *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		Entries: make(map[string]cacheEntry),
		mu:      &sync.RWMutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c Cache) Add(key string, value []byte) {
	c.mu.Lock()
	c.Entries[key] = cacheEntry{createdAt: time.Now(), val: value}
	c.mu.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exist := c.Entries[key]
	if exist {
		return c.Entries[key].val, true
	} else {
		return []byte{}, false
	}
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		c.mu.RLock()
		var staleKeys []string
		for key, value := range c.Entries {
			if time.Since(value.createdAt) > interval {
				staleKeys = append(staleKeys, key)
			}
		}
		c.mu.RUnlock()
		c.mu.Lock()
		for _, key := range staleKeys {
			delete(c.Entries, key)
		}
		c.mu.Unlock()
	}
}

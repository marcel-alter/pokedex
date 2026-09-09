package pokecache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry
	mu      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	values    []byte
}

func NewCache() Cache {
	newCache := Cache{
		Entries: make(map[string]cacheEntry),
	}
	return newCache
}

func (C Cache) Add(key string, val []byte) error {
	_, exist := C.Entries[key]
	if exist {
		return errors.New("Value already exists!")
	}
	C.Entries[key] = cacheEntry{createdAt: time.Now(), values: val}
	return nil
}

func (C Cache) Get(key string) ([]byte, bool) {
	_, exist := C.Entries[key]
	if exist {
		return C.Entries[key].values, true
	} else {
		return []byte{}, false
	}
}

func (C Cache) reapLoop(key string) {

}

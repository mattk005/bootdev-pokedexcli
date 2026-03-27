package pokecache

import (
	"fmt"
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	fmt.Printf("NewCache(%d) has been run\n", interval)
	data := make(map[string]cacheEntry)
	cashe := Cache{data: data}
	go cashe.reapLoop(interval)
	return &cashe
}

func (cache *Cache) Add(key string, value []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.data[key] = cacheEntry{createdAt: time.Now(), val: value}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if _, ok := cache.data[key]; !ok {
		var foo []byte
		return foo, ok
	} else {
		return cache.data[key].val, ok
	}
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(time.Duration(interval))

	for range ticker.C {
		cache.mu.Lock()
		for key := range cache.data {
			if cache.data[key].createdAt.Before(time.Now().Add(-interval)) {
				delete(cache.data, key)
			}
		}
		cache.mu.Unlock()
	}
}

type Cache struct {
	data map[string]cacheEntry
	mu   sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

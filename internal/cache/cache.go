package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value []byte
}

type CacheType struct {
	entries map[string]cacheEntry
	mux *sync.Mutex
}

func NewCache(waitToClear time.Duration) CacheType {
	cache := CacheType{
		entries: map[string]cacheEntry{},
		mux: &sync.Mutex{},
	}
	go cache.cleanCacheLoop(waitToClear)
	return cache
}

func (cache *CacheType) cleanCacheLoop (waitToClear time.Duration) {
	ticker := time.NewTicker(waitToClear)
	for range ticker.C {
		cache.cleanCache(waitToClear)
	}
}

func (cache *CacheType) cleanCache(waitToClear time.Duration) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	for k, v := range cache.entries {
		if v.createdAt.Before(time.Now().Add(-waitToClear)) {
			delete(cache.entries, k)
		}
	}
}

func (cache *CacheType) Add(key string, value []byte) {
	actualCache := *cache
	actualCache.mux.Lock()
	defer actualCache.mux.Unlock()
	actualCache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		value: value,
	}
}

func (cache *CacheType) Get(key string) ([]byte, bool) {
	actualCache := *cache
	entry, ok := actualCache.entries[key]
	if !ok {
		return []byte{}, false
	}

	return entry.value, true
}

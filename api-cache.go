package main

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value []byte
}

type cacheType struct {
	entries map[string]cacheEntry
	mux *sync.Mutex
}

func newCache(waitToClear time.Duration) cacheType {
	cache := cacheType{
		entries: map[string]cacheEntry{},
		mux: &sync.Mutex{},
	}
	go cache.cleanCacheLoop(waitToClear)
	return cache
}

func (cache *cacheType) cleanCacheLoop (waitToClear time.Duration) {
	ticker := time.NewTicker(waitToClear)
	for range ticker.C {
		cache.cleanCache(waitToClear)
	}
}

func (cache *cacheType) cleanCache(waitToClear time.Duration) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	for k, v := range cache.entries {
		if v.createdAt.Before(time.Now().Add(-waitToClear)) {
			delete(cache.entries, k)
		}
	}
}

func (cache *cacheType) add(key string, value []byte) {
	actualCache := *cache
	actualCache.mux.Lock()
	defer actualCache.mux.Unlock()
	actualCache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		value: value,
	}
}

func (cache *cacheType) get(key string) ([]byte, bool) {
	actualCache := *cache
	entry, ok := actualCache.entries[key]
	if !ok {
		return []byte{}, false
	}

	return entry.value, true
}

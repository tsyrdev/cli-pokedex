package pokecache

import (
    "time"
    "sync"
)

type cacheEntry struct {
    createdAt   time.Time
    val         []byte
}

type Cache struct {
    cache   map[string]cacheEntry
    mux     *sync.Mutex 
}

func (c *Cache) Add(key string, value []byte) {
    c.mux.Lock() 
    defer c.mux.Unlock()
    
    c.cache[key] = cacheEntry{
        createdAt: time.Now(),
        val: value, 
    }
} 

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mux.Lock()
    defer c.mux.Unlock()

    v, ok := c.cache[key]
    if !ok {
        return nil, false
    }
    return v.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
        c.reap(time.Now().UTC(), interval)
    }
}

func (c *Cache) reap(now time.Time, last time.Duration) {
    c.mux.Lock()
    defer c.mux.Unlock()
    for k, v := range c.cache {
        if v.createdAt.Before(now.Add(-last)) {
            delete(c.cache, k)
        }
    }
}

func NewCache(interval time.Duration) Cache {
    c := Cache {
        cache:  make(map[string]cacheEntry),
        mux:    &sync.Mutex{},
    }

    go c.reapLoop(interval)

    return c
}



package gache

import (
	"sync"

	lru "gache/gache/leastRecentlyUsed"
)

type cache struct {
	muLock     sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.muLock.Lock()
	defer c.muLock.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.muLock.Lock()
	defer c.muLock.Unlock()
	if c.lru == nil {
		return
	}

	if value, ok := c.lru.Get(key); ok {
		return value.(ByteView), ok
	}

	return
}

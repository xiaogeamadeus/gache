package gache

import (
	"gache/gache/leastRecentlyUsed"
	"sync"
)

type cache struct {
	muLock     sync.Mutex
	lru        *leastRecentlyUsed.Cache
	cacheBytes int64
}

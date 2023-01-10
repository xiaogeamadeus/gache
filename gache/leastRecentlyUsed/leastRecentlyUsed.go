package leastRecentlyUsed

import (
	"container/list"
)

// Cache A LRU cache. Not safe for concurrent access.
type Cache struct {
	maxBytes                int64
	usedBytes               int64
	bidirectionalLinkedList *list.List
	storeMap                map[string]*list.Element
	onEvicted               func(key string, value Value)
}

// entry The node of bidirectional linked list
type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

// New use it to initialize cache.
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:                maxBytes,
		bidirectionalLinkedList: list.New(),
		storeMap:                make(map[string]*list.Element),
		onEvicted:               onEvicted,
	}
}

// Get look up a key's value.
func (cache *Cache) Get(key string) (value Value, ok bool) {
	if node, ok := cache.storeMap[key]; ok {
		cache.bidirectionalLinkedList.MoveToFront(node)
		nodeValueInMap := node.Value.(*entry)
		return nodeValueInMap.value, true
	}

	return
}

// RemoveOldest removes the oldest item.
func (cache *Cache) RemoveOldest() {
	node := cache.bidirectionalLinkedList.Back()
	if node != nil {
		cache.bidirectionalLinkedList.Remove(node)
		nodeValueInMap := node.Value.(*entry)
		delete(cache.storeMap, nodeValueInMap.key)
		cache.usedBytes -= int64(len(nodeValueInMap.key)) + int64(nodeValueInMap.value.Len())
		if cache.onEvicted != nil {
			cache.onEvicted(nodeValueInMap.key, nodeValueInMap.value)
		}
	}
}

// Add update or add new item into cache
func (cache *Cache) Add(key string, value Value) {
	if node, ok := cache.storeMap[key]; ok {
		cache.bidirectionalLinkedList.MoveToFront(node)
		nodeValueInMap := node.Value.(*entry)
		cache.usedBytes += int64(value.Len()) - int64(nodeValueInMap.value.Len())
		nodeValueInMap.value = value
	} else {
		node := cache.bidirectionalLinkedList.PushFront(&entry{key, value})
		cache.storeMap[key] = node
		cache.usedBytes += int64(len(key)) + int64(value.Len())
	}
	for cache.maxBytes != 0 && cache.maxBytes < cache.usedBytes {
		cache.RemoveOldest()
	}
}

// Len get the number of items in cache
func (cache *Cache) Len() int {
	return cache.bidirectionalLinkedList.Len()
}

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
func (c *Cache) Get(key string) (value Value, ok bool) {
	if node, ok := c.storeMap[key]; ok {
		c.bidirectionalLinkedList.MoveToFront(node)
		nodeValueInMap := node.Value.(*entry)
		return nodeValueInMap.value, true
	}

	return
}

// RemoveOldest removes the oldest item.
func (c *Cache) RemoveOldest() {
	node := c.bidirectionalLinkedList.Back()
	if node != nil {
		c.bidirectionalLinkedList.Remove(node)
		nodeValueInMap := node.Value.(*entry)
		delete(c.storeMap, nodeValueInMap.key)
		c.usedBytes -= int64(len(nodeValueInMap.key)) + int64(nodeValueInMap.value.Len())
		if c.onEvicted != nil {
			c.onEvicted(nodeValueInMap.key, nodeValueInMap.value)
		}
	}
}

// Add update or add new item into cache
func (c *Cache) Add(key string, value Value) {
	if node, ok := c.storeMap[key]; ok {
		c.bidirectionalLinkedList.MoveToFront(node)
		nodeValueInMap := node.Value.(*entry)
		c.usedBytes += int64(value.Len()) - int64(nodeValueInMap.value.Len())
		nodeValueInMap.value = value
	} else {
		node := c.bidirectionalLinkedList.PushFront(&entry{key, value})
		c.storeMap[key] = node
		c.usedBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.usedBytes {
		c.RemoveOldest()
	}
}

// Len get the number of items in cache
func (c *Cache) Len() int {
	return c.bidirectionalLinkedList.Len()
}

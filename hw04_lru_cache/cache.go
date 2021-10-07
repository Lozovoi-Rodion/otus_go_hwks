package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*listItem
	sync.Mutex
}

type cacheItem struct {
	Key   Key
	Value interface{}
}

func newCacheItem(key Key, value interface{}) *cacheItem {
	return &cacheItem{key, value}
}

func NewCache(capacity int) *lruCache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*listItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.Lock()
	defer c.Unlock()

	newItem := newCacheItem(key, value)
	if el, ok := c.items[key]; ok {
		el.Value = newItem
		c.queue.PushFront(el)
		return true
	}

	if c.capacity == c.queue.Len() {
		queueItem := c.queue.Back()
		c.queue.Remove(queueItem)
		cachedItem := queueItem.Value.(*cacheItem)

		delete(c.items, cachedItem.Key)
	}

	listItem := c.queue.PushFront(newItem)
	c.items[key] = listItem
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()

	if el, ok := c.items[key]; ok {
		c.queue.PushFront(el)
		item := el.Value.(*cacheItem)
		return item.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*listItem, c.capacity)
	c.queue = NewList()
}

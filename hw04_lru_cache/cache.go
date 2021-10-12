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

type LruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func newCacheItem(key Key, value interface{}) *cacheItem {
	return &cacheItem{key, value}
}

func NewCache(capacity int) *LruCache {
	return &LruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *LruCache) Set(key Key, value interface{}) bool {
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

		delete(c.items, cachedItem.key)
	}

	ListItem := c.queue.PushFront(newItem)
	c.items[key] = ListItem
	return false
}

func (c *LruCache) Get(key Key) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()

	if el, ok := c.items[key]; ok {
		c.queue.PushFront(el)
		item := el.Value.(*cacheItem)
		return item.value, true
	}
	return nil, false
}

func (c *LruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}

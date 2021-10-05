package hw04lrucache

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
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) *lruCache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*listItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if el, ok := c.items[key]; ok {
		el.Value = value
		c.queue.PushFront(el)
		return true
	}

	if c.capacity == c.queue.Len() {
		queueItem := c.queue.Back()
		c.queue.Remove(queueItem)
		for key := range c.items {
			if c.items[key].Value == queueItem.Value {
				delete(c.items, key)
			}
		}
	}
	newItem := c.queue.PushFront(value)
	c.items[key] = newItem

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if el, ok := c.items[key]; ok {
		c.queue.PushFront(el)
		return el.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*listItem, c.capacity)
	c.queue = NewList()
}

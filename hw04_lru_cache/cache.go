package hw04_lru_cache //nolint:golint,stylecheck
import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    sync.Mutex
}

type cacheItem struct {
	Key   Key
	Value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, ok := c.items[key]; ok {
		item.Value = cacheItem{Key: key, Value: value}
		c.queue.MoveToFront(item)

		return true
	}

	if c.capacity == c.queue.Len() {
		last := c.queue.Back()

		delete(c.items, last.Value.(cacheItem).Key)
		c.queue.Remove(last)
	}

	c.queue.PushFront(cacheItem{Key: key, Value: value})
	c.items[key] = c.queue.Front()

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, ok := c.items[key]; ok {
		value := item.Value.(cacheItem).Value
		c.queue.MoveToFront(item)

		return value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = &list{}
	c.items = map[Key]*ListItem{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		items:    map[Key]*ListItem{},
		queue:    &list{},
	}
}

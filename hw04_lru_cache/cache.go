package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                              // Очистить кэш
}

type lruCache struct {
	// Place your code here:
	// - capacity
	// - queue
	// - items
}

type cacheItem struct {
	// Place your code here
}

func (c lruCache) Set(key Key, value interface{}) bool {
	return false
}

func (c lruCache) Get(key Key) (interface{}, bool) {
	return nil, false
}

func (c lruCache) Clear() {

}

func NewCache(capacity int) Cache {
	return &lruCache{}
}

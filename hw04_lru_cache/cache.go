package hw04lrucache

import "sync"

type Key string
type Value string

type Cache interface {
	Set(key Key, value interface{}, c sync.Mutex) bool
	Get(key Key, c sync.Mutex) (interface{}, bool)
	Clear()
}

type lruCache struct {
	//	Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

// Set - метод для добавления значения в кэш по ключу
func Set(key Key, value Value, mu sync.Mutex) bool {
	mu.Lock()         // блокируем мьютекс для синхронизации
	defer mu.Unlock() // освобождаем мьютекс после выполнения операции

	item, exists := mu.cmap[key]
	if !exists {
		item = mu.list.PushFront(CacheItem{Key: key, Value: value})
		c.cmap[key] = item
		if mu.list.Len() > mu.Capacity {
			// если размер списка превышает емкость кэша, удаляем последний элемент
			mu.list.Remove(mu.list.Back())
			delete(mu.cmap, mu.list.Back().Value.Key)
		}
		return true
	}

	// если элемент уже существует, обновляем его значение и перемещаем в начало списка
	item.Value = value
	mu.list.MoveToFront(item)
	return false
}

// Get - метод для получения значения из кэша по ключу
func Get(key Key, c sync.Mutex) (Value, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, exists := c.cmap[key]
	if exists {
		c.list.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

// Clear - метод для очистки кэша
func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, item := range c.list.List {
		delete(c.cmap, item.Value.Key)
	}
	c.cmap = make(cmap[Key] * ListItem)
	c.list = NewList()
}

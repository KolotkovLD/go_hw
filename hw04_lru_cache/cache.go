package hw04lrucache

import "sync"

type Key string
type Value string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	//	Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

// Set - метод для добавления значения в кэш по ключу
func (c *lruCache) Set(key Key, value Value) bool {
	c.mu.Lock()         // блокируем мьютекс для синхронизации
	defer c.mu.Unlock() // освобождаем мьютекс после выполнения операции

	item, exists := c.cmap[key]
	if !exists {
		item = c.list.PushFront(CacheItem{Key: key, Value: value})
		c.cmap[key] = item
		if c.list.Len() > c.Capacity {
			// если размер списка превышает емкость кэша, удаляем последний элемент
			c.list.Remove(c.list.Back())
			delete(c.cmap, c.list.Back().Value.Key)
		}
		return true
	}

	// если элемент уже существует, обновляем его значение и перемещаем в начало списка
	item.Value = value
	c.list.MoveToFront(item)
	return false
}

// Get - метод для получения значения из кэша по ключу
func (c *lruCache) Get(key Key) (Value, bool) {
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

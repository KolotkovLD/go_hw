package hw04lrucache

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

func NewCache(capacity int) lruCache {
	return lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

// Set - метод для добавления значения в кэш по ключу
func (c *lruCache) Set(key Key, value interface{}) bool {

	item, exists := c.items[key]
	if !exists {
		item = c.queue.PushFront(value)
		c.items[key] = item
		if c.queue.Len() > c.capacity {
			// если размер списка превышает емкость кэша, удаляем последний элемент
			c.queue.Remove(c.queue.Back())
		}
		return false
	}

	// если элемент уже существует, обновляем его значение и перемещаем в начало списка
	item.Value = value
	c.queue.MoveToFront(item)
	return true
}

// Get - метод для получения значения из кэша по ключу
func (c *lruCache) Get(key Key) (interface{}, bool) {

	item, exists := c.items[key]
	if exists {
		c.queue.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

//// Clear - метод для очистки кэша
//func (c *lruCache) Clear() {
//
//	for _, item := range c.queue {
//		delete(c.cmap, item.Value.Key)
//	}
//	c.cmap = make(cmap[Key] * ListItem)
//	c.list = NewList()
//}

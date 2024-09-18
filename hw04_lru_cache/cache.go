package hw04lrucache

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
}

type lruCache struct {
	//	Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[string]*ListItem
}

type entry struct {
	key   string
	value interface{}
}

func NewCache(capacity int) lruCache {
	return lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[string]*ListItem, capacity),
	}
}

// Set - метод для добавления значения в кэш по ключу.
func (c *lruCache) Set(key string, value interface{}) bool {
	if item, exist := c.items[key]; exist {
		// если элемент уже существует, обновляем его значение и перемещаем в начало списка
		item.Value.(*entry).value = value
		item = c.queue.MoveToFront(item)
		c.items[key] = item
		return true
	}
	if c.queue.Len() == c.capacity {
		// если размер списка превышает емкость кэша, удаляем последний элемент
		leastUsedValue := c.queue.Back()
		c.queue.Remove(c.queue.Back())
		delete(c.items, leastUsedValue.Value.(*entry).key)
	}

	newEntry := &entry{key, value}
	item := c.queue.PushFront(newEntry)
	c.items[key] = item
	return false
}

// Get - метод для получения значения из кэша по ключу.
func (c *lruCache) Get(key string) (interface{}, bool) {
	if item, exists := c.items[key]; exists {
		item = c.queue.MoveToFront(item)
		c.items[key] = item
		return item.Value.(*entry).value, true
	}
	return nil, false
}

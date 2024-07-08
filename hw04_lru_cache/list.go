package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head *ListItem
	tail *ListItem
	len  int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Next: l.head}
	if l.head != nil {
		l.head.Prev = newItem
	}
	l.head = newItem
	if l.tail == nil {
		l.tail = newItem
	}
	l.len++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Prev: l.tail}
	if l.tail != nil {
		l.tail.Next = newItem
	}
	l.tail = newItem
	if l.head == nil {
		l.head = newItem
	}
	l.len++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.head = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}

	i.Prev = nil
	i.Next = nil
	i.Value = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.head {
		return
	}
	newFrontValue := i.Value
	l.Remove(i)
	l.PushFront(newFrontValue)
}

func NewList() List {
	return &list{}
}

//type Cashe interface {
//	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу.
//	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу.
//	Clear()                              // Очистить кэш.
//}
//
//type CasheItem struct {
//	capasity int
//	order &list
//	keys []string
//}

//// Key - тип ключа для кэша
//type CKey string
//
//// Value - обобщенный тип значения
//type Value interface{}
//
//// CacheItem - структура для элемента кэша
//type CacheItem struct {
//	Key   CKey   // ключ
//	Value Value // значение
//}
//
//// Cache - структура для представления кэша
//type Cache struct {
//	Capacity int // ёмкость кэша
//	mu       sync.Mutex // мьютекс для синхронизации доступа к кэшу
//	list     *List // список последних использованных элементов
//	cmap      map[Key]*ListItem // словарь ключ-элемент
//}
//
//// NewCache - функция для создания нового экземпляра кэша
//func NewCache(capacity int) *Cache {
//	return &Cache{Capacity: capacity, list: NewList(), cmap: make(cmap[Key]*ListItem)}
//}
//
//// Set - метод для добавления значения в кэш по ключу
//func (c *Cache) Set(key Key, value Value) bool {
//	c.mu.Lock() // блокируем мьютекс для синхронизации
//	defer c.mu.Unlock() // освобождаем мьютекс после выполнения операции
//
//	item, exists := c.cmap[key]
//	if !exists {
//		item = c.list.PushFront(CacheItem{Key: key, Value: value})
//		c.cmap[key] = item
//		if c.list.Len() > c.Capacity {
//			// если размер списка превышает емкость кэша, удаляем последний элемент
//			c.list.Remove(c.list.Back())
//			delete(c.cmap, c.list.Back().Value.Key)
//		}
//		return true
//	}
//
//	// если элемент уже существует, обновляем его значение и перемещаем в начало списка
//	item.Value = value
//	c.list.MoveToFront(item)
//	return false
//}
//
//// Get - метод для получения значения из кэша по ключу
//func (c *Cache) Get(key Key) (Value, bool) {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//
//	item, exists := c.cmap[key]
//	if exists {
//		c.list.MoveToFront(item)
//		return item.Value, true
//	}
//	return nil, false
//}
//
//// Clear - метод для очистки кэша
//func (c *Cache) Clear() {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//
//	for _, item := range c.list.List {
//		delete(c.cmap, item.Value.Key)
//	}
//	c.cmap = make(cmap[Key]*ListItem)
//	c.list = NewList()
//}

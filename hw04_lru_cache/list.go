package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem) *ListItem
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
	newItem := &ListItem{Value: v, Prev: l.head}

	if l.head == nil {
		l.head = newItem
		l.tail = newItem
	} else {
		currentItem := l.head
		newItem.Prev = currentItem
		currentItem.Next = newItem
		l.head = newItem
	}

	l.len++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Next: l.tail}
	if l.head == nil {
		l.head = newItem
		l.tail = newItem
	} else {
		currentItem := l.tail
		newItem.Next = currentItem
		currentItem.Prev = newItem
		l.tail = newItem
	}

	l.len++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.tail = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.head = i.Prev
	}

	i.Prev = nil
	i.Next = nil
	i.Value = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) *ListItem {
	if i == l.head {
		return l.head
	}
	newFrontValue := i.Value
	l.Remove(i)
	return l.PushFront(newFrontValue)
}

func NewList() List {
	return &list{}
}

package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	isEmpty() bool
}

type ListItem struct {
	Prev  *ListItem
	Value interface{}
	Next  *ListItem
}

func NewListItem(value interface{}) *ListItem {
	return &ListItem{Value: value}
}

type LinkedList struct {
	head *ListItem
	tail *ListItem
	len  int
}

func NewList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) Front() *ListItem {
	return l.head
}

func (l *LinkedList) Back() *ListItem {
	return l.tail
}

func (l *LinkedList) PushFront(value interface{}) *ListItem {
	item := NewListItem(value)
	if l.isEmpty() {
		l.head = item
		l.tail = item
	} else {
		item.Next = l.head
		l.head.Prev = item
		l.head = item
	}
	l.len++
	return item
}

func (l *LinkedList) PushBack(value interface{}) *ListItem {
	item := NewListItem(value)
	if l.isEmpty() {
		l.head = item
		l.tail = item
	} else {
		l.tail.Next = item
		item.Prev = l.tail
		l.tail = item
	}
	l.len++
	return item
}

func (l *LinkedList) Remove(item *ListItem) {
	switch item {
	case l.head:
		if l.len == 1 {
			l.head = nil
			l.tail = nil
		} else {
			item.Next.Prev = nil
			l.head = item.Next
		}
	case l.tail:
		item.Prev.Next = nil
		l.tail = item.Prev
	default:
		item.Prev.Next = item.Next
		item.Next.Prev = item.Prev
	}

	l.len--
	item.Prev = nil
	item.Next = nil
}

func (l *LinkedList) MoveToFront(item *ListItem) {
	if item == l.head {
		return
	}

	if l.len == 1 && item == l.tail {
		return
	}

	item.Prev.Next = item.Next
	if item.Next != nil {
		item.Next.Prev = item.Prev
	}

	item.Next = l.head
	item.Prev = nil
	l.head.Prev = item
	l.head = item
}

func (l *LinkedList) isEmpty() bool {
	return l.len == 0
}

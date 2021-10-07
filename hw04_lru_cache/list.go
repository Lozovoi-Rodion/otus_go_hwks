package hw04lrucache

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(value interface{}) *listItem
	PushBack(value interface{}) *listItem
	Remove(n *listItem)
	MoveToFront(n *listItem)
	isEmpty() bool
}

type listItem struct {
	Prev  *listItem
	Value interface{}
	Next  *listItem
}

func NewListItem(value interface{}) *listItem {
	return &listItem{Value: value}
}

type doubleLinkedList struct {
	head *listItem
	tail *listItem
	len  int
}

func NewList() *doubleLinkedList {
	return &doubleLinkedList{}
}

func (l *doubleLinkedList) Len() int {
	return l.len
}

func (l *doubleLinkedList) Front() *listItem {
	return l.head
}

func (l *doubleLinkedList) Back() *listItem {
	return l.tail
}

func (l *doubleLinkedList) PushFront(value interface{}) *listItem {
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

func (l *doubleLinkedList) PushBack(value interface{}) *listItem {
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

func (l *doubleLinkedList) Remove(item *listItem) {
	if l.head == item {
		if l.len == 1 {
			l.head = nil
			l.tail = nil
		} else {
			item.Next.Prev = nil
			l.head = item.Next
		}
	} else if l.tail == item {
		item.Prev.Next = nil
		l.tail = item.Prev
	} else {
		item.Prev.Next = item.Next
		item.Next.Prev = item.Prev
	}
	l.len--
	item.Prev = nil
	item.Next = nil
}

func (l *doubleLinkedList) MoveToFront(item *listItem) {
	if item == l.head || l.len == 1 && item == l.tail {
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

func (l *doubleLinkedList) isEmpty() bool {
	return l.len == 0
}

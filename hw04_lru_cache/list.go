package hw04_lru_cache //nolint:golint,stylecheck

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
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
		Next:  l.front,
	}

	if l.len > 0 {
		l.front.Prev = item
	} else {
		l.back = item
	}

	l.front = item
	l.len++

	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
		Prev:  l.back,
	}

	if l.len > 0 {
		l.back.Next = item
	} else {
		l.front = item
	}

	l.back = item
	l.len++

	return item
}

func (l *list) Remove(i *ListItem) {
	var (
		next = i.Next
		prev = i.Prev
	)

	if next != nil {
		next.Prev = prev
	} else {
		l.back = prev
	}

	if prev != nil {
		prev.Next = next
	} else {
		l.front = next
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil || l.len == 1 {
		return
	}

	i.Prev.Next = i.Next

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}

	i.Prev = nil
	i.Next = l.front

	l.front.Prev = i
	l.front = i
}

func NewList() List {
	return &list{}
}

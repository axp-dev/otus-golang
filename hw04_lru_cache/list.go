package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *ListItem                  // первый элемент списка
	Back() *ListItem                   // последний элемент списка
	PushFront(v interface{}) *ListItem // добавить значение в начало
	PushBack(v interface{}) *ListItem  // добавить значение в конец
	Remove(i *ListItem)                // удалить элемент
	MoveToFront(i *ListItem)           // переместить элемент в начало
}

type ListItem struct {
	Value interface{} // значение
	Next  *ListItem   // следующий элемент
	Prev  *ListItem   // предыдущий элемент
}

type list struct {
	// Place your code here
}

func (l list) Len() int {
	return 0
}

func (l list) Front() *ListItem {
	return nil
}

func (l list) Back() *ListItem {
	return nil
}

func (l list) PushFront(v interface{}) *ListItem {
	return nil
}

func (l list) PushBack(v interface{}) *ListItem {
	return nil
}

func (l list) Remove(i *ListItem) {

}

func (l list) MoveToFront(i *ListItem) {

}

func NewList() List {
	return &list{}
}

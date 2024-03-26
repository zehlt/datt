package datt

type node[T any] struct {
	value T
	next  *node[T]
}

type LinkedList[T any] struct {
	front *node[T]
	back  *node[T]
	len   int
}

func (l *LinkedList[T]) PushFront(v T) {
	new := &node[T]{value: v}

	if l.len == 0 {
		l.front = new
		l.back = new
		l.len++
		return
	}

	previousFront := l.front
	new.next = previousFront
	l.front = new

	l.len++
}

func (l *LinkedList[T]) PushBack(v T) {
	new := &node[T]{value: v}

	if l.len == 0 {
		l.front = new
		l.back = new
		l.len++
		return
	}

	previousBack := l.back
	previousBack.next = new
	l.back = new
	l.len++
}

func (l *LinkedList[T]) PopFront() (T, bool) {
	if l.len == 0 {
		var d T
		return d, false
	}

	if l.len == 1 {
		front := l.front
		l.front = nil
		l.back = nil
		l.len--
		return front.value, true
	}

	previous := l.front
	l.front = l.front.next
	l.len--
	return previous.value, true
}

func (l *LinkedList[T]) Front() (T, bool) {
	if l.len == 0 {
		var d T
		return d, false
	}

	return l.front.value, true
}

func (l *LinkedList[T]) Len() int {
	return l.len
}

func (l *LinkedList[T]) Do(f func(v T)) {
	current := l.front

	for current != nil {
		f(current.value)
		current = current.next
	}
}

func (l *LinkedList[T]) Clear() {
	l.len = 0
	l.front = nil
	l.back = nil
}

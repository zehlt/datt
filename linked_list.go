package datt

type node[T any] struct {
	value T
	next  *node[T]
}

type linkedList[T any] struct {
	front *node[T]
	back  *node[T]
	len   int
}

// func (l *linkedList[T]) PushFront(v T) {
// 	if l.len == 0 {
// 		new := &node[T]{value: v}
// 		l.front = new
// 		l.back = new
// 	}
// 	l.len++
// }

func (l *linkedList[T]) PushBack(v T) {
	if l.len == 0 {
		new := &node[T]{value: v}
		l.front = new
		l.back = new
		l.len++
		return
	}

	newNode := &node[T]{value: v}
	previousBack := l.back
	previousBack.next = newNode
	l.back = newNode
	l.len++
}

func (l *linkedList[T]) PopFront() (T, bool) {
	if l.len == 0 {
		var d T
		return d, false
	}

	if l.len == 1 {
		front := l.front
		l.front = nil
		l.len--
		return front.value, true
	}

	previous := l.front
	l.front = l.front.next
	l.len--
	return previous.value, true
}

func (l *linkedList[T]) Front() (T, bool) {
	if l.len == 0 {
		var d T
		return d, false
	}

	return l.front.value, true
}

func (l *linkedList[T]) Len() int {
	return l.len
}

func (l *linkedList[T]) Do(f func(v T)) {
	current := l.front

	for current != nil {
		f(current.value)
		current = current.next
	}
}

func (l *linkedList[T]) Clear() {
	l.len = 0
	l.front = nil
	l.back = nil
}

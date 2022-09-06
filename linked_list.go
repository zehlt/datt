package datt

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyList       = errors.New("empty list")
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrNotFound        = errors.New("data not found")
)

type node[T any] struct {
	Data T
	Next *node[T]
}

type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	lenght int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head:   nil,
		tail:   nil,
		lenght: 0,
	}
}

func (l *LinkedList[T]) PushHead(data T) {
	previousHead := l.head

	newHead := &node[T]{
		Data: data,
		Next: previousHead,
	}

	if l.lenght <= 0 {
		l.head = newHead
		l.tail = newHead
	} else {
		l.head = newHead
	}

	l.lenght++
}

func (l *LinkedList[T]) PushTail(data T) {
	newTail := &node[T]{
		Data: data,
		Next: nil,
	}

	if l.lenght <= 0 {
		l.head = newTail
		l.tail = newTail
	} else {
		previousTail := l.tail
		previousTail.Next = newTail
		l.tail = newTail
	}

	l.lenght++
}

func (l *LinkedList[T]) PopHead() (T, error) {
	if l.lenght <= 0 {
		var empty T
		return empty, ErrEmptyList
	}

	previousHead := l.head

	if l.lenght == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = previousHead.Next
	}

	l.lenght--
	return previousHead.Data, nil
}

func (l *LinkedList[T]) PopTail() (T, error) {
	if l.lenght <= 0 {
		var empty T
		return empty, ErrEmptyList
	}

	previousTail := l.tail

	if l.lenght == 1 {
		l.head = nil
		l.tail = nil
	} else {
		penultimate := l.head
		for i := 0; i < l.lenght-2; i++ {
			penultimate = penultimate.Next
		}
		penultimate.Next = nil
		l.tail = penultimate
	}

	l.lenght--

	return previousTail.Data, nil
}

func (l *LinkedList[T]) PeekAt(index int) (T, error) {
	var d T

	if l.lenght <= 0 {
		return d, ErrEmptyList
	}

	if index >= l.lenght {
		return d, ErrIndexOutOfRange
	}

	if index == 0 {
		return l.PeekHead()
	}

	if index == l.lenght-1 {
		return l.PeekTail()
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Data, nil
}

func (l *LinkedList[T]) PopAt(index int) (T, error) {
	var d T

	if l.lenght <= 0 {
		return d, ErrEmptyList
	}

	if index >= l.lenght {
		return d, ErrIndexOutOfRange
	}

	if index == 0 {
		return l.PopHead()
	}

	if index == l.lenght-1 {
		return l.PopTail()
	}

	currentnode := l.head
	var previousnode *node[T]

	for i := 0; i < index; i++ {
		previousnode = currentnode
		currentnode = currentnode.Next
	}

	previousnode.Next = currentnode.Next
	l.lenght--
	return currentnode.Data, nil
}

func (l *LinkedList[T]) PeekHead() (T, error) {
	if l.lenght <= 0 {
		var d T
		return d, ErrEmptyList
	}

	return l.head.Data, nil
}

func (l *LinkedList[T]) PeekTail() (T, error) {
	if l.lenght <= 0 {
		var d T
		return d, ErrEmptyList
	}

	return l.tail.Data, nil
}

func (l *LinkedList[T]) Length() int {
	return l.lenght
}

func (l *LinkedList[T]) Iterator(callback func(value T) bool) {
	currentNode := l.head

	for i := 0; i < l.lenght; i++ {
		shouldStop := callback(currentNode.Data)
		if shouldStop {
			break
		}

		currentNode = currentNode.Next
	}
}

func (l *LinkedList[T]) String() string {
	var builder strings.Builder

	l.Iterator(func(value T) bool {
		builder.WriteString(fmt.Sprintf("[%v] -> ", value))
		return false
	})

	return builder.String()
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.lenght <= 0
}

func (l *LinkedList[T]) Clear() {
	if l.lenght <= 0 {
		return
	}

	// TODO: checf if data leak may need to null all data
	l.head = nil
	l.tail = nil
	l.lenght = 0
}

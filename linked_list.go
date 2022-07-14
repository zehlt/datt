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

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

type linkedList[T comparable] struct {
	head   *Node[T]
	tail   *Node[T]
	lenght int
}

func NewLinkedList[T comparable]() linkedList[T] {
	return linkedList[T]{
		head:   nil,
		tail:   nil,
		lenght: 0,
	}
}

func (l *linkedList[T]) PushHead(data T) {
	previousHead := l.head

	newHead := &Node[T]{
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

func (l *linkedList[T]) PushTail(data T) {
	newTail := &Node[T]{
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

func (l *linkedList[T]) PopHead() (T, error) {
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

func (l *linkedList[T]) PopTail() (T, error) {
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

func (l *linkedList[T]) PeekAt(index int) (T, error) {
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

// TODO: add insert at

func (l *linkedList[T]) Pop(data T) (T, error) {
	var d T

	if l.lenght <= 0 {
		return d, ErrEmptyList
	}

	currentNode := l.head
	var previousNode *Node[T]

	for i := 0; i < l.lenght; i++ {
		previousNode = currentNode

		if currentNode.Data == data {
			previousNode.Next = currentNode.Next
			l.lenght--

			return currentNode.Data, nil
		}
		currentNode = currentNode.Next
	}

	return d, ErrNotFound
}

func (l *linkedList[T]) IndexOf(data T) (int, error) {
	if l.lenght <= 0 {
		return 0, ErrEmptyList
	}

	currentNode := l.head

	for i := 0; i < l.lenght; i++ {
		if currentNode.Data == data {
			return i, nil
		}

		currentNode = currentNode.Next
	}

	return 0, ErrNotFound
}

func (l *linkedList[T]) Contains(data T) bool {
	if l.lenght <= 0 {
		return false
	}

	currentNode := l.head

	for i := 0; i < l.lenght; i++ {
		if currentNode.Data == data {
			return true
		}

		currentNode = currentNode.Next
	}

	return false
}

func (l *linkedList[T]) PopAt(index int) (T, error) {
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

	currentNode := l.head
	var previousNode *Node[T]

	for i := 0; i < index; i++ {
		previousNode = currentNode
		currentNode = currentNode.Next
	}

	previousNode.Next = currentNode.Next
	l.lenght--
	return currentNode.Data, nil
}

func (l *linkedList[T]) PeekHead() (T, error) {
	if l.lenght <= 0 {
		var d T
		return d, ErrEmptyList
	}

	return l.head.Data, nil
}

func (l *linkedList[T]) PeekTail() (T, error) {
	if l.lenght <= 0 {
		var d T
		return d, ErrEmptyList
	}

	return l.tail.Data, nil
}

func (l *linkedList[T]) Length() int {
	return l.lenght
}

func (l *linkedList[T]) String() string {
	var builder strings.Builder

	currentNode := l.head
	for currentNode != nil {
		builder.WriteString(fmt.Sprintf("[%v] -> ", currentNode.Data))
		currentNode = currentNode.Next
	}

	return builder.String()
}

func (l *linkedList[T]) Clear() {
	if l.lenght <= 0 {
		return
	}

	// TODO: checf if data leak may need to null all data
	l.head = nil
	l.tail = nil
	l.lenght = 0
}

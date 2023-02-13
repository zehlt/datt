package datt

type StackLinkedList[T any] struct {
	ll linkedList[T]
}

// O(1)
func (s *StackLinkedList[T]) Len() int {
	return s.ll.Len()
}

// O(1)
func (s *StackLinkedList[T]) Push(v T) {
	s.ll.PushFront(v)
}

// O(1)
func (s *StackLinkedList[T]) Peek() (T, bool) {
	return s.ll.Front()
}

// O(1)
func (s *StackLinkedList[T]) Pop() (T, bool) {
	return s.ll.PopFront()
}

// O(N)
func (s *StackLinkedList[T]) Do(f func(v T)) {
	s.ll.Do(f)
}

// O(1)
func (s *StackLinkedList[T]) Clear() {
	s.ll.Clear()
}

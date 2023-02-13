package datt

type QueueLinkedList[T any] struct {
	ll linkedList[T]
}

// O(1)
func (q *QueueLinkedList[T]) Len() int {
	return q.ll.Len()
}

// O(1)
func (q *QueueLinkedList[T]) Enqueue(v T) {
	q.ll.PushBack(v)
}

// O(1)
func (q *QueueLinkedList[T]) Dequeue() (T, bool) {
	return q.ll.PopFront()
}

// O(1)
func (q *QueueLinkedList[T]) Peek() (T, bool) {
	return q.ll.Front()
}

// O(N)
func (q *QueueLinkedList[T]) Do(f func(v T)) {
	q.ll.Do(f)
}

// O(1)
func (q *QueueLinkedList[T]) Clear() {
	q.ll.Clear()
}

package datt

type QueueCircularArray[T any] struct {
}

func (q *QueueCircularArray[T]) Len() int {
	return 0
}

func (q *QueueCircularArray[T]) Enqueue(v T) {
}

func (q *QueueCircularArray[T]) Dequeue() (T, bool) {
	var d T
	return d, false
}

func (q *QueueCircularArray[T]) Peek() (T, bool) {
	var d T
	return d, false
}

func (q *QueueCircularArray[T]) Do(f func(v T)) {
}

func (q *QueueCircularArray[T]) Clear() {
}

package datt

type queue[T comparable] struct {
	l linkedList[T]
}

func NewQueue[T comparable]() queue[T] {
	return queue[T]{
		l: NewLinkedList[T](),
	}
}

func (q *queue[T]) PeekFront() (T, error) {
	return q.l.PeekHead()
}

func (q *queue[T]) PeekBack() (T, error) {
	return q.l.PeekTail()
}

func (q *queue[T]) Enqueue(data T) {
	q.l.PushTail(data)
}

func (q *queue[T]) Dequeue() (T, error) {
	return q.l.PopHead()
}

func (q *queue[T]) Length() int {
	return q.l.Length()
}

func (q *queue[T]) IsEmpty() bool {
	return q.l.IsEmpty()
}

func (q *queue[T]) Clear() {
	q.l.Clear()
}

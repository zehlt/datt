package datt

type Queue[T any] struct {
	l *LinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		l: NewLinkedList[T](),
	}
}

func (q *Queue[T]) PeekFront() (T, error) {
	return q.l.PeekHead()
}

func (q *Queue[T]) PeekBack() (T, error) {
	return q.l.PeekTail()
}

func (q *Queue[T]) Enqueue(data T) {
	q.l.PushTail(data)
}

func (q *Queue[T]) Dequeue() (T, error) {
	return q.l.PopHead()
}

func (q *Queue[T]) Length() int {
	return q.l.Length()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.l.IsEmpty()
}

func (q *Queue[T]) Clear() {
	q.l.Clear()
}

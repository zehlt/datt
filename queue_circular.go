package datt

type QueueCircularArray[T any] struct {
	front int
	next  int
	arr   []T
	cap   int
	len   int
}

func NewQueueCircularArray[T any](initCap int) QueueCircularArray[T] {
	return QueueCircularArray[T]{
		cap: initCap,
		arr: make([]T, initCap),
	}
}

func (q *QueueCircularArray[T]) Enqueue(v T) {
	if q.len == q.cap {
		q.growComplex()
	}

	if q.next == q.cap {
		q.next = 0
	}

	q.arr[q.next] = v
	q.next++
	q.len++
}

func (q *QueueCircularArray[T]) growComplex() {
	newArr := make([]T, q.cap*2)

	for i := 0; i < q.len; i++ {
		if q.front == q.cap {
			q.front = 0
		}

		newArr[i] = q.arr[q.front]
		q.front++
	}

	q.arr = newArr
	q.cap *= 2
	q.front = 0
	q.next = q.len
}

func (q *QueueCircularArray[T]) Dequeue() (T, bool) {
	var v T
	if q.front == q.next {
		return v, false
	}

	if q.front == q.cap {
		q.front = 0
	}

	v = q.arr[q.front]
	q.front++
	q.len--
	return v, true
}

func (q *QueueCircularArray[T]) Peek() (T, bool) {
	if q.front == q.next {
		if q.len == q.cap {
			return q.arr[q.front], true
		}

		var d T
		return d, false
	}

	return q.arr[q.front], true
}

func (q *QueueCircularArray[T]) Do(f func(v T)) {
	start := q.front

	for i := 0; i < q.len; i++ {
		if start == q.cap {
			start = 0
		}
		f(q.arr[start])
		start++
	}
}

func (q *QueueCircularArray[T]) Len() int {
	return q.len
}

func (q *QueueCircularArray[T]) Clear() {
	q.next = 0
	q.front = 0
	q.len = 0
}

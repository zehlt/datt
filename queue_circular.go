package datt

type QueueCircularArray[T any] struct {
	front int
	rear  int
	len   int
	cap   int
	arr   []T
}

const defaultInitialCapacity = 4

// O(1)
func (q *QueueCircularArray[T]) Enqueue(v T) {
	if q.len == q.cap {
		if q.cap <= 0 {
			q.cap = defaultInitialCapacity
			q.arr = make([]T, defaultInitialCapacity)
		} else {
			q.grow()
		}
	}

	if q.len > 0 {
		q.rear++

		if q.rear == q.cap {
			q.rear = 0
		}
	}

	q.arr[q.rear] = v
	q.len++
}

func (q *QueueCircularArray[T]) grow() {
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
	q.rear = q.len - 1
}

// O(1)
func (q *QueueCircularArray[T]) Dequeue() (T, bool) {
	var v T
	if q.len == 0 {
		return v, false
	}

	v = q.arr[q.front]
	if q.len > 1 {
		q.front++
	}

	q.len--
	return v, true
}

// O(1)
func (q *QueueCircularArray[T]) Peek() (T, bool) {
	if q.len == 0 {
		var v T
		return v, false
	}

	return q.arr[q.front], true
}

// O(N)
func (q *QueueCircularArray[T]) Do(f func(v T)) {
	for i := 0; i < q.len; i++ {
		if q.front == q.cap {
			q.front = 0
		}

		f(q.arr[q.front])
		q.front++
	}
}

// O(1)
func (q *QueueCircularArray[T]) Len() int {
	return q.len
}

// O(1)
func (q *QueueCircularArray[T]) Clear() {
	q.len = 0
	q.rear = 0
	q.front = 0
}

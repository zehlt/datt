package datt

type PriorityQueue[T any] struct {
	compare CompareFunc[T]
	len     int
	cap     int
	arr     []T
}

func NewPriorityQueue[T any](compare CompareFunc[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		compare: compare,
		cap:     1,
		arr:     make([]T, 1),
	}
}

// O(log(N))
func (q *PriorityQueue[T]) Enqueue(v T) {
	q.len++
	q.ensureCapacity()

	q.arr[q.getLastIndex()] = v
	q.bubbleUp(q.getLastIndex())
}

// O(log(N))
func (q *PriorityQueue[T]) Dequeue() (T, bool) {
	if q.len == 0 {
		var d T
		return d, false
	}

	if q.len == 1 {
		out := q.arr[0]
		q.len--
		return out, true
	}

	out := q.arr[0]
	q.arr[0] = q.arr[q.getLastIndex()]
	q.len--
	q.bubbleDown(0)

	return out, true
}

// O(1)
func (q *PriorityQueue[T]) Peek() (T, bool) {
	if q.len == 0 {
		var d T
		return d, false
	}

	return q.arr[0], true
}

// O(1)
func (q *PriorityQueue[T]) Len() int {
	return q.len
}

// O(N)
func (q *PriorityQueue[T]) Do(f func(v T)) {
	for i := 0; i < q.len; i++ {
		f(q.arr[i])
	}
}

// O(1)
func (q *PriorityQueue[T]) Clear() {
	q.len = 0
}

func (q *PriorityQueue[T]) bubbleUp(currentIndex int) {
	if currentIndex == 0 {
		return
	}

	currentValue := q.arr[currentIndex]
	parentIndex := q.getParentIndex(currentIndex)
	parentValue := q.arr[parentIndex]

	result := q.compare(currentValue, parentValue)

	switch result {
	case LOWER:
		q.arr[currentIndex] = parentValue
		q.arr[parentIndex] = currentValue

		q.bubbleUp(parentIndex)
	}
}

func (q *PriorityQueue[T]) bubbleDown(currentIndex int) {
	currentValue := q.arr[currentIndex]

	leftIndex := q.getLeftChildIndex(currentIndex)
	rightIndex := q.getRightChildIndex(currentIndex)

	if rightIndex >= q.len {
		if leftIndex >= q.len {
			return
		}

		leftValue := q.arr[leftIndex]

		result := q.compare(currentValue, leftValue)
		switch result {
		case HIGHER:
			q.arr[currentIndex] = leftValue
			q.arr[leftIndex] = currentValue
		}
		return
	}

	leftValue := q.arr[q.getLeftChildIndex(currentIndex)]
	rightValue := q.arr[q.getRightChildIndex(currentIndex)]
	result := q.compare(leftValue, rightValue)

	var smallestIndex int
	switch result {
	case LOWER, EQUAL:
		smallestIndex = q.getLeftChildIndex(currentIndex)
	case HIGHER:
		smallestIndex = q.getRightChildIndex(currentIndex)
	}

	smallestValue := q.arr[smallestIndex]
	result = q.compare(currentValue, smallestValue)

	switch result {
	case HIGHER:
		q.arr[currentIndex] = smallestValue
		q.arr[smallestIndex] = currentValue
		q.bubbleDown(smallestIndex)
	}
}

func (q *PriorityQueue[T]) getParentIndex(current int) int {
	return (current - 1) / 2
}

func (q *PriorityQueue[T]) getLeftChildIndex(current int) int {
	return (current * 2) + 1
}

func (q *PriorityQueue[T]) getRightChildIndex(current int) int {
	return (current * 2) + 2
}

func (q *PriorityQueue[T]) getLastIndex() int {
	return q.len - 1
}

func (q *PriorityQueue[T]) ensureCapacity() {
	if q.len == q.cap+1 {
		newCap := q.cap + (q.cap * 2)
		n := make([]T, newCap)
		copy(n, q.arr)
		q.arr = n
		q.cap = newCap
	}
}

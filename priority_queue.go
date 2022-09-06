package datt

func NewPriorityQueue[T any](compare func(current T, other T) CompareResult) PriorityQueue[T] {
	return &priorityQueue[T]{
		binaryHeap: NewBinaryHeap(compare),
	}
}

type PriorityQueue[T any] interface {
	Push(value T)
	Pop() (T, error)
	Peek() (T, error)
	Length() int
	Clear()
	String() string
}

type priorityQueue[T any] struct {
	binaryHeap *BinaryHeap[T]
}

func (pq *priorityQueue[T]) Push(value T) {
	pq.binaryHeap.Push(value)
}

func (pq *priorityQueue[T]) Pop() (T, error) {
	return pq.binaryHeap.Pop()
}

func (pq *priorityQueue[T]) Peek() (T, error) {
	return pq.binaryHeap.PeekFront()
}

func (pq *priorityQueue[T]) Length() int {
	return pq.binaryHeap.Length()
}

func (pq *priorityQueue[T]) Clear() {
	pq.binaryHeap.Clear()
}

func (pq *priorityQueue[T]) String() string {
	return pq.binaryHeap.String()
}

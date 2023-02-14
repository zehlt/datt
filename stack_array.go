package datt

type StackArray[T any] struct {
	head int
	arr  []T
}

func NewStackArray[T any](cap int) *StackArray[T] {
	return &StackArray[T]{
		head: -1,
		arr:  make([]T, cap),
	}
}

// O(1)
func (s *StackArray[T]) Len() int {
	if s.empty() {
		return 0
	}

	return s.head + 1
}

// O(1)
func (s *StackArray[T]) Push(v T) {
	s.head++
	s.ensureCapacity()

	s.arr[s.head] = v
}

func (s *StackArray[T]) ensureCapacity() {
	arrlen := len(s.arr)
	if s.head == arrlen {
		var d T
		s.arr = append(s.arr, d)
	}
}

func (s *StackArray[T]) empty() bool {
	return s.head == -1
}

// O(1)
func (s *StackArray[T]) Peek() (T, bool) {
	var v T
	if s.empty() {
		return v, false
	}

	return s.arr[s.head], true
}

// O(1)
func (s *StackArray[T]) Pop() (T, bool) {
	var v T
	if s.empty() {
		return v, false
	}

	v = s.arr[s.head]
	s.head--
	return v, true
}

// O(N)
func (s *StackArray[T]) Do(f func(v T)) {
	for i := s.head; i >= 0; i-- {
		f(s.arr[i])
	}
}

// O(1)
func (s *StackArray[T]) Clear() {
	s.head = -1
}

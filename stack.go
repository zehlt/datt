package datt

type stack[T comparable] struct {
	l linkedList[T]
}

func NewStack[T comparable]() stack[T] {
	return stack[T]{
		l: NewLinkedList[T](),
	}
}

func (s *stack[T]) Push(data T) {
	s.l.PushHead(data)
}

func (s *stack[T]) Pop() (T, error) {
	return s.l.PopHead()
}

func (s *stack[T]) PeekFront() (T, error) {
	return s.l.PeekHead()
}

func (s *stack[T]) PeekBack() (T, error) {
	return s.l.PeekTail()
}

func (s *stack[T]) Length() int {
	return s.l.Length()
}

func (s *stack[T]) IsEmpty() bool {
	return s.l.IsEmpty()
}

func (s *stack[T]) Clear() {
	s.l.Clear()
}

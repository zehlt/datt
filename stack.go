package datt

type Stack[T any] struct {
	l LinkedList[T]
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		l: NewLinkedList[T](),
	}
}

func (s *Stack[T]) Push(data T) {
	s.l.PushHead(data)
}

func (s *Stack[T]) Pop() (T, error) {
	return s.l.PopHead()
}

func (s *Stack[T]) PeekFront() (T, error) {
	return s.l.PeekHead()
}

func (s *Stack[T]) PeekBack() (T, error) {
	return s.l.PeekTail()
}

func (s *Stack[T]) Length() int {
	return s.l.Length()
}

func (s *Stack[T]) IsEmpty() bool {
	return s.l.IsEmpty()
}

func (s *Stack[T]) Clear() {
	s.l.Clear()
}

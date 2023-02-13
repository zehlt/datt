package datt_test

import (
	"testing"

	"github.com/zehlt/datt"
)

func TestStackLinkedListLen(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		s := datt.StackLinkedList[string]{}
		got := s.Len()

		AssertEqual(t, got, 0)
	})

	t.Run("stack with 2 element", func(t *testing.T) {
		s := datt.StackLinkedList[string]{}
		s.Push("1")
		s.Push("2")
		s.Push("7s")

		got := s.Len()
		AssertEqual(t, got, 3)
	})
}

func TestStackLinkedListPeek(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		s := datt.StackLinkedList[string]{}
		_, ok := s.Peek()

		AssertEqual(t, ok, false)
	})

	t.Run("stack with 3 element must peek the last pushed", func(t *testing.T) {
		s := datt.StackLinkedList[string]{}
		s.Push("sa")
		s.Push("henry")
		s.Push("doe")

		v, ok := s.Peek()
		AssertEqual(t, ok, true)
		AssertEqual(t, v, "doe")
	})
}

func TestStackLinkedListPush(t *testing.T) {
	t.Run("push pop push pop", func(t *testing.T) {
		s := datt.StackLinkedList[string]{}
		s.Push("sa")
		s.Pop()
		s.Push("bob")
		v, ok := s.Pop()

		AssertEqual(t, ok, true)
		AssertEqual(t, v, "bob")
	})
}

func TestStackLinkedListPop(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		s := datt.StackLinkedList[string]{}
		_, ok := s.Pop()

		AssertEqual(t, ok, false)
	})

	t.Run("stack with 3 element must pop the last pushed", func(t *testing.T) {
		s := datt.StackLinkedList[string]{}
		s.Push("sa")
		s.Push("henry")
		s.Push("doe")

		v, ok := s.Pop()
		AssertEqual(t, ok, true)
		AssertEqual(t, v, "doe")
		AssertEqual(t, s.Len(), 2)
	})
}

func TestStackLinkedListClear(t *testing.T) {
	t.Run("stack with 3 element must empty", func(t *testing.T) {
		s := datt.StackLinkedList[string]{}
		s.Push("henry")
		s.Push("doe")
		s.Push("sa")

		s.Clear()
		_, got := s.Peek()
		AssertEqual(t, s.Len(), 0)
		AssertEqual(t, got, false)
	})

	t.Run("add one element after clear", func(t *testing.T) {
		s := datt.StackLinkedList[string]{}
		s.Push("henry")
		s.Push("doe")
		s.Push("sa")

		s.Clear()
		s.Push("alice")
		got, ok := s.Peek()

		AssertEqual(t, s.Len(), 1)
		AssertEqual(t, got, "alice")
		AssertEqual(t, ok, true)
	})
}

func TestStackLinkedListDo(t *testing.T) {
	t.Run("empty stack should not iterate", func(t *testing.T) {
		s := datt.StackLinkedList[int]{}

		var iter int
		s.Do(func(v int) {
			iter++
		})

		AssertEqual(t, iter, 0)
	})

	t.Run("stack with elements should iterate on all", func(t *testing.T) {
		s := datt.StackLinkedList[int]{}
		var data [4]int = [4]int{1, 2, 3, 4}

		for _, v := range data {
			s.Push(v)
		}

		var got [4]int
		var i int = 3
		s.Do(func(v int) {
			got[i] = v
			i--
		})

		AssertEqual(t, got, data)
	})
}

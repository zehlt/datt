package datt_test

import (
	"testing"

	"github.com/zehlt/datt"
)

func TestQueueLinkedListLen(t *testing.T) {
	t.Run("empty queue should return 0", func(t *testing.T) {
		q := datt.QueueLinkedList[int]{}
		got := q.Len()

		AssertEqual(t, got, 0)
	})

	t.Run("after enqueue len should increase by 1", func(t *testing.T) {
		q := datt.QueueLinkedList[int]{}
		q.Enqueue(7)
		q.Enqueue(7)
		got := q.Len()

		AssertEqual(t, got, 2)
	})
}

func TestQueueLinkedListPeek(t *testing.T) {
	t.Run("front on empty queue should return false", func(t *testing.T) {
		q := datt.QueueLinkedList[int]{}
		_, ok := q.Peek()

		AssertEqual(t, ok, false)
	})

	t.Run("with 1 value should return true and the value ", func(t *testing.T) {
		q := datt.QueueLinkedList[int]{}
		q.Enqueue(114)
		got, ok := q.Peek()

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 114)
	})

	t.Run("with 2 value should return true and the first value added ", func(t *testing.T) {
		q := datt.QueueLinkedList[int]{}
		q.Enqueue(114)
		q.Enqueue(338)
		got, ok := q.Peek()

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 114)
	})
}

func TestQueueLinkedListEnqueue(t *testing.T) {
	t.Run("multiple enqueue", func(t *testing.T) {
		q := datt.QueueLinkedList[float32]{}
		q.Enqueue(16.3)
		q.Enqueue(1.398)
		q.Enqueue(10.398)

		AssertEqual(t, q.Len(), 3)
	})

	t.Run("enqueue dequeue enqueue dequeue", func(t *testing.T) {
		q := datt.QueueLinkedList[float32]{}
		q.Enqueue(16.3)
		q.Enqueue(1.398)
		q.Dequeue()
		q.Enqueue(10.398)
		got, ok := q.Dequeue()

		AssertEqual(t, q.Len(), 1)
		AssertEqual(t, ok, true)
		AssertEqual(t, got, 1.398)
	})
}

func TestQueueLinkedListDequeue(t *testing.T) {
	t.Run("empty queue should not return any value", func(t *testing.T) {
		q := datt.QueueLinkedList[int]{}
		_, ok := q.Dequeue()

		AssertEqual(t, ok, false)
		AssertEqual(t, q.Len(), 0)
	})

	t.Run("with 1 value should dequeue the only value", func(t *testing.T) {
		q := datt.QueueLinkedList[rune]{}
		q.Enqueue('c')
		got, ok := q.Dequeue()

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 'c')
		AssertEqual(t, q.Len(), 0)
	})

	t.Run("with 2 value should dequeue the fist value added", func(t *testing.T) {
		q := datt.QueueLinkedList[rune]{}
		q.Enqueue('c')
		q.Enqueue('ひ')
		got, ok := q.Dequeue()

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 'c')
		AssertEqual(t, q.Len(), 1)
	})

	t.Run("multiple dequeue", func(t *testing.T) {
		q := datt.QueueLinkedList[rune]{}
		q.Enqueue('c')
		q.Enqueue('ひ')
		q.Dequeue()
		got, ok := q.Dequeue()

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 'ひ')
		AssertEqual(t, q.Len(), 0)
	})
}

func TestQueueLinkedListClear(t *testing.T) {
	t.Run("clear empty queue should not change len", func(t *testing.T) {
		q := datt.QueueLinkedList[int]{}
		q.Clear()

		AssertEqual(t, q.Len(), 0)
	})

	t.Run("clear queue with one element", func(t *testing.T) {
		q := datt.QueueLinkedList[string]{}
		q.Enqueue("henry")
		q.Clear()

		AssertEqual(t, q.Len(), 0)
	})
}

func TestQueueLinkedListDo(t *testing.T) {
	t.Run("empty queue should not iterate", func(t *testing.T) {
		q := datt.QueueLinkedList[int]{}

		var iter int
		q.Do(func(v int) {
			iter++
		})

		AssertEqual(t, iter, 0)
	})

	t.Run("queue with elements should iterate on all", func(t *testing.T) {
		q := datt.QueueLinkedList[int]{}
		var data [4]int = [4]int{1, 2, 3, 4}

		for _, v := range data {
			q.Enqueue(v)
		}

		var got [4]int
		var i int
		q.Do(func(v int) {
			got[i] = v
			i++
		})

		AssertEqual(t, got, data)
	})
}

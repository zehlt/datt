package datt_test

import (
	"testing"

	"github.com/zehlt/datt"
)

func TestQueueCircularArrayLen(t *testing.T) {
	t.Run("empty with cap of 5", func(t *testing.T) {
		q := datt.NewQueueCircularArray[float32](5)
		got := q.Len()

		AssertEqual(t, got, 0)
	})

	t.Run("enqueue 1", func(t *testing.T) {
		q := datt.NewQueueCircularArray[float32](5)
		q.Enqueue(26.2)
		q.Enqueue(19.236)
		got := q.Len()

		AssertEqual(t, got, 2)
	})
}

func TestQueueCircularArrayPeek(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		q := datt.NewQueueCircularArray[float32](5)
		_, ok := q.Peek()

		AssertEqual(t, ok, false)
	})

	t.Run("enqueue 1", func(t *testing.T) {
		q := datt.NewQueueCircularArray[float32](5)
		q.Enqueue(26.2)
		got, ok := q.Peek()

		AssertEqual(t, got, 26.2)
		AssertEqual(t, ok, true)
	})

	t.Run("enqueue 3", func(t *testing.T) {
		q := datt.NewQueueCircularArray[float32](5)
		q.Enqueue(84.652)
		q.Enqueue(26.2)
		q.Enqueue(8.652)
		got, ok := q.Peek()

		AssertEqual(t, got, 84.652)
		AssertEqual(t, ok, true)
	})

	t.Run("enqueue 3 and dequeue 3", func(t *testing.T) {
		q := datt.NewQueueCircularArray[float32](3)
		q.Enqueue(84.652)
		q.Enqueue(26.2)
		q.Enqueue(8.652)

		q.Dequeue()
		q.Dequeue()
		q.Dequeue()

		q.Enqueue(6.0)
		got, ok := q.Peek()

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 6.0)
	})

	t.Run("force the cycle of the circular array", func(t *testing.T) {
		q := datt.NewQueueCircularArray[string](5)
		q.Enqueue("a")
		q.Dequeue()

		q.Enqueue("b")
		q.Enqueue("c")
		q.Enqueue("d")
		q.Enqueue("e")
		q.Enqueue("f")

		got, ok := q.Peek()
		AssertEqual(t, got, "b")
		AssertEqual(t, ok, true)
		AssertEqual(t, q.Len(), 5)
	})

	t.Run("queue should grow the capacity", func(t *testing.T) {
		q := datt.NewQueueCircularArray[string](2)
		q.Enqueue("a")
		q.Enqueue("b")
		q.Enqueue("c")
		q.Enqueue("d")

		got, ok := q.Peek()
		AssertEqual(t, got, "a")
		AssertEqual(t, ok, true)
		AssertEqual(t, q.Len(), 4)
	})

	t.Run("dequeue to the farest", func(t *testing.T) {
		q := datt.NewQueueCircularArray[string](3)
		q.Enqueue("a")
		q.Enqueue("b")
		q.Enqueue("c")
		q.Dequeue()
		q.Dequeue()
		q.Dequeue()

		_, ok := q.Dequeue()
		AssertEqual(t, ok, false)
		AssertEqual(t, q.Len(), 0)
	})

	t.Run("dequeue to the farest++", func(t *testing.T) {
		q := datt.NewQueueCircularArray[string](3)
		q.Enqueue("a")
		q.Enqueue("b")
		q.Enqueue("c")
		q.Dequeue()
		q.Dequeue()
		q.Dequeue()

		q.Enqueue("d")
		got, ok := q.Dequeue()

		AssertEqual(t, got, "d")
		AssertEqual(t, ok, true)
		AssertEqual(t, q.Len(), 0)
	})

	t.Run("force grow: 1 1 [d b c] -> 0 4 [b c d e  ]", func(t *testing.T) {
		q := datt.NewQueueCircularArray[string](3)
		q.Enqueue("a")
		q.Enqueue("b")
		q.Enqueue("c")
		q.Dequeue()

		q.Enqueue("d")
		q.Enqueue("e")
		got, ok := q.Peek()

		AssertEqual(t, got, "b")
		AssertEqual(t, ok, true)
		AssertEqual(t, q.Len(), 4)
	})

}

func TestQueueCircularArrayDequeue(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		q := datt.NewQueueCircularArray[float32](5)
		_, ok := q.Dequeue()

		AssertEqual(t, ok, false)
	})

	t.Run("1 element", func(t *testing.T) {
		q := datt.NewQueueCircularArray[string](5)
		q.Enqueue("alice")
		got, ok := q.Dequeue()

		AssertEqual(t, ok, true)
		AssertEqual(t, got, "alice")
	})

	t.Run("3 element", func(t *testing.T) {
		q := datt.NewQueueCircularArray[string](5)
		q.Enqueue("john")
		q.Enqueue("alice")
		q.Enqueue("bob")
		got, ok := q.Dequeue()

		AssertEqual(t, ok, true)
		AssertEqual(t, got, "john")
		AssertEqual(t, q.Len(), 2)
	})

	t.Run("dequeue multiple time", func(t *testing.T) {
		q := datt.NewQueueCircularArray[string](5)
		q.Enqueue("john")
		q.Enqueue("alice")
		q.Enqueue("bob")
		q.Dequeue()
		q.Dequeue()
		got, ok := q.Dequeue()

		AssertEqual(t, ok, true)
		AssertEqual(t, got, "bob")
		AssertEqual(t, q.Len(), 0)
	})
}

func TestQueueCircularArray(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		q := datt.NewQueueCircularArray[string](5)
		q.Clear()
		_, ok := q.Dequeue()

		AssertEqual(t, q.Len(), 0)
		AssertEqual(t, ok, false)
	})

	t.Run("clear multiple elements should empty the queue", func(t *testing.T) {
		q := datt.NewQueueCircularArray[int](4)
		q.Enqueue(4)
		q.Enqueue(3)
		q.Enqueue(1)
		q.Enqueue(7)
		q.Enqueue(9)

		q.Clear()
		_, ok := q.Dequeue()

		AssertEqual(t, q.Len(), 0)
		AssertEqual(t, ok, false)
	})

	t.Run("enqueue should work normally after clear", func(t *testing.T) {
		q := datt.NewQueueCircularArray[int](4)
		q.Enqueue(4)
		q.Enqueue(9)

		q.Clear()

		q.Enqueue(11)
		q.Enqueue(7)
		q.Enqueue(2)
		got, ok := q.Peek()

		AssertEqual(t, q.Len(), 3)
		AssertEqual(t, ok, true)
		AssertEqual(t, got, 11)
	})
}

func TestQueueCircularArrayDo(t *testing.T) {
	t.Run("empty queue should not iterate", func(t *testing.T) {
		q := datt.NewQueueCircularArray[int](4)

		var iter int
		q.Do(func(v int) {
			iter++
		})

		AssertEqual(t, iter, 0)
	})

	t.Run("queue with elements should iterate on all", func(t *testing.T) {
		q := datt.NewQueueCircularArray[int](4)
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

	t.Run("queue with circle should iterate on all", func(t *testing.T) {
		q := datt.NewQueueCircularArray[int](3)
		var want [3]int = [3]int{2, 3, 4}

		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		q.Dequeue()
		q.Enqueue(4)

		var got [3]int
		var i int
		q.Do(func(v int) {
			got[i] = v
			i++
		})

		AssertEqual(t, got, want)
	})
}

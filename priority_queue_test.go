package datt_test

import (
	"testing"

	"github.com/zehlt/datt"
)

func TestPriorityQueueLen(t *testing.T) {
	t.Run("empty queue", func(t *testing.T) {
		q := datt.NewPriorityQueue(datt.CompareOrdered[int])
		got := q.Len()
		AssertEqual(t, got, 0)
	})

	t.Run("3 element", func(t *testing.T) {
		q := datt.NewPriorityQueue(datt.CompareOrdered[int])
		q.Enqueue(10)
		q.Enqueue(-5)
		q.Enqueue(499)

		q.Enqueue(499)
		q.Enqueue(499)
		q.Enqueue(499)

		got := q.Len()
		AssertEqual(t, got, 6)
	})
}

func TestPriorityQueuePeek(t *testing.T) {
	t.Run("empty queue", func(t *testing.T) {
		q := datt.NewPriorityQueue(datt.CompareOrdered[string])

		_, ok := q.Peek()
		AssertEqual(t, ok, false)
		AssertEqual(t, q.Len(), 0)
	})

	t.Run("should return the only element without changing the len", func(t *testing.T) {
		q := datt.NewPriorityQueue(datt.CompareOrdered[string])
		q.Enqueue("alice")
		got, ok := q.Peek()

		AssertEqual(t, got, "alice")
		AssertEqual(t, ok, true)
		AssertEqual(t, q.Len(), 1)
	})

	t.Run("2 element", func(t *testing.T) {
		q := datt.NewPriorityQueue(datt.CompareOrdered[int])
		q.Enqueue(1)
		q.Enqueue(2)
		got, ok := q.Peek()

		AssertEqual(t, got, 1)
		AssertEqual(t, ok, true)
		AssertEqual(t, q.Len(), 2)
	})

	t.Run("4 element", func(t *testing.T) {
		q := datt.NewPriorityQueue(datt.CompareOrdered[int])
		q.Enqueue(2)
		q.Enqueue(3)
		q.Enqueue(1)
		q.Enqueue(4)
		got, ok := q.Peek()

		AssertEqual(t, got, 1)
		AssertEqual(t, ok, true)
		AssertEqual(t, q.Len(), 4)
	})
}

func TestPriorityQueueDequeue(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		q := datt.NewPriorityQueue(datt.CompareOrdered[int])
		_, ok := q.Dequeue()

		AssertEqual(t, ok, false)
		AssertEqual(t, q.Len(), 0)
	})

	t.Run("empty", func(t *testing.T) {
		q := datt.NewPriorityQueue(datt.CompareOrdered[int])
		q.Enqueue(55)
		got, ok := q.Dequeue()

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 55)
		AssertEqual(t, q.Len(), 0)
	})

	t.Run("tc1", func(t *testing.T) {
		q := datt.NewPriorityQueue(datt.CompareOrdered[int])
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		q.Enqueue(4)
		q.Enqueue(5)
		q.Enqueue(6)
		q.Enqueue(7)
		q.Enqueue(8)
		got, ok := q.Dequeue()

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 1)
		AssertEqual(t, q.Len(), 7)
	})
}

func TestPriorityQueueClear(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		q := datt.NewPriorityQueue(datt.CompareOrdered[int])
		q.Clear()

		AssertEqual(t, q.Len(), 0)
	})

	t.Run("remove all enqueues", func(t *testing.T) {
		q := datt.NewPriorityQueue(datt.CompareOrdered[string])
		q.Enqueue("bob")
		q.Enqueue("alice")
		q.Enqueue("john")
		q.Enqueue("doe")
		q.Clear()

		q.Enqueue("harry")
		got, ok := q.Dequeue()

		AssertEqual(t, q.Len(), 0)
		AssertEqual(t, ok, true)
		AssertEqual(t, got, "harry")

	})
}

func TestPriorityQueueDo(t *testing.T) {
	testCases := []struct {
		desc     string
		enqueues []int
		want     []int
	}{
		{
			desc:     "level order",
			enqueues: []int{1, 2, 3, 4, 5, 6},
			want:     []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			q := datt.NewPriorityQueue(datt.CompareOrdered[int])
			for _, v := range tC.enqueues {
				q.Enqueue(v)
			}

			got := []int{}
			q.Do(func(v int) {
				got = append(got, v)
			})

			AssertEqual(t, got, tC.want)
		})
	}
}

func TestPriorityQueueTC(t *testing.T) {
	testCases := []struct {
		desc     string
		enqueues []int
		want     []int
	}{
		{
			desc:     "c1",
			enqueues: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want:     []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			desc:     "empty",
			enqueues: []int{},
			want:     []int{},
		},
		{
			desc:     "one element",
			enqueues: []int{5},
			want:     []int{5},
		},
		{
			desc:     "one element",
			enqueues: []int{5, 4, 3, 2, 1, 0, -1},
			want:     []int{-1, 0, 1, 2, 3, 4, 5},
		},
		{
			desc:     "one element same",
			enqueues: []int{5, 4, 4, 2, 1, 0, -1},
			want:     []int{-1, 0, 1, 2, 4, 4, 5},
		},
		{
			desc:     "same element",
			enqueues: []int{3, 5, 4, 6, 7, 5},
			want:     []int{3, 4, 5, 5, 6, 7},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			q := datt.NewPriorityQueue(datt.CompareOrdered[int])

			for _, enq := range tC.enqueues {
				q.Enqueue(enq)
			}

			got := []int{}
			for q.Len() > 0 {
				v, _ := q.Dequeue()
				got = append(got, v)
			}

			AssertEqual(t, got, tC.want)
		})
	}
}

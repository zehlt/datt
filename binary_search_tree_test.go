package datt_test

import (
	"testing"

	"github.com/zehlt/datt"
)

func TestBinarySearchTreeCompareFunc(t *testing.T) {
	t.Run("setting a wrong compare func must panic", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(func(a int, b int) datt.CompareResult {
			return -5
		})

		bst.Insert(3)

		AssertPanic(t, func() {
			bst.Insert(7)
		})

		AssertPanic(t, func() {
			bst.Has(3)
		})

		AssertPanic(t, func() {
			bst.Remove(3)
		})
	})

}

func TestBinarySearchTreeLen(t *testing.T) {
	t.Run("empty bst", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[string])
		got := bst.Len()

		AssertEqual(t, got, 0)
	})

	t.Run("after insert len should increase", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[string])
		bst.Insert("a")
		bst.Insert("b")
		bst.Insert("c")
		got := bst.Len()

		AssertEqual(t, got, 3)
	})
}

func TestBinarySearchTreeHas(t *testing.T) {
	t.Run("empty bst", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[string])
		got := bst.Has("a")

		AssertEqual(t, got, false)
	})

	t.Run("contain only the root element", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[string])
		bst.Insert("a")
		got := bst.Has("a")

		AssertEqual(t, got, true)
	})

	t.Run("insert second element higher", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[int])
		bst.Insert(5)
		bst.Insert(10)
		got := bst.Has(10)

		AssertEqual(t, got, true)
	})

	t.Run("insert two equal element second is ignored", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[int])
		bst.Insert(5)
		bst.Insert(4)
		bst.Insert(3)
		bst.Insert(4)

		AssertEqual(t, bst.Len(), 3)
	})

	t.Run("branch right then left", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[float32])
		bst.Insert(5)
		bst.Insert(10)
		bst.Insert(8)
		bst.Insert(6)
		got := bst.Has(6)

		AssertEqual(t, got, true)
	})

	t.Run("try to find a value not inserted", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[float32])
		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(6)
		bst.Insert(8)
		got := bst.Has(1)

		AssertEqual(t, got, false)
	})
}

func TestBinarySearchTreeRemove(t *testing.T) {
	t.Run("try to remove on empty bst", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[rune])
		got := bst.Remove('a')

		AssertEqual(t, got, false)
		AssertEqual(t, bst.Len(), 0)
	})

	t.Run("bst does not contain the element", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[rune])
		bst.Insert('a')
		bst.Insert('d')
		bst.Insert('f')
		got := bst.Remove('i')

		AssertEqual(t, got, false)
		AssertEqual(t, bst.Len(), 3)
	})

	t.Run("try to remove the root", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[rune])
		bst.Insert('d')
		got := bst.Remove('d')

		AssertEqual(t, got, true)
		AssertEqual(t, bst.Has('d'), false)
		AssertEqual(t, bst.Len(), 0)
	})

	t.Run("remove only the left leaf", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[int])
		bst.Insert(5)
		bst.Insert(4)
		got := bst.Remove(4)

		AssertEqual(t, got, true)
		AssertEqual(t, bst.Has(5), true)
		AssertEqual(t, bst.Has(4), false)
		AssertEqual(t, bst.Len(), 1)
	})

	t.Run("remove only the right leaf", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[int])
		bst.Insert(5)
		bst.Insert(6)
		got := bst.Remove(6)

		AssertEqual(t, got, true)
		AssertEqual(t, bst.Has(5), true)
		AssertEqual(t, bst.Has(6), false)
		AssertEqual(t, bst.Len(), 1)
	})

	t.Run("delete root with two child should take left leaf", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[int])
		bst.Insert(5)
		bst.Insert(6)
		bst.Insert(4)
		got := bst.Remove(5)

		AssertEqual(t, got, true)
		AssertEqual(t, bst.Has(4), true)
		AssertEqual(t, bst.Has(5), false)
		AssertEqual(t, bst.Has(6), true)
		AssertEqual(t, bst.Len(), 2)
	})

	t.Run("remove intermediate value", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[int])
		bst.Insert(5)
		bst.Insert(4)
		bst.Insert(3)
		got := bst.Remove(4)

		AssertEqual(t, got, true)
		AssertEqual(t, bst.Has(5), true)
		AssertEqual(t, bst.Has(4), false)
		AssertEqual(t, bst.Has(3), true)
		AssertEqual(t, bst.Len(), 2)
	})

	t.Run("digLeft", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[int])
		bst.Insert(9)
		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(2)
		bst.Insert(7)
		bst.Insert(3)

		got := bst.Remove(9)
		AssertEqual(t, got, true)
		AssertEqual(t, bst.Has(9), false)
		AssertEqual(t, bst.Len(), 5)
	})

	t.Run("remove intermediate value", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[int])
		bst.Insert(5)
		bst.Insert(4)
		bst.Insert(3)
		got := bst.Remove(4)

		AssertEqual(t, got, true)
		AssertEqual(t, bst.Has(5), true)
		AssertEqual(t, bst.Has(4), false)
		AssertEqual(t, bst.Has(3), true)
		AssertEqual(t, bst.Len(), 2)
	})

	t.Run("remove with dig left", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[int])
		bst.Insert(10)
		bst.Insert(7)
		bst.Insert(12)
		bst.Insert(11)
		bst.Insert(6)
		bst.Insert(9)
		bst.Insert(8)
		bst.Insert(5)
		got := bst.Remove(7)

		AssertEqual(t, got, true)
		AssertEqual(t, bst.Len(), 7)
	})
}

func TestBinarySearchTreeClear(t *testing.T) {
	t.Run("clear on empty bst does nothing", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[rune])
		bst.Clear()
		got := bst.Len()

		AssertEqual(t, got, 0)
	})

	t.Run("after clear we should be able use it again", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[rune])
		bst.Insert('雪')
		bst.Insert('é')
		bst.Clear()
		bst.Insert('9')

		AssertEqual(t, bst.Has('9'), true)
		AssertEqual(t, bst.Has('雪'), false)
	})
}

func TestBinarySearchTreeDo(t *testing.T) {
	t.Run("use wrong order traversal constant", func(t *testing.T) {
		bst := datt.NewBinarySearchTree(datt.CompareOrdered[int])

		AssertPanic(t, func() {
			bst.Do(-1, func(v int) {})
		})
	})

	testCases := []struct {
		desc    string
		order   datt.OrderTraversal
		inserts []int
		want    []int
	}{
		{
			desc:    "pre order",
			order:   datt.PreOrder,
			inserts: []int{6, 7, 9, 4, 2, 3},
			want:    []int{6, 4, 2, 3, 7, 9},
		},
		{
			desc:    "pre order",
			order:   datt.PreOrder,
			inserts: []int{10, 12, 11, 8, 9, 7, 6},
			want:    []int{10, 8, 7, 6, 9, 12, 11},
		},
		{
			desc:    "in order",
			order:   datt.InOrder,
			inserts: []int{6, 7, 9, 4, 2, 3},
			want:    []int{2, 3, 4, 6, 7, 9},
		},
		{
			desc:    "in",
			order:   datt.InOrder,
			inserts: []int{10, 12, 11, 8, 9, 7, 6},
			want:    []int{6, 7, 8, 9, 10, 11, 12},
		},
		{
			desc:    "post order",
			order:   datt.PostOrder,
			inserts: []int{6, 7, 9, 4, 2, 3},
			want:    []int{3, 2, 4, 9, 7, 6},
		},
		{
			desc:    "post order",
			order:   datt.PostOrder,
			inserts: []int{10, 12, 11, 8, 9, 7, 6},
			want:    []int{6, 7, 9, 8, 11, 12, 10},
		},
		{
			desc:    "level order 1",
			order:   datt.LevelOrder,
			inserts: []int{6, 7, 9, 4, 2, 3},
			want:    []int{6, 4, 7, 2, 9, 3},
		},
		{
			desc:    "level order 2",
			order:   datt.LevelOrder,
			inserts: []int{10, 12, 11, 8, 9, 7, 6},
			want:    []int{10, 8, 12, 7, 9, 11, 6},
		},
		{
			desc:    "level order without insertion",
			order:   datt.LevelOrder,
			inserts: []int{},
			want:    []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			bst := datt.NewBinarySearchTree(datt.CompareOrdered[int])
			for _, v := range tC.inserts {
				bst.Insert(v)
			}

			got := []int{}
			bst.Do(tC.order, func(v int) {
				got = append(got, v)
			})

			AssertEqual(t, got, tC.want)
		})
	}
}

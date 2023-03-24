package datt_test

import (
	"testing"

	"github.com/zehlt/datt"
)

func TestRedBlackTreeLen(t *testing.T) {
	t.Run("len on empty tree return 0", func(t *testing.T) {
		rb := datt.NewRedBlackTree[int, string](datt.CompareOrdered[int])

		got := rb.Len()

		AssertEqual(t, got, 0)
	})

	t.Run("len on multiple insert must increase", func(t *testing.T) {
		rb := datt.NewRedBlackTree[int, string](datt.CompareOrdered[int])
		rb.Insert(125, "pierre")
		rb.Insert(51, "henry")
		rb.Insert(12, "john")
		rb.Insert(-10, "alice")

		got := rb.Len()

		AssertEqual(t, got, 4)
	})
}

func TestRedBlackTreeHas(t *testing.T) {
	t.Run("does not contain", func(t *testing.T) {
		rb := datt.NewRedBlackTree[int, string](datt.CompareOrdered[int])

		got := rb.Has(10)

		AssertEqual(t, got, false)
	})

	t.Run("tree with one element", func(t *testing.T) {
		rb := datt.NewRedBlackTree[int, string](datt.CompareOrdered[int])
		rb.Insert(10, "alice")

		got := rb.Has(10)

		AssertEqual(t, got, true)
	})

	t.Run("insert with bigger key", func(t *testing.T) {
		rb := datt.NewRedBlackTree[int, string](datt.CompareOrdered[int])
		rb.Insert(10, "alice")
		rb.Insert(20, "bob")

		got := rb.Has(20)

		AssertEqual(t, got, true)
	})
}

func TestRedBlackTreeGet(t *testing.T) {
	t.Run("empty redblacktree", func(t *testing.T) {
		rb := datt.NewRedBlackTree[string, int](datt.CompareOrdered[string])

		_, ok := rb.Get("alice")

		AssertEqual(t, ok, false)
	})

	t.Run("one element rbt", func(t *testing.T) {
		rb := datt.NewRedBlackTree[string, int](datt.CompareOrdered[string])

		rb.Insert("alice", 25)
		got, ok := rb.Get("alice")

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 25)
	})

	t.Run("two elements with second lower", func(t *testing.T) {
		rb := datt.NewRedBlackTree[string, int](datt.CompareOrdered[string])

		rb.Insert("bob", 33)
		rb.Insert("alice", 25)
		got, ok := rb.Get("alice")

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 25)
	})

	t.Run("three elements left", func(t *testing.T) {
		rb := datt.NewRedBlackTree[string, int](datt.CompareOrdered[string])

		rb.Insert("ken", 67)
		rb.Insert("bob", 33)
		rb.Insert("alice", 25)
		got, ok := rb.Get("alice")

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 25)
	})

	t.Run("three elements right", func(t *testing.T) {
		rb := datt.NewRedBlackTree[string, int](datt.CompareOrdered[string])

		rb.Insert("alice", 25)
		rb.Insert("bob", 33)
		rb.Insert("ken", 67)
		got, ok := rb.Get("ken")

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 67)
	})

	t.Run("three elements cross", func(t *testing.T) {
		rb := datt.NewRedBlackTree[string, int](datt.CompareOrdered[string])

		rb.Insert("alice", 25)
		rb.Insert("ken", 67)
		rb.Insert("bob", 33)
		got, ok := rb.Get("bob")

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 33)
	})

	t.Run("override root", func(t *testing.T) {
		rb := datt.NewRedBlackTree[string, int](datt.CompareOrdered[string])

		rb.Insert("bob", 33)
		rb.Insert("bob", 28)
		got, ok := rb.Get("bob")

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 28)
	})

	t.Run("override any", func(t *testing.T) {
		rb := datt.NewRedBlackTree[string, int](datt.CompareOrdered[string])

		rb.Insert("bob", 33)
		rb.Insert("alice", 26)
		rb.Insert("ken", 68)
		rb.Insert("ken", 69)

		got, ok := rb.Get("ken")

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 69)
	})
}

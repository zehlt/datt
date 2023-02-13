package datt_test

import (
	"reflect"
	"testing"
)

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/zehlt/datt"
// )

// func TestLinkedListNew(t *testing.T) {
// 	t.Run("new", func(t *testing.T) {
// 		s := datt.NewLinkedList[int]()

// 		if s == nil {
// 			t.Error("returned an nil ptr")
// 		}
// 	})
// }

// func TestLinkedListLen(t *testing.T) {
// 	t.Run("ll Len should be 0", func(t *testing.T) {
// 		s := datt.NewLinkedList[string]()
// 		got := s.Len()

// 		AssertEqual(t, got, 0)
// 	})

// 	t.Run("ll Len should be 1 after push", func(t *testing.T) {
// 		s := datt.NewLinkedList[string]()
// 		s.PushFront("henry")

// 		got := s.Len()
// 		AssertEqual(t, got, 1)
// 	})

// 	t.Run("ll Len should be 2 after 2 push", func(t *testing.T) {
// 		s := datt.NewLinkedList[string]()
// 		s.PushFront("henry")
// 		s.PushFront("marie")

// 		got := s.Len()
// 		AssertEqual(t, got, 2)
// 	})
// }

// func TestLinkedListPushFront(t *testing.T) {
// 	t.Run("must return a node with the passing value", func(t *testing.T) {
// 		ll := datt.NewLinkedList[float64]()
// 		got := ll.PushFront(123.0)

// 		AssertEqual(t, got.Value, 123.0)
// 		AssertEqual(t, ll.Len(), 1)
// 	})
// }

// func TestLinkedListFront(t *testing.T) {
// 	t.Run("when ll is empty front must return nil", func(t *testing.T) {
// 		ll := datt.NewLinkedList[rune]()
// 		v := ll.Front()

// 		AssertEqual(t, v, nil)
// 	})

// 	t.Run("must return the head node", func(t *testing.T) {
// 		ll := datt.NewLinkedList[rune]()
// 		got := ll.PushFront('赤')
// 		want := ll.Front()

// 		AssertEqual(t, got, want)
// 	})

// 	t.Run("must return the last pushed head node", func(t *testing.T) {
// 		ll := datt.NewLinkedList[rune]()
// 		ll.PushFront('雪')

// 		got := ll.PushFront('é')
// 		want := ll.Front()

// 		AssertEqual(t, got, want)
// 	})
// }

// func TestLinkedListBack(t *testing.T) {
// 	t.Run("when ll is empty back must return nil", func(t *testing.T) {
// 		ll := datt.NewLinkedList[rune]()
// 		got := ll.Back()
// 		AssertEqual(t, got, nil)
// 	})

// 	t.Run("when ll contains 1 element must return it", func(t *testing.T) {
// 		ll := datt.NewLinkedList[byte]()
// 		want := ll.PushFront(45)
// 		got := ll.Back()

// 		AssertEqual(t, got, want)
// 	})

// 	t.Run("when ll contains 2 element must return the fist pushed", func(t *testing.T) {
// 		ll := datt.NewLinkedList[byte]()
// 		want := ll.PushFront(125)
// 		ll.PushFront(45)
// 		got := ll.Back()

// 		AssertEqual(t, got, want)
// 	})
// }

// func TestLinkedListClear(t *testing.T) {
// 	t.Run("must remove the previous nodes", func(t *testing.T) {
// 		ll := datt.NewLinkedList[rune]()
// 		ll.PushFront('1')
// 		ll.PushFront('&')
// 		ll.Clear()

// 		AssertEqual(t, ll.Len(), 0)
// 	})
// }

// func TestNode(t *testing.T) {
// 	t.Run("Value property must be the given one", func(t *testing.T) {
// 		node := datt.Node[float32]{
// 			Value: 122,
// 		}

// 		AssertEqual(t, node.Value, 122)
// 	})

// 	t.Run("Next on single node not linked to a list must return nil", func(t *testing.T) {
// 		node := datt.Node[float32]{Value: 88.001}
// 		want := node.Next()

// 		AssertEqual(t, want, nil)
// 	})

// 	t.Run("Next on single node must return nil", func(t *testing.T) {
// 		ll := datt.NewLinkedList[byte]()

// 		first := ll.PushFront(1)
// 		got := first.Next()

// 		AssertEqual(t, got, nil)
// 	})

// 	t.Run("Next on linked node must return the next one", func(t *testing.T) {
// 		ll := datt.NewLinkedList[byte]()

// 		want := ll.PushFront(1)
// 		current := ll.PushFront(12)
// 		got := current.Next()

// 		AssertEqual(t, got, want)
// 	})

// 	t.Run("Next on single node must return nil", func(t *testing.T) {
// 		node := datt.Node[float32]{Value: 88.001}
// 		want := node.Prev()

// 		AssertEqual(t, want, nil)
// 	})

// 	t.Run("Prev on linked node must return the previous one", func(t *testing.T) {
// 		ll := datt.NewLinkedList[byte]()

// 		current := ll.PushFront(12)
// 		want := ll.PushFront(1)
// 		got := current.Prev()

// 		AssertEqual(t, got, want)
// 	})
// }

func AssertEqual[T comparable](t testing.TB, got T, want T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// func AssertError(t testing.TB, got error, want error) {
// 	t.Helper()

// 	if got != want {
// 		t.Errorf("got %s want %s", got, want)
// 	}
// }

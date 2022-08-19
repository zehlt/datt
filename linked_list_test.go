package datt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkQueuePush(b *testing.B) {
	ll := NewLinkedList[string]()

	for i := 0; i < b.N; i++ {
		ll.PushHead("bab")
	}
}

func BenchmarkQueuePushAndPop(b *testing.B) {
	ll := NewLinkedList[string]()

	for i := 0; i < b.N; i++ {
		ll.PushHead("bab")
		ll.PushHead("bub")
		// ll.PopTail()
	}
}

func TestPushHeadOnce(t *testing.T) {
	l := NewLinkedList[string]()
	expected := "pierre"

	l.PushHead(expected)
	assert.Equal(t, l.Length(), 1)

	got, err := l.PopHead()
	assert.NoError(t, err)

	assert.Equal(t, expected, got)
	assert.Equal(t, l.Length(), 0)
}

func TestPeakOneItem(t *testing.T) {
	l := NewLinkedList[string]()
	expected := "pierre"

	l.PushHead(expected)
	assert.Equal(t, l.Length(), 1)

	got, err := l.PeekHead()
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
	assert.Equal(t, l.Length(), 1)

	got, err = l.PeekTail()
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
	assert.Equal(t, l.Length(), 1)
}

func TestPeakEmpty(t *testing.T) {
	l := NewLinkedList[string]()
	assert.Equal(t, l.Length(), 0)

	_, err := l.PeekHead()
	assert.Equal(t, ErrEmptyList, err)

	_, err = l.PeekTail()
	assert.Equal(t, ErrEmptyList, err)
}

func TestPushTailOnce(t *testing.T) {
	l := NewLinkedList[string]()
	expected := "pierre"

	l.PushTail(expected)
	assert.Equal(t, 1, l.Length())

	got, err := l.PopTail()
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
	assert.Equal(t, 0, l.Length())
}

func TestPushTailTwice(t *testing.T) {
	l := NewLinkedList[string]()
	expected1 := "pierre"
	expected2 := "henry"

	l.PushTail(expected1)
	l.PushTail(expected2)
	assert.Equal(t, l.Length(), 2)

	got1, err := l.PopTail()
	assert.NoError(t, err)

	got2, err := l.PopTail()
	assert.NoError(t, err)

	assert.Equal(t, expected2, got1)
	assert.Equal(t, expected1, got2)
	assert.Equal(t, l.Length(), 0)
}

func TestPushHeadTwice(t *testing.T) {
	l := NewLinkedList[string]()
	expected1 := "pierre"
	expected2 := "henry"

	l.PushHead(expected1)
	l.PushHead(expected2)
	assert.Equal(t, l.Length(), 2)

	got1, err := l.PopHead()
	assert.NoError(t, err)

	got2, err := l.PopHead()
	assert.NoError(t, err)

	assert.Equal(t, expected2, got1)
	assert.Equal(t, expected1, got2)
	assert.Equal(t, l.Length(), 0)
}

func TestPushTailMultiple(t *testing.T) {
	l := NewLinkedList[string]()
	expected1 := "pierre"
	expected2 := "henry"
	expected3 := "marie"
	expected4 := "ana"

	l.PushTail(expected1)
	l.PushTail(expected2)
	l.PushTail(expected3)
	l.PushTail(expected4)
	assert.Equal(t, l.Length(), 4)

	got1, err := l.PopTail()
	assert.NoError(t, err)

	got2, err := l.PopTail()
	assert.NoError(t, err)

	got3, err := l.PopTail()
	assert.NoError(t, err)

	got4, err := l.PopTail()
	assert.NoError(t, err)

	assert.Equal(t, expected4, got1)
	assert.Equal(t, expected3, got2)
	assert.Equal(t, expected2, got3)
	assert.Equal(t, expected1, got4)
	assert.Equal(t, l.Length(), 0)
}

func TestPushHeadMultiple(t *testing.T) {
	l := NewLinkedList[string]()
	expected1 := "pierre"
	expected2 := "henry"
	expected3 := "marie"

	l.PushHead(expected1)
	l.PushHead(expected2)
	l.PushHead(expected3)
	assert.Equal(t, l.Length(), 3)

	got1, err := l.PopHead()
	assert.NoError(t, err)

	got2, err := l.PopHead()
	assert.NoError(t, err)

	got3, err := l.PopHead()
	assert.NoError(t, err)

	assert.Equal(t, expected3, got1)
	assert.Equal(t, expected2, got2)
	assert.Equal(t, expected1, got3)
	assert.Equal(t, l.Length(), 0)
}

func TestErrPopHeadTail(t *testing.T) {
	l := NewLinkedList[string]()
	_, err := l.PopTail()
	assert.Error(t, err)
	assert.Equal(t, err, ErrEmptyList)
}

func TestErrPopHeadEmpty(t *testing.T) {
	l := NewLinkedList[string]()
	_, err := l.PopHead()
	assert.Error(t, err)
	assert.Equal(t, err, ErrEmptyList)
}

func TestClearEmpty(t *testing.T) {
	l := NewLinkedList[string]()
	assert.Equal(t, 0, l.Length())

	l.Clear()
	assert.Equal(t, 0, l.Length())
}

func TestClearNotEmptyList(t *testing.T) {
	l := NewLinkedList[string]()
	l.PushHead("aaa")
	l.PushTail("bbb")
	assert.Equal(t, 2, l.Length())

	l.Clear()
	assert.Equal(t, 0, l.Length())
}

func TestPeekAtEmptyList(t *testing.T) {
	l := NewLinkedList[string]()
	_, err := l.PeekAt(0)
	assert.Equal(t, ErrEmptyList, err)
}

func TestPeekAtOutOfRange(t *testing.T) {
	l := NewLinkedList[string]()
	l.PushHead("aaa")
	l.PushTail("bbb")

	_, err := l.PeekAt(2)
	assert.Equal(t, ErrIndexOutOfRange, err)
}

func TestPeekAtHead(t *testing.T) {
	l := NewLinkedList[string]()

	l.PushHead("aaa")
	l.PushHead("bbb")
	l.PushTail("ccc")
	assert.Equal(t, 3, l.Length())

	got, err := l.PeekAt(0)
	assert.NoError(t, err)
	assert.Equal(t, 3, l.Length())
	assert.Equal(t, "bbb", got)
}

func TestPeekAtTail(t *testing.T) {
	l := NewLinkedList[string]()

	l.PushHead("aaa")
	l.PushHead("bbb")
	l.PushTail("ccc")
	assert.Equal(t, 3, l.Length())

	got, err := l.PeekAt(2)
	assert.NoError(t, err)
	assert.Equal(t, 3, l.Length())
	assert.Equal(t, "ccc", got)
}

func TestPeekAtMiddle(t *testing.T) {
	l := NewLinkedList[string]()

	l.PushHead("aaa")
	l.PushHead("bbb")
	l.PushTail("ccc")
	assert.Equal(t, 3, l.Length())

	got, err := l.PeekAt(1)
	assert.NoError(t, err)
	assert.Equal(t, 3, l.Length())
	assert.Equal(t, "aaa", got)
}

//fdsdf

func TestPopAtEmptyList(t *testing.T) {
	l := NewLinkedList[string]()
	_, err := l.PopAt(0)
	assert.Equal(t, ErrEmptyList, err)
}

func TestPopAtOutOfRange(t *testing.T) {
	l := NewLinkedList[string]()
	l.PushHead("aaa")
	l.PushTail("bbb")

	_, err := l.PopAt(2)
	assert.Equal(t, ErrIndexOutOfRange, err)
}

func TestPopAtHead(t *testing.T) {
	l := NewLinkedList[string]()

	l.PushHead("aaa")
	l.PushHead("bbb")
	l.PushTail("ccc")
	assert.Equal(t, 3, l.Length())

	got, err := l.PopAt(0)
	assert.NoError(t, err)
	assert.Equal(t, 2, l.Length())
	assert.Equal(t, "bbb", got)
}

func TestPopAtTail(t *testing.T) {
	l := NewLinkedList[string]()

	l.PushHead("aaa")
	l.PushHead("bbb")
	l.PushTail("ccc")
	assert.Equal(t, 3, l.Length())

	got, err := l.PopAt(2)
	assert.NoError(t, err)
	assert.Equal(t, 2, l.Length())
	assert.Equal(t, "ccc", got)
}

func TestPopAtMiddle(t *testing.T) {
	l := NewLinkedList[string]()

	l.PushHead("aaa")
	l.PushHead("bbb")
	l.PushTail("ccc")
	assert.Equal(t, 3, l.Length())

	got, err := l.PopAt(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, l.Length())
	assert.Equal(t, "aaa", got)
}

// func TestContainsEmptylist(t *testing.T) {
// 	l := NewLinkedList[string]()
// 	b := l.Contains("pierre")
// 	assert.Equal(t, false, b)
// }

// func TestContainsFalse(t *testing.T) {
// 	l := NewLinkedList[string]()
// 	l.PushHead("aaa")
// 	l.PushHead("bbb")
// 	l.PushTail("ccc")
// 	assert.Equal(t, 3, l.Length())

// 	b := l.Contains("pierre")
// 	assert.Equal(t, false, b)
// }

// func TestContainsTrue(t *testing.T) {
// 	l := NewLinkedList[string]()
// 	l.PushHead("aaa")
// 	l.PushHead("bbb")
// 	l.PushTail("ccc")
// 	assert.Equal(t, 3, l.Length())

// 	b := l.Contains("aaa")
// 	assert.Equal(t, true, b)
// }

// //

// func TestIndexOfEmptylist(t *testing.T) {
// 	l := NewLinkedList[string]()
// 	_, err := l.IndexOf("pierre")
// 	assert.Error(t, err)
// 	assert.Equal(t, ErrEmptyList, err)
// }

// func TestIndexOfNotFound(t *testing.T) {
// 	l := NewLinkedList[string]()
// 	l.PushHead("aaa")
// 	l.PushHead("bbb")
// 	l.PushTail("ccc")
// 	assert.Equal(t, 3, l.Length())

// 	_, err := l.IndexOf("pierre")
// 	assert.Error(t, err)
// 	assert.Equal(t, ErrNotFound, err)
// }

// func TestIndexOfFound(t *testing.T) {
// 	l := NewLinkedList[string]()
// 	l.PushHead("aaa")
// 	l.PushHead("bbb")
// 	l.PushTail("ccc")
// 	assert.Equal(t, 3, l.Length())

// 	index, err := l.IndexOf("aaa")
// 	assert.NoError(t, err)
// 	assert.Equal(t, 1, index)
// }

// //

// func TestPopEmptylist(t *testing.T) {
// 	l := NewLinkedList[string]()
// 	_, err := l.Pop("pierre")
// 	assert.Error(t, err)
// 	assert.Equal(t, ErrEmptyList, err)
// }

// func TestPopNotFound(t *testing.T) {
// 	l := NewLinkedList[string]()
// 	l.PushHead("aaa")
// 	l.PushHead("bbb")
// 	l.PushTail("ccc")
// 	assert.Equal(t, 3, l.Length())

// 	_, err := l.Pop("pierre")
// 	assert.Error(t, err)
// 	assert.Equal(t, ErrNotFound, err)
// 	assert.Equal(t, 3, l.Length())
// }

// func TestPopFound(t *testing.T) {
// 	l := NewLinkedList[string]()
// 	l.PushHead("aaa")
// 	l.PushHead("bbb")
// 	l.PushTail("ccc")
// 	assert.Equal(t, 3, l.Length())

// 	index, err := l.Pop("aaa")
// 	assert.NoError(t, err)
// 	assert.Equal(t, "aaa", index)
// 	assert.Equal(t, 2, l.Length())
// }

func TestEmptyLinkedList(t *testing.T) {
	l := NewLinkedList[string]()
	l.PushHead("aaa")
	l.PushHead("bbb")
	l.PushTail("ccc")
	assert.Equal(t, 3, l.Length())
	assert.False(t, l.IsEmpty())

	l.PopTail()
	l.PopTail()
	l.PopTail()
	assert.Equal(t, 0, l.Length())
	assert.True(t, l.IsEmpty())
}

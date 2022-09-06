package datt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackLinkedListLengthIncreaseAndDecrease(t *testing.T) {
	q := NewStack[string]()
	assert.Equal(t, 0, q.Length())

	q.Push("pierre")
	assert.Equal(t, 1, q.Length())

	_, err := q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 0, q.Length())
}

func TestStackLinkedListPushOnce(t *testing.T) {
	q := NewStack[string]()
	expected := "pierre"

	q.Push(expected)

	got, err := q.Pop()
	assert.NoError(t, err)

	assert.Equal(t, expected, got)
}

func TestStackLinkedListPushMultiple(t *testing.T) {
	q := NewStack[string]()
	expected1 := "pierre"
	expected2 := "marie"
	expected3 := "henry"
	expected4 := "carol"

	q.Push(expected1)
	q.Push(expected2)
	q.Push(expected3)
	q.Push(expected4)

	got1, err := q.Pop()
	assert.NoError(t, err)
	got2, err := q.Pop()
	assert.NoError(t, err)
	got3, err := q.Pop()
	assert.NoError(t, err)
	got4, err := q.Pop()
	assert.NoError(t, err)

	assert.Equal(t, expected4, got1)
	assert.Equal(t, expected3, got2)
	assert.Equal(t, expected2, got3)
	assert.Equal(t, expected1, got4)
}

func TestStackLinkedListClearMultiple(t *testing.T) {
	q := NewStack[string]()
	expected1 := "pierre"
	expected2 := "marie"
	expected3 := "henry"
	expected4 := "carol"

	q.Push(expected1)
	q.Push(expected2)
	q.Push(expected3)
	q.Push(expected4)
	assert.Equal(t, 4, q.Length())

	q.Clear()
	assert.Equal(t, 0, q.Length())
}

func TestStackLinkedListPeakFrontAndBack(t *testing.T) {
	q := NewStack[string]()
	expected1 := "pierre"
	expected2 := "marie"
	expected3 := "henry"
	expected4 := "carol"

	q.Push(expected1)
	q.Push(expected2)
	q.Push(expected3)
	q.Push(expected4)

	front, err := q.PeekFront()
	assert.NoError(t, err)
	assert.Equal(t, expected4, front)

	back, err := q.PeekBack()
	assert.NoError(t, err)
	assert.Equal(t, expected1, back)

	assert.Equal(t, 4, q.Length())
}

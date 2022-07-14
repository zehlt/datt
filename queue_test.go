package datt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthIncreaseAndDecrease(t *testing.T) {
	q := NewQueue[string]()
	assert.Equal(t, 0, q.Length())

	q.Enqueue("pierre")
	assert.Equal(t, 1, q.Length())

	_, err := q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 0, q.Length())
}

func TestEnqueueOnce(t *testing.T) {
	q := NewQueue[string]()
	expected := "pierre"

	q.Enqueue(expected)

	got, err := q.Dequeue()
	assert.NoError(t, err)

	assert.Equal(t, expected, got)
}

func TestEnqueueMultiple(t *testing.T) {
	q := NewQueue[string]()
	expected1 := "pierre"
	expected2 := "marie"
	expected3 := "henry"
	expected4 := "carol"

	q.Enqueue(expected1)
	q.Enqueue(expected2)
	q.Enqueue(expected3)
	q.Enqueue(expected4)

	got1, err := q.Dequeue()
	assert.NoError(t, err)
	got2, err := q.Dequeue()
	assert.NoError(t, err)
	got3, err := q.Dequeue()
	assert.NoError(t, err)
	got4, err := q.Dequeue()
	assert.NoError(t, err)

	assert.Equal(t, expected1, got1)
	assert.Equal(t, expected2, got2)
	assert.Equal(t, expected3, got3)
	assert.Equal(t, expected4, got4)
}

func TestClearMultiple(t *testing.T) {
	q := NewQueue[string]()
	expected1 := "pierre"
	expected2 := "marie"
	expected3 := "henry"
	expected4 := "carol"

	q.Enqueue(expected1)
	q.Enqueue(expected2)
	q.Enqueue(expected3)
	q.Enqueue(expected4)
	assert.Equal(t, 4, q.Length())

	q.Clear()
	assert.Equal(t, 0, q.Length())
}

func TestPeakFrontAndBack(t *testing.T) {
	q := NewQueue[string]()
	expected1 := "pierre"
	expected2 := "marie"
	expected3 := "henry"
	expected4 := "carol"

	q.Enqueue(expected1)
	q.Enqueue(expected2)
	q.Enqueue(expected3)
	q.Enqueue(expected4)

	front, err := q.PeekFront()
	assert.NoError(t, err)
	assert.Equal(t, expected1, front)

	back, err := q.PeekBack()
	assert.NoError(t, err)
	assert.Equal(t, expected4, back)

	assert.Equal(t, 4, q.Length())
}

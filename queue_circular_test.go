package datt_test

import (
	"testing"

	"github.com/zehlt/datt"
)

func TestQueueCircularArray(t *testing.T) {
	q := datt.QueueCircularArray[float32]{}
	got := q.Len()

	AssertEqual(t, got, 0)
}

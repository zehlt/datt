package datt

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyBinaryHeap = errors.New("empty list")
)

func NewBinaryHeap[T any](compare func(current T, other T) CompareResult) *BinaryHeap[T] {
	return &BinaryHeap[T]{
		compare: compare,
	}
}

type BinaryHeap[T any] struct {
	compare func(current T, other T) CompareResult
	lenght  int
	arr     []T
}

func (bh *BinaryHeap[T]) PeekFront() (T, error) {
	if bh.lenght <= 0 {
		var t T
		return t, ErrEmptyBinaryHeap
	}

	return bh.arr[0], nil
}

func (bh *BinaryHeap[T]) PeekBack() (T, error) {
	if bh.lenght <= 0 {
		var t T
		return t, ErrEmptyBinaryHeap
	}

	return bh.arr[bh.lenght-1], nil
}

func (bh *BinaryHeap[T]) Push(value T) {
	bh.arr = append(bh.arr, value)
	bh.lenght++

	if bh.lenght <= 1 {
		return
	}

	bh.growUp(value, bh.lenght-1)
}

func (bh *BinaryHeap[T]) Pop() (T, error) {
	if bh.lenght <= 0 {
		var t T
		return t, ErrEmptyBinaryHeap
	}

	if bh.lenght == 1 {
		value := bh.arr[0]
		bh.lenght--
		return value, nil
	}

	rootValue, _ := bh.PeekFront()
	lastValue, _ := bh.PeekBack()

	bh.arr[0] = lastValue
	bh.lenght--

	bh.grownDown(lastValue, 0)
	return rootValue, nil
}

func (bh *BinaryHeap[T]) Iterate(callback func(value T) bool) {
	for i := 0; i < bh.lenght; i++ {
		shouldStop := callback(bh.arr[i])
		if shouldStop {
			break
		}
	}
}

func (bh *BinaryHeap[T]) Length() int {
	return bh.lenght
}

func (bh *BinaryHeap[T]) Clear() {
	bh.arr = make([]T, 0)
	bh.lenght = 0
}

func (bh *BinaryHeap[T]) String() string {
	var builder strings.Builder

	bh.Iterate(func(value T) bool {
		builder.WriteString(fmt.Sprintf("[%v]", value))
		return false
	})

	return builder.String()
}

func (bh *BinaryHeap[T]) leftChildIndex(index int) int {
	return (index * 2) + 1
}

func (bh *BinaryHeap[T]) rightChildIndex(index int) int {
	return (index * 2) + 2
}

func (bh *BinaryHeap[T]) parentIndex(index int) int {
	return (index - 1) / 2
}

func (bh *BinaryHeap[T]) grownDown(value T, position int) {
	if position >= bh.lenght-1 {
		return
	}

	leftIndex := bh.leftChildIndex(position)
	rightIndex := bh.rightChildIndex(position)

	if leftIndex >= bh.lenght || rightIndex >= bh.lenght {
		return
	}

	leftValue := bh.arr[leftIndex]
	rightValue := bh.arr[rightIndex]

	leftChildResult := bh.compare(leftValue, rightValue)
	if leftChildResult == HIGHER {
		result := bh.compare(value, leftValue)

		if result == LOWER {
			bh.arr[position] = leftValue
			bh.arr[leftIndex] = value
			bh.grownDown(value, leftIndex)
		} else {
			return
		}
	} else {
		result := bh.compare(value, rightValue)

		if result == LOWER {
			bh.arr[position] = rightValue
			bh.arr[rightIndex] = value
			bh.grownDown(value, rightIndex)
		} else {
			return
		}
	}
}

func (bh *BinaryHeap[T]) growUp(value T, position int) {
	if position == 0 {
		return
	}

	parentIndex := bh.parentIndex(position)
	parentValue := bh.arr[parentIndex]

	result := bh.compare(value, parentValue)

	switch result {
	case HIGHER:
		bh.arr[parentIndex] = value
		bh.arr[position] = parentValue
	case LOWER:
		return
	case EQUAL:
		return
	}

	bh.growUp(value, parentIndex)
}

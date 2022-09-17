package datt

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEntityDoesNotExist           = errors.New("entity does not exist")
	ErrInternalUnableToCreateEntity = errors.New("internal error: unable to create an entity")
)

type cellType int

const (
	emptyCell cellType = iota
	occupiedCell
)

type cell[T any] struct {
	t          cellType
	generation uint64
	next       int
	data       T
}

func (c *cell[T]) String() string {
	if c.t == occupiedCell {
		return fmt.Sprintf("[gen:%d, data:%v]", c.generation, c.data)
	}
	return fmt.Sprintf("(next:%d)", c.next)
}

type ArenaKey struct {
	Gen uint64
	Id  int
}

type Arena[T any] struct {
	cells       []cell[T]
	current_gen uint64
	first_free  int
	size        int
}

func NewArena[T any]() *Arena[T] {
	return &Arena[T]{}
}

func (a *Arena[T]) Create(value T) ArenaKey {
	if a.first_free >= a.size {
		a.first_free++
		a.cells = append(a.cells, cell[T]{
			t:          occupiedCell,
			generation: a.current_gen,
			data:       value,
		})
		a.size++
		return ArenaKey{
			Gen: a.current_gen,
			Id:  a.size - 1,
		}
	}

	currentCellIndex := a.first_free
	nextFree := a.cells[currentCellIndex].next

	a.cells[currentCellIndex] = cell[T]{
		t:          occupiedCell,
		generation: a.current_gen,
		data:       value,
	}
	a.first_free = nextFree

	return ArenaKey{
		Gen: a.current_gen,
		Id:  currentCellIndex,
	}
}

func (a *Arena[T]) Contains(key ArenaKey) bool {
	if key.Id >= a.size || key.Id < 0 {
		return false
	}

	cell := a.cells[key.Id]
	if cell.t == emptyCell {
		return false
	}

	if cell.generation != key.Gen {
		return false
	}

	return true
}

func (a *Arena[T]) Remove(key ArenaKey) error {
	ok := a.Contains(key)
	if !ok {
		return ErrEntityDoesNotExist
	}

	a.current_gen++
	a.cells[key.Id] = cell[T]{
		t:          emptyCell,
		generation: a.current_gen,
		next:       a.first_free,
	}
	a.first_free = key.Id

	return nil
}

func (a *Arena[T]) Get(key ArenaKey) (T, error) {
	ok := a.Contains(key)
	if !ok {
		var t T
		return t, ErrEntityDoesNotExist
	}

	return a.cells[key.Id].data, nil
}

func (a *Arena[T]) Set(key ArenaKey, value T) error {
	ok := a.Contains(key)
	if !ok {
		return ErrEntityDoesNotExist
	}

	a.cells[key.Id].data = value
	return nil
}

func (a *Arena[T]) Iterate(callback func(index int, value T) bool) {
	for i, cell := range a.cells {
		shouldStop := callback(i, cell.data)
		if shouldStop {
			break
		}
	}
}

func (a *Arena[T]) Clear() {
	a.cells = make([]cell[T], 0)
	a.current_gen = 0
	a.first_free = 0
	a.size = 0
}

func (a *Arena[T]) String() string {
	var builder strings.Builder

	for id, cell := range a.cells {
		builder.WriteString(fmt.Sprintf("%d -> %s\n", id, cell.String()))
	}

	return builder.String()
}

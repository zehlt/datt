package datt

import (
	"errors"
)

var (
	ErrDispatchNoHandlerRegistered = errors.New("type has not been registered")
)

type Dispatch[T comparable, D any] struct {
	handlers map[T][]func(data D)
}

func NewDispatch[T comparable, D any]() *Dispatch[T, D] {
	return &Dispatch[T, D]{
		handlers: make(map[T][]func(data D)),
	}
}

func (d *Dispatch[T, D]) Subscribe(t T, f func(data D)) {
	d.ensureArrayExists(t)

	d.handlers[t] = append(d.handlers[t], f)
}

func (d *Dispatch[T, D]) ensureArrayExists(t T) {
	_, ok := d.handlers[t]
	if !ok {
		d.handlers[t] = make([]func(data D), 0)
	}
}

func (d *Dispatch[T, D]) Publish(t T, data D) error {
	functions, ok := d.handlers[t]
	if !ok {
		return ErrDispatchNoHandlerRegistered
	}

	for _, fn := range functions {
		fn(data)
	}

	return nil
}

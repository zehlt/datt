package datt

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrItemAlreadyExistsSet = errors.New("item already exists")
	ErrItemDoesNotExistSet  = errors.New("item does not exist")
)

type Set[T comparable] struct {
	hash map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		hash: make(map[T]bool),
	}
}

func (s *Set[T]) Append(value T) error {
	_, ok := s.hash[value]
	if ok {
		return ErrItemAlreadyExistsSet
	}

	s.hash[value] = true
	return nil
}

func (s *Set[T]) Remove(value T) error {
	_, ok := s.hash[value]
	if !ok {
		return ErrItemDoesNotExistSet
	}

	delete(s.hash, value)
	return nil
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	newSet := NewSet[T]()

	for key := range s.hash {
		newSet.Append(key)
	}

	for key := range other.hash {
		newSet.Append(key)
	}

	return newSet
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	newSet := NewSet[T]()

	var biggest map[T]bool
	var smallest map[T]bool

	if len(s.hash) > len(other.hash) {
		biggest = s.hash
		smallest = other.hash
	} else {
		biggest = other.hash
		smallest = s.hash
	}

	for key := range biggest {
		_, ok := smallest[key]
		if ok {
			newSet.Append(key)
		}
	}

	return newSet
}

func (s *Set[T]) Xor(other *Set[T]) *Set[T] {
	newSet := NewSet[T]()

	for key := range s.hash {
		_, ok := other.hash[key]
		if !ok {
			newSet.Append(key)
		}
	}

	for key := range other.hash {
		_, ok := s.hash[key]
		if !ok {
			newSet.Append(key)
		}
	}

	return newSet
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	newSet := NewSet[T]()

	for key := range s.hash {
		_, ok := other.hash[key]
		if !ok {
			newSet.Append(key)
		}
	}

	return newSet
}

func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for key := range s.hash {
		_, ok := other.hash[key]
		if !ok {
			return false
		}
	}

	return true
}

func (s *Set[T]) IsEqual(other *Set[T]) bool {
	if s.Length() != other.Length() {
		return false
	}

	for key := range s.hash {
		_, ok := other.hash[key]
		if !ok {
			return false
		}
	}

	return true
}

func (s *Set[T]) Iterate(callback func(value T) bool) {
	for key := range s.hash {
		shouldStop := callback(key)
		if shouldStop {
			break
		}
	}
}

func (s *Set[T]) Length() int {
	return len(s.hash)
}

func (s *Set[T]) String() string {
	var builder strings.Builder

	s.Iterate(func(value T) bool {
		builder.WriteString(fmt.Sprintf("%v, ", value))
		return false
	})

	return builder.String()
}

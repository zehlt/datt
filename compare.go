package datt

import "golang.org/x/exp/constraints"

type CompareResult int

const (
	LOWER CompareResult = iota
	EQUAL
	HIGHER
)

func CompareOrdered[T constraints.Ordered](current T, other T) CompareResult {
	if current > other {
		return HIGHER
	} else if current < other {
		return LOWER
	} else {
		return EQUAL
	}
}

type EqualResult int

const (
	DIFFERENT EqualResult = iota
	SAME
)

func EqualComparable[T comparable](current T, other T) EqualResult {
	if current == other {
		return SAME
	} else {
		return DIFFERENT
	}
}

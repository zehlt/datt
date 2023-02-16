package datt

type CompareResult int

const (
	LOWER CompareResult = iota
	EQUAL
	HIGHER
)

type CompareFunc[T any] func(current T, other T) CompareResult

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~string | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func CompareOrdered[T Ordered](current T, other T) CompareResult {
	if current > other {
		return HIGHER
	} else if current < other {
		return LOWER
	} else {
		return EQUAL
	}
}

type OrderTraversal int

const (
	PreOrder OrderTraversal = iota
	InOrder
	PostOrder
	LevelOrder
)

package datt

type Tuple[A any, B any] struct {
	First  A
	Second B
}

type Triple[A any, B any, C any] struct {
	First  A
	Second B
	Third  C
}

type Quadruple[A any, B any, C any, D any] struct {
	First  A
	Second B
	Third  C
	Fourth D
}

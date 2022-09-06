package datt

import (
	"log"

	"golang.org/x/exp/constraints"
)

type Vertex[T comparable] struct {
	value     T
	adjacents []*Vertex[T]
}

func NewVertex[T comparable](value T) *Vertex[T] {
	return &Vertex[T]{
		value: value,
	}
}

func (v *Vertex[T]) SetValue(value T) {
	v.value = value
}

func (v *Vertex[T]) AddAdjacentExclusive(other *Vertex[T]) {
	// TODO: check if aleardy add
	v.adjacents = append(v.adjacents, other)
}

func (v *Vertex[T]) AddAdjacentMutual(other *Vertex[T]) {
	// TODO: check if aleardy add
	v.adjacents = append(v.adjacents, other)
	other.AddAdjacentExclusive(v)
}

func DfsTraverse[T comparable](vertex *Vertex[T], visitedHash map[T]bool) {
	visitedHash[vertex.value] = true

	log.Println(vertex.value)

	for _, adj := range vertex.adjacents {
		_, ok := visitedHash[adj.value]
		if !ok {
			DfsTraverse(adj, visitedHash)
		}
	}
}

func BfsTraverse[T comparable](startVertex *Vertex[T]) {
	visitedHash := make(map[T]bool)

	queue := NewQueue[*Vertex[T]]()
	queue.Enqueue(startVertex)

	for !queue.IsEmpty() {
		currentVertex, _ := queue.Dequeue()
		visitedHash[currentVertex.value] = true

		log.Println(currentVertex.value)

		for _, adj := range currentVertex.adjacents {
			isAlreadyVisited := visitedHash[adj.value]

			if !isAlreadyVisited {
				queue.Enqueue(adj)
			}
		}

	}
}

type WeightedVertex[V constraints.Ordered, W any] struct {
	value     V
	adjacents map[*WeightedVertex[V, W]]W
}

func NewWeightedVertex[V constraints.Ordered, W any](value V) *WeightedVertex[V, W] {
	return &WeightedVertex[V, W]{
		value:     value,
		adjacents: make(map[*WeightedVertex[V, W]]W),
	}
}

func (v *WeightedVertex[V, W]) AddAdjacentExclusive(other *WeightedVertex[V, W], weight W) {
	_, ok := v.adjacents[other]
	if !ok {
		v.adjacents[other] = weight
	}
}

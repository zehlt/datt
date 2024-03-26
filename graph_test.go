package datt_test

import (
	"testing"

	"github.com/zehlt/datt"
)

func TestGraphAdjacencyListNew(t *testing.T) {
	t.Run("creating 10 element graph", func(t *testing.T) {
		want := 10
		graph := datt.NewGraphAdjacencyList(want)
		got := graph.Vertices()

		AssertEqual(t, got, want)
	})
}

func TestGraphAdjacencyListAddEdge(t *testing.T) {
	t.Run("add edge on empty graph must panic", func(t *testing.T) {
		g := datt.NewGraphAdjacencyList(0)

		AssertPanic(t, func() {
			g.AddEdge(2, 3)
		})
	})

	t.Run("add edge on bound ", func(t *testing.T) {
		g := datt.NewGraphAdjacencyList(10)

		g.AddEdge(2, 3)

		var got []int
		g.Adjacents(2, func(adj int) {
			got = append(got, adj)
		})
		want := []int{3}

		AssertEqual(t, got, want)
	})

	// log.Println("MY MAN")
	// graph := datt.NewGraphAdjacencyList(100)
	// AssertNotEqual(t, graph, nil)
}

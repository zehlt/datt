package datt

type GraphAdjacencyList struct {
	vertices    int
	adjacencies []LinkedList[int]
}

func NewGraphAdjacencyList(vertices int) *GraphAdjacencyList {
	return &GraphAdjacencyList{
		vertices:    vertices,
		adjacencies: make([]LinkedList[int], vertices),
	}
}

func (g *GraphAdjacencyList) AddEdge(v int, w int) {
	g.boundCheck(v)
	g.boundCheck(w)

	g.adjacencies[v].PushFront(w)
	g.adjacencies[w].PushFront(v)
}

func (g *GraphAdjacencyList) Adjacents(v int, fn func(adj int)) {
	g.boundCheck(v)

	g.adjacencies[v].Do(func(a int) {
		fn(a)
	})
}

func (g *GraphAdjacencyList) Vertices() int {
	return g.vertices
}

func (g *GraphAdjacencyList) boundCheck(v int) {
	if v >= g.vertices || v < 0 {
		panic("out of bound")
	}
}

// func (g *GraphAdjacencyList) Edges() int {
// 	return 0
// }

// func (g *GraphAdjacencyList) Degree() int {
// 	return 0
// }

// func (g *GraphAdjacencyList) MaxDegree() int {
// 	return 0
// }

// func (g *GraphAdjacencyList) AverageDegree() int {
// 	return 0
// }

// func (g *GraphAdjacencyList) DepthFirstSearch() {

// }

// func (g *GraphAdjacencyList) BreathFirstSearch() int {
// 	return 0
// }

// func (g *GraphAdjacencyList) String() string {
// 	return ""
// }

// --

type Edge struct {
}

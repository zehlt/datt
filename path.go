package datt

type Graph interface {
	AddEdge(v int, w int)
	Adjacents(v int, fn func(adj int))
	Vertices() int
}

type DepthFirstSearch struct {
	g      Graph
	marked []bool
	edgeTo []int
	s      int
}

func (p *DepthFirstSearch) NewDepthFirstSearch(g Graph, s int) *DepthFirstSearch {
	size := g.Vertices()

	path := &DepthFirstSearch{
		g:      g,
		marked: make([]bool, size),
		edgeTo: make([]int, size),
		s:      s,
	}
	path.dfs(s)

	return path
}

func (p *DepthFirstSearch) HasPathTo(v int) bool {
	return p.marked[v]
}

func (p *DepthFirstSearch) PathTo(v int, fn func(a int)) {
	if !p.HasPathTo(v) {
		return
	}

	var arr []int
	var road int = v

	for road != p.s {
		road = p.edgeTo[road]
		arr = append(arr, road)
	}

	for _, edge := range arr {
		fn(edge)
	}
}

func (p *DepthFirstSearch) dfs(v int) {
	p.marked[v] = true

	p.g.Adjacents(v, func(adj int) {
		if p.marked[adj] {
			return
		}

		p.edgeTo[adj] = v
		p.dfs(adj)
	})
}

type BreadthFirstSearch struct {
	g      Graph
	marked []bool
	edgeTo []int
	distTo []int
	queue  QueueCircularArray[int]
}

func NewBreadthFirstSearch(g Graph, v int) *BreadthFirstSearch {
	size := g.Vertices()
	path := &BreadthFirstSearch{
		g:      g,
		marked: make([]bool, size),
		edgeTo: make([]int, size),
		distTo: make([]int, size),
		queue:  QueueCircularArray[int]{},
	}

	path.bfs(v)

	return path
}

func (p *BreadthFirstSearch) HasPathTo(v int) {
}

func (p *BreadthFirstSearch) PathTo(v int, fn func(a int)) {
}

func (p *BreadthFirstSearch) LengthTo(v int) int {
	return 0
}

func (p *BreadthFirstSearch) bfs(v int) {
	p.queue.Enqueue(v)
	p.distTo[v] = 0
	p.marked[v] = true
	distance := 1

	for p.queue.Len() > 0 {
		vertex, _ := p.queue.Dequeue()

		p.g.Adjacents(vertex, func(adj int) {
			if p.marked[adj] {
				return
			}

			p.queue.Enqueue(adj)
			p.edgeTo[adj] = vertex
			p.distTo[adj] = distance
			p.marked[adj] = true
		})

		distance++
	}
}

// type ConnectedComponent struct {
// 	marked []bool
// 	cc     []int
// }

package datt

// Weighted with path compression
type UnionFind struct {
	cap  int
	arr  []int
	size []int
}

// O(N)
func NewUnionFind(cap int) *UnionFind {
	if cap <= 0 {
		return nil
	}

	uf := &UnionFind{
		arr:  make([]int, cap),
		size: make([]int, cap),
		cap:  cap,
	}

	for i := 0; i < cap; i++ {
		uf.arr[i] = i
		uf.size[i] = 1
	}

	return uf
}

// O(log N)
func (u *UnionFind) Union(a int, b int) bool {
	if u.isOutOfRange(a, b) {
		return false
	}

	roota := u.root(a)
	rootb := u.root(b)

	if roota == rootb {
		return false
	}

	sizea := u.size[roota]
	sizeb := u.size[rootb]

	if sizea > sizeb {
		u.arr[rootb] = roota
		u.size[roota] += sizeb
	} else {
		u.arr[roota] = rootb
		u.size[rootb] += sizea
	}

	return true
}

// O(log N)
func (u *UnionFind) Find(a int, b int) bool {
	if u.isOutOfRange(a, b) {
		return false
	}

	return u.root(a) == u.root(b)
}

func (u *UnionFind) root(v int) int {
	for v != u.arr[v] {
		u.arr[v] = u.arr[u.arr[v]]
		v = u.arr[v]
	}

	return v
}

func (u *UnionFind) isOutOfRange(a int, b int) bool {
	return a >= u.cap || a < 0 || b >= u.cap || b < 0
}
